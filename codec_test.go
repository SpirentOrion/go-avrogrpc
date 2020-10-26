//go:generate go run ./go-avrogrpc -i test.avpr -o codegen_test.go -p avrogrpc -t ./go-avrogrpc/codegen.template

package avrogrpc

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// testServer implements TestServer.
type testServer struct {
	UnimplementedTestServer
}

func (s *testServer) Hello(ctx context.Context, in map[string]interface{}) (interface{}, error) {
	name := in["name"]
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello %s!", name), nil
}

func (s *testServer) Echo(ctx context.Context, in map[string]interface{}) (interface{}, error) {
	return in["record"], nil
}

func (s *testServer) Ack(ctx context.Context, in map[string]interface{}) (interface{}, error) {
	return nil, nil
}

func (s *testServer) ServerCounter(in map[string]interface{}, srv Test_ServerCounterServer) error {
	for i := 1; i <= 10; i++ {
		if err := srv.Send(i); err != nil {
			return err
		}
	}
	return nil
}

func (s *testServer) ClientCounter(srv Test_ClientCounterServer) error {
	var values []int32
	for {
		m, err := srv.Recv()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		in, _ := m.(map[string]interface{})
		value, _ := in["counter"].(int32)
		values = append(values, value)
	}
	return srv.SendAndClose(values)
}

func (s *testServer) BidirCounter(srv Test_BidirCounterServer) error {
	for {
		m, err := srv.Recv()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		in, _ := m.(map[string]interface{})
		value, _ := in["counter"].(int32)
		if err = srv.Send(value); err != nil {
			return err
		}
	}
	return nil
}

func TestCodec(t *testing.T) {
	scratch, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(scratch)

	sock := filepath.Join(scratch, "sock")
	addr, err := net.ResolveUnixAddr("unix", sock)
	require.NoError(t, err)

	// Start the server
	lis, err := net.ListenUnix("unix", addr)
	require.NoError(t, err)

	s := grpc.NewServer()
	RegisterTestServer(s, &testServer{})

	go s.Serve(lis)
	defer s.GracefulStop()

	// Connect a client
	conn, err := grpc.Dial("unix:"+addr.String(),
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.CallContentSubtype("avro")))
	require.NoError(t, err)
	defer conn.Close()

	c := NewTestClient(conn)

	// Test the Hello message
	req := map[string]interface{}{"name": "David"}
	resp, err := c.Hello(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, resp, "Hello David!")

	resp, err = c.Hello(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, resp, "Hello World!")

	// Test the Echo message
	hash, err := hex.DecodeString("f76e57f5310c47eeae187c8f4a143637")
	require.NoError(t, err)

	req = map[string]interface{}{
		"record": map[string]interface{}{
			"name": "David",
			"kind": "BAR",
			"hash": hash,
		},
	}
	resp, err = c.Echo(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, req["record"], resp)

	// Test the Ack message
	resp, err = c.Ack(context.Background(), nil)
	require.NoError(t, err)
	require.Nil(t, resp)

	// Test response streaming
	srvCounter, err := c.ServerCounter(context.Background(), nil)
	require.NoError(t, err)

	var m interface{}
	for i := 1; ; i++ {
		m, err = srvCounter.Recv()
		if err == io.EOF {
			require.Nil(t, m)
			require.Equal(t, 10, i-1)
			break
		}
		require.NoError(t, err)
		require.EqualValues(t, i, m)
	}

	// Test request streaming
	cliCounter, err := c.ClientCounter(context.Background())
	for i := 1; i <= 10; i++ {
		err = cliCounter.Send(map[string]interface{}{"counter": i})
		require.NoError(t, err)
	}

	m, err = cliCounter.CloseAndRecv()
	require.NoError(t, err)
	require.IsType(t, []interface{}{}, m)
	require.Len(t, m, 10)

	// Test bi-diretional streaming
	biCounter, err := c.BidirCounter(context.Background())
	for i := 1; i <= 10; i++ {
		err = biCounter.Send(map[string]interface{}{"counter": i})
		require.NoError(t, err)

		m, err = biCounter.Recv()
		require.NoError(t, err)
		require.EqualValues(t, i, m)
	}

	err = biCounter.CloseSend()
	require.NoError(t, err)

	// Test the Unimplemented message
	_, err = c.Unimplemented(context.Background(), nil)
	require.Error(t, err)
	require.Equal(t, codes.Unimplemented, status.Code(err))
}

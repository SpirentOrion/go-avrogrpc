# A gRPC+Avro Codec for Go

This is an [Avro](https://avro.apache.org/docs/current/) codec for the [Go
implementation of gRPC](https://github.com/grpc/grpc-go). It also includes a
code generator that compiles [Avro protocol
definitions](https://avro.apache.org/docs/current/spec.html#Protocol+Declaration)
(i.e. `.avpr` files in JSON format) into gRPC client and server stubs.

## Installation

Install the package and the associated code generator:

    go get -u github.com/SpirentOrion/go-avrogrpc/...

## Usage

gRPC client and server stubs supporting the Avro codec are generated from an
[Avro protocol definition](https://avro.apache.org/docs/current/spec.html#Protocol+Declaration),
typically stored in an `.avpr` file:

    go-avrogrpc -i /path/to/avpr.file -o /path/to/source.go
    
This is the moral equivalent of using `protoc` to compile a protocol
buffer-encoded gRPC service definition into Go client and server code. However,
there are [important differences](#grpc-vs-avro-rpc).

By default, code is generated using a package name that derives from the Avro
protocol name. This may be overriden from the command line:

    go-avrogrpc -i /path/to/avpr.file -o /path/to/source.go -p mypackage

Multiple Avro protocols may be compiled into a single package without conflict,
as long as their protocol names are unique:

    go-avrogrpc -i /path/to/avpr.file1 -o /path/to/source1.go -p mypackage
    go-avrogrpc -i /path/to/avpr.file2 -o /path/to/source2.go -p mypackage

At some point within your own code, you must register the Avro codec with the
gRPC library:

```go
package mypackage

import _ "github.com/SpirentOrion/go-avrorpc"
```

On the client-side, to use the Avro codec, the `CallOption` `CallContentSubtype`
should be used as follows:

```go
response, err := myclient.MyCall(ctx, request, grpc.CallContentSubtype("avro"))
```

As a reminder, all `CallOption`s may be converted into `DialOption`s that become
the default for all RPCs sent through a client using `grpc.WithDefaultCallOptions`:

```go
myclient := grpc.Dial(ctx, target, grpc.WithDefaultCallOptions(grpc.CallContentSubtype("avro")))
```

Messages will be sent along with headers indicating the codec (`Content-Type`
set to `application/grpc+avro`).

On the server-side, the codec is automatically registered via the package import
above.

For more information on the use of custom codecs with gRPC, see the [`grpc-go`
docs](https://github.com/grpc/grpc-go/blob/master/Documentation/encoding.md).

## gRPC vs Avro RPC

There are two major differences between gRPC and Avro RPC:

* Avro protocol specifications are written in JSON instead of the usual protocol
  buffer language
* Avro protocol specifications define Avro RPC service interfaces and not gRPC
  service interfaces

The second point is particularly important since there are semantic differences
between gRPC and Avro RPC:

* Avro RPC supports one-way messages (i.e. "fire and forget" messages), but gRPC
  does not
* Avro RPC supports structured errors, but gRPC fixes on an error code and a string
* gRPC supports streaming of both request and response messages, but Avro RPC
  does not
  
The `go-avrogrpc` code generator reconciles these differences as best as it can.

One-way messages are nominally supported. The gRPC server handler is expected to
return `nil`.

Structured errors are not supported at all. `go-avrogrpc` will warn (with a
message to `stderr`) if it encounters these but will not fail.

gRPC-style streaming is fully supported. However, since Avro protocol
specifications have no way to define this, support for streaming is indicated
using an extra, non-standard `stream` attribute on the Avro protocol's message
objects:

* `"stream": "request"` indicates client-to-server streaming
* `"stream": "response"` indicates server-to-client streaming
* `"stream": "bidir"` indicates bi-directional streaming

For examples, see [`test.avpr`](test.avpr).

## Programming Model

The codec and code generator take a strong dependency on the [`goavro`
library](https://github.com/linkedin/goavro).

Since `goavro`-driven serialization / deserialization requires a
message-specific codec, the actual gRPC codec implementation is a no-op:

```go
func (codec) Marshal(v interface{}) ([]byte, error) {
	// NB: generated code handles Avro marshaling using a message-specific codec
	return v.([]byte), nil
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	// NB: generated code handles Avro unmarshaling using a message-specific codec
	buf := v.(*[]byte)
	*buf = data
	return nil
}
```

This is a "pass-through" implementation that relies on generated client and
server code to handle the actual marshaling and unmarshaling. `go-avrogrpc`
generates this code automatically using the correct codec and
`goavro.BinaryFromNative` (for marshsaling) and `goavro.NativeFromBinary` (for
unmarshaling).

If you're using Avro with Go, then you're probably already familiar with
`goavro` and how it uses the Go type system. If not, you should at least read
[this](https://github.com/linkedin/goavro#translating-data).

With the exception of message types, the [Go Generated-code
Reference](https://grpc.io/docs/languages/go/generated-code/) from the gRPC
documentation is 100% valid when using `go-avrogrpc` generated code. However,
since `goavro` powers the codec, the message types need special consideration:

| Message  | Avro protocol type | Generated Go type        | Concrete Go type         |
| ---      | ---                | ---                      | ---                      |
| request  | anonymous record   | `map[string]interface{}` | `map[string]interface{}` |
| response | from `.avpr`       | `interface{}`            | depends on `goavro`      |

For example, unary methods on server interfaces are:

    Foo(context.Context, map[string]interface{}) (interface{}, error)

instead of:

    Foo(context.Context, *MsgA) (*MsgB, error)

Here, the request type is `map[string]interface{}` (since `goavro` translates
`record` to `map[string]interface{}`). The response type is `interface{}`. It's
up to the server implementation to return the correct concrete type. For
example, if the Avro response type is also a `record` then the return should be
`map[string]interface{}`. However, this isn't always the case. If the Avro
response type is a primitive type like `int`, then the handler should just
return `42`.

If there is a type mismatch between what a server handler returns and what
`goavro.BinaryFromNative` expects (based on the `goavro.Codec` built from the
Avro protocol specification), then this will produce a runtime error.

## Acknowledgements

* The [Go codegen template](go-avrogrpc/codegen.template) was
  inspired by inspection of code output by `protoc`. The [`grpc-go`
  project](https://github.com/grpc/grpc-go) Copyright header is maintained in
  this project's [LICENSE](LICENSE).
* Handling of gRPC streams using an extra attribute in the `.avpr` file was
  inspired by a [VMworld talk](https://www.youtube.com/watch?v=Zrvx4ZAuMCU)
  given by Clement Pang and Srujan Narkedamalli.

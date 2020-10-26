package avrogrpc

import "google.golang.org/grpc/encoding"

type codec struct{}

func init() {
	encoding.RegisterCodec(codec{})
}

func (codec) Name() string {
	return "avro"
}

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

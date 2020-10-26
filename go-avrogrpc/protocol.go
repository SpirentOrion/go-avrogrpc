package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/linkedin/goavro/v2"
)

// Protocol is the parsed, validated form of an Avro protocol definition.
type Protocol struct {
	Namespace string
	Protocol  string
	Doc       string
	Messages  map[string]*Message
}

// Message defines a specific message within an Avro protocol definition.
type Message struct {
	Doc               string
	RequestSchema     string
	RequestCodec      *goavro.Codec
	ResponseSchema    string
	ResponseCodec     *goavro.Codec
	ErrorsSchema      string
	ErrorsCodec       *goavro.Codec
	HasExplicitErrors bool
	OneWay            bool
	Stream            string
}

const (
	// RequestStream indicates that the client writes a sequence of messages
	// and sends them to the server, using a provided stream. Once the
	// client has finished writing the messages, it waits for the server to
	// read them and return its response.
	RequestStream = "request"

	// ResponseStream indicates that the client sends a request to the
	// server and gets a stream to read a sequence of messages back.
	ResponseStream = "response"

	// BidirStream indicates that both sides send a sequence of messages
	// using a read-write stream. The two streams operate independently, so
	// clients and servers can read and write in whatever order they like.
	BidirStream = "bidir"
)

type rawProtocol struct {
	Namespace string                   `json:"namespace"`
	Protocol  string                   `json:"protocol"`
	Doc       string                   `json:"doc"`
	Types     []map[string]interface{} `json:"types"`
	Messages  map[string]*rawMessage   `json:"messages"`
}

type rawMessage struct {
	Doc      string                   `json:"doc"`
	Request  []map[string]interface{} `json:"request"`
	Response interface{}              `json:"response"`
	Errors   []interface{}            `json:"errors"`
	OneWay   bool                     `json:"one-way"`
	Stream   string                   `json:"stream"`
}

// ParseProtocol parses an Avro protocol definition (given in JSON format) from
// the io.Reader r.
func ParseProtocol(r io.Reader) (*Protocol, error) {
	const errPrefix = "invalid Avro protocol:"

	// Parse the JSON-formatted Avro protocol
	dec := json.NewDecoder(r)

	var rp rawProtocol
	if err := dec.Decode(&rp); err != nil {
		return nil, errors.New(errPrefix + err.Error())
	}

	// Perform basic validations. Additional validations will be performed
	// during codec construction below.
	if rp.Protocol == "" {
		return nil, errors.New(errPrefix + " missing or empty name attribute")
	}

	if len(rp.Messages) == 0 {
		return nil, errors.New(errPrefix + " missing or empty messages object")
	}

	namedTypes := make(map[string]map[string]interface{}, len(rp.Types))
	for i, typ := range rp.Types {
		name, ok := typ["name"].(string)
		if !ok {
			return nil, fmt.Errorf("%s missing name attribute (type %d)", errPrefix, i+1)
		}

		if _, ok = typ["type"].(string); !ok {
			return nil, fmt.Errorf("%s missing type attribute (type %d)", errPrefix, i+1)
		}

		namedTypes[name] = typ
		namespace, ok := typ["namespace"].(string)
		if ok {
			namedTypes[namespace+"."+name] = typ
		}
	}

	for name, rmsg := range rp.Messages {
		if strings.TrimSpace(name) == "" {
			return nil, errors.New(errPrefix + "missing or empty message name attribute")
		}

		if rmsg.OneWay && (rmsg.Response != "null" || len(rmsg.Errors) > 0) {
			return nil, fmt.Errorf("%s the one-way message parameter may only be true when the response type is \"null\" and no errors are listed (%s)", errPrefix, name)
		}

		switch rmsg.Stream {
		case "":
		case RequestStream, ResponseStream, BidirStream:
			if rmsg.OneWay {
				return nil, fmt.Errorf("%s the stream parameter may not be specified for one-way messages (%s)", errPrefix, name)
			}
		default:
			return nil, fmt.Errorf("%s invalid stream parameter (%s)", errPrefix, name)
		}
	}

	// Construct the protocol with message-specific codecs
	p := &Protocol{
		Namespace: rp.Namespace,
		Protocol:  rp.Protocol,
		Doc:       rp.Doc,
		Messages:  make(map[string]*Message, len(rp.Messages)),
	}

	for name, rmsg := range rp.Messages {
		// Dynamically generate the request schema
		reqSchema, err := json.Marshal(map[string]interface{}{
			"type":   "record",
			"name":   name,
			"doc":    rmsg.Doc,
			"fields": fixupRequestParams(rmsg.Request, namedTypes),
		})
		if err != nil {
			panic(err)
		}

		// Dynamically generate the response schema
		respSchema, err := json.Marshal(fixupSchema(rmsg.Response, namedTypes))
		if err != nil {
			panic(err)
		}

		// Dynamically generate the errors schema
		var (
			errsSchema        []byte
			hasExplicitErrors bool
		)
		if rmsg.Errors != nil {
			if errsSchema, err = json.Marshal(fixupErrors(rmsg.Errors, namedTypes)); err != nil {
				panic(err)
			}
			hasExplicitErrors = len(rmsg.Errors) != 0
		}

		// Construct the message and its codecs
		msg := &Message{
			Doc:               rmsg.Doc,
			RequestSchema:     string(reqSchema),
			ResponseSchema:    string(respSchema),
			ErrorsSchema:      string(errsSchema),
			HasExplicitErrors: hasExplicitErrors,
			OneWay:            rmsg.OneWay,
			Stream:            rmsg.Stream,
		}

		if msg.RequestCodec, err = goavro.NewCodec(msg.RequestSchema); err != nil {
			return nil, fmt.Errorf("%s invalid \"%s\" message request: %s", errPrefix, name, err.Error())
		}

		if msg.ResponseCodec, err = goavro.NewCodec(msg.ResponseSchema); err != nil {
			return nil, fmt.Errorf("%s invalid \"%s\" message response: %s", errPrefix, name, err.Error())
		}

		if errsSchema != nil {
			if msg.ErrorsCodec, err = goavro.NewCodec(msg.ErrorsSchema); err != nil {
				return nil, fmt.Errorf("%s invalid \"%s\" message errors: %s", errPrefix, name, err.Error())
			}
		}

		p.Messages[name] = msg
	}

	return p, nil
}

// fixupRequestParams merges named type definitions into a request schema, which
// is given as a list of parameter schemas.
func fixupRequestParams(params []map[string]interface{}, namedTypes map[string]map[string]interface{}) []map[string]interface{} {
	// A request is a list of named-type parameter schemas, equivalent to an
	// anonymous record, so fixup each parameter schema separately
	for _, param := range params {
		mergeSchema(param, fixupSchema(param["type"], namedTypes))
	}
	return params
}

// fixupErrors merges named type definitions into an errors schema, which is
// given as a union.
//
// Note that gRPC doesn't support arbitrarily-structured errors so this is
// surfaced for completeness/reference only.
func fixupErrors(errors []interface{}, namedTypes map[string]map[string]interface{}) interface{} {
	// Ensure that the errors union contains "string"
	var hasString bool
	for i := range errors {
		if errors[i].(string) == "string" {
			hasString = true
			break
		}
	}
	if !hasString {
		errors = append([]interface{}{"string"}, errors...)
	}
	return fixupSchema(errors, namedTypes)
}

// fixupSchema merges named type definitions into a schema.
func fixupSchema(schema interface{}, namedTypes map[string]map[string]interface{}) interface{} {
	switch x0 := schema.(type) {
	case string:
		// This is a type name. If it matches a named type, then recurse
		// on that type's schema. Otherwise, it is a primitive type and
		// return as-is.
		if typ, ok := namedTypes[x0]; ok {
			return fixupSchema(typ, namedTypes)
		}
		return x0

	case map[string]interface{}:
		switch x1 := x0["type"].(type) {
		case string:
			switch x1 {
			case "error":
				x0["type"] = "record"
				fallthrough

			case "record":
				// This is a record, so recurse on each field
				fields, _ := x0["fields"].([]interface{})
				for i, field := range fields {
					fields[i] = fixupSchema(field, namedTypes)
				}

			case "array":
				// This is a array, so recurse on item schema
				x0["items"] = fixupSchema(x0["items"], namedTypes)

			case "map":
				// This is a map, so recurse on value schema
				x0["values"] = fixupSchema(x0["values"], namedTypes)

			default:
				// This is another named type
				x0 = mergeSchema(x0, fixupSchema(x0["type"], namedTypes))
			}

		case []interface{}:
			// This is a union, so recurse on each element
			for i := range x1 {
				x1[i] = fixupSchema(x1[i], namedTypes)
			}

		default:
			panic(fmt.Sprintf("unhandled type: %T", x1))
		}
		return x0

	case []interface{}:
		// This is a union, so recurse on each element
		for i := range x0 {
			x0[i] = fixupSchema(x0[i], namedTypes)
		}
		return x0

	default:
		panic(fmt.Sprintf("unhandled type: %T", x0))
	}
}

// mergeSchema intelligently merges one schema ("other") into a named-type
// schema ("schema"). If "other" is itself a named-type schema, then the merge
// happens attribute-wise. Otherwise, the type attribute of "schema" is replaced
// with "other".
func mergeSchema(schema map[string]interface{}, other interface{}) map[string]interface{} {
	namedType, ok := other.(map[string]interface{})
	if !ok {
		schema["type"] = other
		return schema
	}

	for k, v := range namedType {
		if k == "name" {
			continue
		}
		schema[k] = v
	}

	return schema
}

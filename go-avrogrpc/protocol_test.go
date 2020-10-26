package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	// ref: https://github.com/apache/avro/blob/master/share/test/schemas/simple.avpr
	exampleSimpleProtocol = `
{
  "namespace": "org.apache.avro.test",
  "protocol": "Simple",
  "doc": "Protocol used for testing.",
  "version": "1.6.2",
  "javaAnnotation": [
    "javax.annotation.Generated(\"avro\")",
    "org.apache.avro.TestAnnotation"
  ],
  "types": [
    {
      "name": "Kind",
      "type": "enum",
      "symbols": [
        "FOO",
        "BAR",
        "BAZ"
      ],
      "javaAnnotation": "org.apache.avro.TestAnnotation"
    },
    {
      "name": "MD5",
      "type": "fixed",
      "size": 16,
      "javaAnnotation": "org.apache.avro.TestAnnotation"
    },
    {
      "name": "TestRecord",
      "type": "record",
      "javaAnnotation": "org.apache.avro.TestAnnotation",
      "fields": [
        {
          "name": "name",
          "type": "string",
          "order": "ignore",
          "javaAnnotation": "org.apache.avro.TestAnnotation"
        },
        {
          "name": "kind",
          "type": "Kind",
          "order": "descending"
        },
        {
          "name": "hash",
          "type": "MD5"
        }
      ]
    },
    {
      "name": "TestError",
      "type": "error",
      "fields": [
        {
          "name": "message",
          "type": "string"
        }
      ]
    },
    {
      "name": "TestRecordWithUnion",
      "type": "record",
      "fields": [
        {
          "name": "kind",
          "type": [
            "null",
            "Kind"
          ]
        },
        {
          "name": "value",
          "type": [
            "null",
            "string"
          ]
        }
      ]
    }
  ],
  "messages": {
    "hello": {
      "doc": "Send a greeting",
      "request": [
        {
          "name": "greeting",
          "type": "string",
          "aliases": [
            "salute"
          ],
          "customProp": "customValue"
        }
      ],
      "response": "string"
    },
    "echo": {
      "doc": "Pretend you're in a cave!",
      "request": [
        {
          "name": "record",
          "type": "TestRecord"
        }
      ],
      "response": "TestRecord"
    },
    "add": {
      "specialProp": "test",
      "request": [
        {
          "name": "arg1",
          "type": "int"
        },
        {
          "name": "arg2",
          "type": "int"
        }
      ],
      "response": "int"
    },
    "echoBytes": {
      "request": [
        {
          "name": "data",
          "type": "bytes"
        }
      ],
      "response": "bytes"
    },
    "error": {
      "doc": "Always throws an error.",
      "request": [],
      "response": "null",
      "errors": [
        "TestError"
      ]
    },
    "ack": {
      "doc": "Send a one way message",
      "request": [],
      "response": "null",
      "one-way": true,
      "javaAnnotation": "org.apache.avro.TestAnnotation"
    }
  }
}`

	// This is a paste of the simple protocol above, but using
	// "TestRecordWithUnion" instead of "TestRecord"
	exampleComplexProtocol = `
{
  "namespace": "org.apache.avro.test",
  "protocol": "Simple",
  "doc": "Protocol used for testing.",
  "version": "1.6.2",
  "javaAnnotation": [
    "javax.annotation.Generated(\"avro\")",
    "org.apache.avro.TestAnnotation"
  ],
  "types": [
    {
      "name": "Kind",
      "type": "enum",
      "symbols": [
        "FOO",
        "BAR",
        "BAZ"
      ],
      "javaAnnotation": "org.apache.avro.TestAnnotation"
    },
    {
      "name": "MD5",
      "type": "fixed",
      "size": 16,
      "javaAnnotation": "org.apache.avro.TestAnnotation"
    },
    {
      "name": "TestRecord",
      "type": "record",
      "javaAnnotation": "org.apache.avro.TestAnnotation",
      "fields": [
        {
          "name": "name",
          "type": "string",
          "order": "ignore",
          "javaAnnotation": "org.apache.avro.TestAnnotation"
        },
        {
          "name": "kind",
          "type": "Kind",
          "order": "descending"
        },
        {
          "name": "hash",
          "type": "MD5"
        }
      ]
    },
    {
      "name": "TestError",
      "type": "error",
      "fields": [
        {
          "name": "message",
          "type": "string"
        }
      ]
    },
    {
      "name": "TestRecordWithUnion",
      "type": "record",
      "fields": [
        {
          "name": "kind",
          "type": [
            "null",
            "Kind"
          ]
        },
        {
          "name": "value",
          "type": [
            "null",
            "string"
          ]
        }
      ]
    }
  ],
  "messages": {
    "hello": {
      "doc": "Send a greeting",
      "request": [
        {
          "name": "greeting",
          "type": "string",
          "aliases": [
            "salute"
          ],
          "customProp": "customValue"
        }
      ],
      "response": "string"
    },
    "echo": {
      "doc": "Pretend you're in a cave!",
      "request": [
        {
          "name": "record",
          "type": "TestRecordWithUnion"
        }
      ],
      "response": "TestRecordWithUnion"
    },
    "add": {
      "specialProp": "test",
      "request": [
        {
          "name": "arg1",
          "type": "int"
        },
        {
          "name": "arg2",
          "type": "int"
        }
      ],
      "response": "int"
    },
    "echoBytes": {
      "request": [
        {
          "name": "data",
          "type": "bytes"
        }
      ],
      "response": "bytes"
    },
    "error": {
      "doc": "Always throws an error.",
      "request": [],
      "response": "null",
      "errors": [
        "TestError"
      ]
    },
    "ack": {
      "doc": "Send a one way message",
      "request": [],
      "response": "null",
      "one-way": true,
      "javaAnnotation": "org.apache.avro.TestAnnotation"
    }
  }
}`

	// ref: https://github.com/apache/avro/blob/master/share/test/schemas/namespace.avpr
	exampleNamespaceProtocol = `
{
  "namespace": "org.apache.avro.test.namespace",
  "protocol": "TestNamespace",
  "types": [
    {
      "name": "org.apache.avro.test.util.MD5",
      "type": "fixed",
      "size": 16
    },
    {
      "name": "TestRecord",
      "type": "record",
      "fields": [
        {
          "name": "hash",
          "type": "org.apache.avro.test.util.MD5"
        }
      ]
    },
    {
      "name": "TestError",
      "namespace": "org.apache.avro.test.errors",
      "type": "error",
      "fields": [
        {
          "name": "message",
          "type": "string"
        }
      ]
    }
  ],
  "messages": {
    "echo": {
      "request": [
        {
          "name": "record",
          "type": "TestRecord"
        }
      ],
      "response": "TestRecord"
    },
    "error": {
      "request": [],
      "response": "null",
      "errors": [
        "org.apache.avro.test.errors.TestError"
      ]
    }
  }
}`
)

func TestParseSimpleProtocol(t *testing.T) {
	p, err := ParseProtocol(bytes.NewBufferString(exampleSimpleProtocol))
	require.NoError(t, err)
	require.NotNil(t, p)
}

func TestParseComplexProtocol(t *testing.T) {
	p, err := ParseProtocol(bytes.NewBufferString(exampleComplexProtocol))
	require.NoError(t, err)
	require.NotNil(t, p)
}

func TestParseNamespaceProtocol(t *testing.T) {
	p, err := ParseProtocol(bytes.NewBufferString(exampleNamespaceProtocol))
	require.NoError(t, err)
	require.NotNil(t, p)
}

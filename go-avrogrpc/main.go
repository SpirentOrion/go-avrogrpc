//go:generate go-bindata -fs=false -nomemcopy -o codegen.go codegen.template

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
	"unicode"
)

var (
	inputPath  = flag.String("i", "", "path to Avro protocol (.avpr) file")
	outputPath = flag.String("o", "", "path to generated code (.go) file")
	pkgName    = flag.String("p", "", "package name for generated code")
	tmplPath   = flag.String("t", "", "path to code generation template")
)

func usage() {
	const msg = `usage: %s [flags]

Generates a Go-language gRPC client and service implementation from an Avro
protocol file. Input is an Avro protocol definition, read from the given file
or stdin (default). Note that although this is a JSON-formatted file, it is a
different format than an ordinary Avro schema file. See
https://avro.apache.org/docs/current/spec.html#Protocol+Declaration for the
specification. Output is written to the given file or stdout (default). The
code generation template (in Go template form) may be optionally overridden,
otherwise the built-in template is used. If no package name is given, then
the Avro protocol name is used as the Go package name.

Flags:

`
	fmt.Fprintf(os.Stderr, msg, os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	// Parse the Avro protocol definition
	var (
		proto    *Protocol
		metadata string
		err      error
	)
	if s := *inputPath; s != "" {
		var f *os.File
		if f, err = os.Open(s); err != nil {
			log.Fatal(err)
		}
		proto, err = ParseProtocol(f)
		metadata = s
		f.Close()
	} else {
		proto, err = ParseProtocol(os.Stdin)
	}
	if err != nil {
		log.Fatal(err)
	}

	for name, msg := range proto.Messages {
		if msg.HasExplicitErrors {
			log.Printf("gRPC doesn't support arbitrarily-structured errors, so the errors schema for \"%s\" is unused", name)
		}
	}

	// Parse the code generation template
	tmplFuncMap := template.FuncMap{
		"camelBack": camelBack,
		"camelCase": camelCase,
		"inc":       func(n int) int { return n + 1 },
	}

	tmpl := template.New("codegen").Funcs(tmplFuncMap)
	if s := *tmplPath; s != "" {
		var buf []byte
		if buf, err = ioutil.ReadFile(s); err != nil {
			log.Fatal(err)
		}
		tmpl, err = tmpl.Parse(string(buf))
	} else {
		tmpl, err = tmpl.Parse(string(MustAsset("codegen.template")))
	}
	if err != nil {
		log.Fatal(err)
	}

	// Open the output file
	var w io.WriteCloser
	if s := *outputPath; s != "" {
		if w, err = os.Create(s); err != nil {
			log.Fatal(err)
		}
	} else {
		w = os.Stdout
	}

	// Generate code
	if *pkgName == "" {
		s := camelBack(proto.Protocol)
		pkgName = &s
	}

	data := struct {
		Package  string
		Protocol *Protocol
		Metadata string
	}{
		Package:  *pkgName,
		Protocol: proto,
		Metadata: metadata,
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, &data); err != nil {
		log.Fatal(err)
	}

	// Format and output the code
	source, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	if _, err = io.Copy(w, bytes.NewBuffer(source)); err != nil {
		log.Fatal(err)
	}
}

func camelBack(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	return string(append([]rune{unicode.ToLower(runes[0])}, []rune(runes[1:])...))
}

func camelCase(s string) string {
	return strings.Title(s)
}

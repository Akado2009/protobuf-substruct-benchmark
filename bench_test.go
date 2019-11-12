package main

import (
	"testing"
	// anotherinner "github.com/Akado2009/protobuf-substruct-benchmark/protobuf-struct/anotherinner"
	// inner "github.com/Akado2009/protobuf-substruct-benchmark/protobuf-struct/inner"
	// anotherinner "github.com/Akado2009/protobuf-substruct-benchmark/protobuf-struct/anotherinner"
)

var sasds []string

const (
	constName = "AnyRandomStructure"
)

func init() {

}

func BenchmarkAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		name := "AnyRandomStructure"
		if name != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, name)
		}
	}
}

func BenchmarkByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		name := "AnyRandomStructure"
		if name != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, name)
		}
	}
}

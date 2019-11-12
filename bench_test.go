package main

import "testing"

var sasds []string

const (
	constName = "AnyRandomStructure"
)

func init() {

}

func BenchmarkAny(b *testing.B) {
	name := "AnyRandomStructure"
	if name != constName {
		b.Fatalf("Error: expcted name %s, but got %s", constName, name)
	}
}

func BenchmarkByte(b *testing.B) {
	name := "AnyRandomStructure"
	if name != constName {
		b.Fatalf("Error: expcted name %s, but got %s", constName, name)
	}
}

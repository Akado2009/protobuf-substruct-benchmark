package main

import (
	"log"
	"testing"

	AnotherInnerStruct "github.com/Akado2009/protobuf-substruct-benchmark/generated-code/protobuf-struct/anotherinner"
	UniversalParser "github.com/Akado2009/protobuf-substruct-benchmark/generated-code/protobuf-struct/universal"
	InnerStruct "github.com/Akado2009/protobuf-substruct-benchmark/protobuf-struct/generated-code/inner"
	"github.com/golang/protobuf/proto"
)

const (
	constName = "AnyRandomStructure"
)

func BenchmarkAnySmallInner(b *testing.B) {
	inner := InnerStruct.InnerMessage{
		Name:       "randomName",
		Id:         int32(100),
		SecondName: "secondRandomName",
	}
	data, err := proto.Marshal(inner)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	wrapper := UniversalParser.UniversalMessage{
		Name: constName,
		Msg:  data,
	}
	for i := 0; i < b.N; i++ {
		parser := UniversalParser.UniversalMessage{}
		err = proto.Unmarshal(data, wrapper)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if parser.Name != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, parser.Name)
		}
	}
}

func BenchmarkByteSmallInner(b *testing.B) {
	inner := InnerStruct.InnerMessage{
		Name:       "randomName",
		Id:         int32(100),
		SecondName: "secondRandomName",
	}
	data, err := proto.Marshal(inner)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	wrapper := UniversalParser.UniversalMessage{
		Name: constName,
		Msg:  data,
	}
	for i := 0; i < b.N; i++ {
		parser := UniversalParser.UniversalMessage{}
		err = proto.Unmarshal(data, wrapper)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if parser.Name != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, parser.Name)
		}
	}
}

func BenchmarkAnyLargeInner(b *testing.B) {
	aInner := AnotherInnerStruct.AnotherInnerMessage{
		Name:       "randomName",
		Id:         int32(100),
		SecondName: "secondRandomName",
		ThirdName:  "thirdRandomName",
		IdFloat:    float32(100.02),
		FourthName: "fourthRandomName",
		FifthName:  "fifthRandomName",
	}
	data, err := proto.Marshal(aInner)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	wrapper := UniversalParser.UniversalMessage{
		Name: constName,
		Msg:  data,
	}
	for i := 0; i < b.N; i++ {
		parser := UniversalParser.UniversalMessage{}
		err = proto.Unmarshal(data, wrapper)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if parser.Name != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, parser.Name)
		}
	}
}

func BenchmarkByteLargeInner(b *testing.B) {
	aInner := AnotherInnerStruct.AnotherInnerMessage{
		Name:       "randomName",
		Id:         int32(100),
		SecondName: "secondRandomName",
		ThirdName:  "thirdRandomName",
		IdFloat:    float32(100.02),
		FourthName: "fourthRandomName",
		FifthName:  "fifthRandomName",
	}
	data, err := proto.Marshal(aInner)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	wrapper := UniversalParser.UniversalMessage{
		Name: constName,
		Msg:  data,
	}
	for i := 0; i < b.N; i++ {
		parser := UniversalParser.UniversalMessage{}
		err = proto.Unmarshal(data, wrapper)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if parser.Name != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, parser.Name)
		}
	}
}

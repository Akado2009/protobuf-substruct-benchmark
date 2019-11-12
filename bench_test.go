package main

import (
	"bytes"
	"log"
	"testing"

	AnotherInnerStruct "github.com/Akado2009/protobuf-substruct-benchmark/protobuf-struct/generated-code/anotherinner"
	InnerStruct "github.com/Akado2009/protobuf-substruct-benchmark/protobuf-struct/generated-code/inner"
	UniversalParser "github.com/Akado2009/protobuf-substruct-benchmark/protobuf-struct/generated-code/universal"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

const (
	constName  = "AnyRandomStructure"
	bufferSize = 20
)

func BenchmarkAnySmallInnerWoAllocation(b *testing.B) {
	inner := &InnerStruct.InnerMessage{
		Name:       "randomName",
		Id:         int32(100),
		SecondName: "secondRandomName",
	}
	msg, err := ptypes.MarshalAny(inner)
	wrapper := UniversalParser.UniversalMessage{
		Name: constName,
		Msg:  msg,
	}
	data, err := proto.Marshal(&wrapper)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	for i := 0; i < b.N; i++ {
		parser := UniversalParser.UniversalMessage{}
		err = proto.Unmarshal(data, &parser)

		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		newInner := InnerStruct.InnerMessage{}
		err = proto.Unmarshal(parser.Msg.Value, &newInner)
		// err := ptypes.UnmarshalAny(parser.Msg, &newInner)

		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if parser.Name != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, parser.Name)
		}
	}
}

func BenchmarkByteSmallInnerWoAllocation(b *testing.B) {
	inner := InnerStruct.InnerMessage{
		Name:       "randomName",
		Id:         int32(100),
		SecondName: "secondRandomName",
	}
	data, err := proto.Marshal(&inner)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	buf := make([]byte, bufferSize, bufferSize)
	copy(buf[:], constName)
	overallData := append(buf, data...)
	for i := 0; i < b.N; i++ {
		structName := string(bytes.Trim(overallData[:bufferSize], "\x00")[:])
		structSelf := overallData[bufferSize:]
		newInner := InnerStruct.InnerMessage{}
		err = proto.Unmarshal(structSelf, &newInner)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if structName != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, structName)
		}
	}
}

func BenchmarkAnySmallInnerWAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inner := &InnerStruct.InnerMessage{
			Name:       "randomName",
			Id:         int32(100),
			SecondName: "secondRandomName",
		}
		msg, err := ptypes.MarshalAny(inner)
		wrapper := UniversalParser.UniversalMessage{
			Name: constName,
			Msg:  msg,
		}
		data, err := proto.Marshal(&wrapper)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		parser := UniversalParser.UniversalMessage{}
		err = proto.Unmarshal(data, &parser)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		newInner := InnerStruct.InnerMessage{}
		err = proto.Unmarshal(parser.Msg.Value, &newInner)
		// err = ptypes.UnmarshalAny(parser.Msg, &newInner)

		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		if parser.Name != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, parser.Name)
		}
	}
}

func BenchmarkByteSmallInnerWAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inner := InnerStruct.InnerMessage{
			Name:       "randomName",
			Id:         int32(100),
			SecondName: "secondRandomName",
		}
		data, err := proto.Marshal(&inner)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		buf := make([]byte, bufferSize, bufferSize)
		copy(buf[:], constName)
		overallData := append(buf, data...)
		structName := string(bytes.Trim(overallData[:bufferSize], "\x00")[:])
		structSelf := overallData[bufferSize:]
		newInner := InnerStruct.InnerMessage{}
		err = proto.Unmarshal(structSelf, &newInner)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if structName != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, structName)
		}
	}
}

func BenchmarkAnyLargeInnerWoAllocation(b *testing.B) {
	aInner := AnotherInnerStruct.AnotherInnerMessage{
		Name:       "randomName",
		Id:         int32(100),
		SecondName: "secondRandomName",
		ThirdName:  "thirdRandomName",
		IdFloat:    float32(100.02),
		FourthName: "fourthRandomName",
		FifthName:  "fifthRandomName",
	}
	msg, err := ptypes.MarshalAny(&aInner)
	wrapper := UniversalParser.UniversalMessage{
		Name: constName,
		Msg:  msg,
	}
	data, err := proto.Marshal(&wrapper)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	for i := 0; i < b.N; i++ {
		parser := UniversalParser.UniversalMessage{}
		err = proto.Unmarshal(data, &parser)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		newAInner := AnotherInnerStruct.AnotherInnerMessage{}
		err = proto.Unmarshal(parser.Msg.Value, &newAInner)

		// err = ptypes.UnmarshalAny(parser.Msg, &newAInner)

		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if parser.Name != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, parser.Name)
		}
	}
}

func BenchmarkByteLargeInnerWoAllocation(b *testing.B) {
	aInner := AnotherInnerStruct.AnotherInnerMessage{
		Name:       "randomName",
		Id:         int32(100),
		SecondName: "secondRandomName",
		ThirdName:  "thirdRandomName",
		IdFloat:    float32(100.02),
		FourthName: "fourthRandomName",
		FifthName:  "fifthRandomName",
	}
	data, err := proto.Marshal(&aInner)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	buf := make([]byte, bufferSize, bufferSize)
	copy(buf[:], constName)
	overallData := append(buf, data...)
	for i := 0; i < b.N; i++ {
		structName := string(bytes.Trim(overallData[:bufferSize], "\x00")[:])
		structSelf := overallData[bufferSize:]
		newInner := InnerStruct.InnerMessage{}
		err = proto.Unmarshal(structSelf, &newInner)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if structName != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, structName)
		}
	}
}

func BenchmarkAnyLargeInnerWAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aInner := AnotherInnerStruct.AnotherInnerMessage{
			Name:       "randomName",
			Id:         int32(100),
			SecondName: "secondRandomName",
			ThirdName:  "thirdRandomName",
			IdFloat:    float32(100.02),
			FourthName: "fourthRandomName",
			FifthName:  "fifthRandomName",
		}
		msg, err := ptypes.MarshalAny(&aInner)
		wrapper := UniversalParser.UniversalMessage{
			Name: constName,
			Msg:  msg,
		}
		data, err := proto.Marshal(&wrapper)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		parser := UniversalParser.UniversalMessage{}
		err = proto.Unmarshal(data, &parser)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		newAInner := AnotherInnerStruct.AnotherInnerMessage{}
		err = proto.Unmarshal(parser.Msg.Value, &newAInner)

		// err = ptypes.UnmarshalAny(parser.Msg, &newAInner)

		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if parser.Name != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, parser.Name)
		}
	}
}

func BenchmarkByteLargeInnerWAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aInner := AnotherInnerStruct.AnotherInnerMessage{
			Name:       "randomName",
			Id:         int32(100),
			SecondName: "secondRandomName",
			ThirdName:  "thirdRandomName",
			IdFloat:    float32(100.02),
			FourthName: "fourthRandomName",
			FifthName:  "fifthRandomName",
		}
		data, err := proto.Marshal(&aInner)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		buf := make([]byte, bufferSize, bufferSize)
		copy(buf[:], constName)
		overallData := append(buf, data...)
		structName := string(bytes.TrimRight(overallData[:bufferSize], "\x00")[:])
		structSelf := overallData[bufferSize:]
		newInner := InnerStruct.InnerMessage{}
		err = proto.Unmarshal(structSelf, &newInner)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if structName != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, structName)
		}
	}
}

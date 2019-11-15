package main

import (
	"log"
	"testing"

	AnotherInnerStruct "github.com/Akado2009/protobuf-substruct-benchmark/protobuf-struct/generated-code/anotherinner"
	InnerStruct "github.com/Akado2009/protobuf-substruct-benchmark/protobuf-struct/generated-code/inner"
	UniversalParser "github.com/Akado2009/protobuf-substruct-benchmark/protobuf-struct/generated-code/universal"

	firstmessage "github.com/Akado2009/protobuf-substruct-benchmark/test-proto/firstmessage"
	general "github.com/Akado2009/protobuf-substruct-benchmark/test-proto/general"
	generaloneof "github.com/Akado2009/protobuf-substruct-benchmark/test-proto/generaloneof"

	HeaderStruct "github.com/Akado2009/protobuf-substruct-benchmark/protobuf-struct/generated-code/header"
	secondmessage "github.com/Akado2009/protobuf-substruct-benchmark/test-proto/secondmessage"

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
		structName := string(overallData[:bufferSize])
		structSelf := overallData[bufferSize:]
		newInner := InnerStruct.InnerMessage{}
		err = proto.Unmarshal(structSelf, &newInner)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if structName != constName {
			//			b.Fatalf("Error: expcted name %s, but got %s", constName, structName)
		}
	}
}

func BenchmarkGeneralSmallInnerWoAllocation(b *testing.B) {
	inner := firstmessage.FirstMessage{
		Name:       constName,
		Id:         int32(100),
		SecondName: "secondRandomName",
	}
	generalStruct := general.General{
		Fmsg: &inner,
	}
	data, err := proto.Marshal(&generalStruct)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	for i := 0; i < b.N; i++ {
		parser := general.General{}
		err = proto.Unmarshal(data, &parser)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if parser.Msg10 != nil {
		}
		if parser.Msg9 != nil {
		}
		if parser.Msg8 != nil {
		}
		if parser.Msg7 != nil {
		}
		if parser.Msg6 != nil {
		}
		if parser.Msg5 != nil {
		}
		if parser.Msg4 != nil {
		}
		if parser.Msg3 != nil {
		}
		if parser.Msg2 != nil {
		}
		if parser.Msg1 != nil {
		}
		if parser.Smsg != nil {
		}
		if parser.Fmsg != nil {
			if parser.Fmsg.Name != constName {
				b.Fatalf("Error: expcted name %s, but got %s", constName, parser.Fmsg.Name)
			}
		}
	}
}

func BenchmarkHeaderSmallInnerWoAllocation(b *testing.B) {
	inner := firstmessage.FirstMessage{
		Name:       constName,
		Id:         int32(100),
		SecondName: "secondRandomName",
	}
	data, err := proto.Marshal(&inner)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	hS := HeaderStruct.HeaderMessager{
		Type:    constName,
		Message: data,
	}
	headerData, err := proto.Marshal(&hS)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	for i := 0; i < b.N; i++ {
		headerParser := HeaderStruct.HeaderMessager{}
		err = proto.Unmarshal(headerData, &headerParser)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		headerMessage := firstmessage.FirstMessage{}
		err = proto.Unmarshal(headerParser.Message, &headerMessage)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if headerParser.Type != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, headerParser.Type)
		}
	}
}

func BenchmarkOneOfSmallInnerWoAllocation(b *testing.B) {
	inner := firstmessage.FirstMessage{
		Name:       constName,
		Id:         int32(100),
		SecondName: "secondRandomName",
	}

	hS := generaloneof.GeneralOneOf{
		Msg: &generaloneof.GeneralOneOf_Fmsg{
			Fmsg: &inner,
		},
	}

	headerData, err := proto.Marshal(&hS)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	for i := 0; i < b.N; i++ {
		headerParser := generaloneof.GeneralOneOf{}
		err = proto.Unmarshal(headerData, &headerParser)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		// _ := secondmessage.SecondMessage{}
		switch x := headerParser.Msg.(type) {
		case *generaloneof.GeneralOneOf_Msg1:
			continue
		case *generaloneof.GeneralOneOf_Msg2:
			continue
		case *generaloneof.GeneralOneOf_Msg3:
			continue
		case *generaloneof.GeneralOneOf_Msg4:
			continue
		case *generaloneof.GeneralOneOf_Msg5:
			continue
		case *generaloneof.GeneralOneOf_Msg6:
			continue
		case *generaloneof.GeneralOneOf_Msg7:
			continue
		case *generaloneof.GeneralOneOf_Msg8:
			continue
		case *generaloneof.GeneralOneOf_Msg9:
			continue
		case *generaloneof.GeneralOneOf_Msg10:
			continue
		case *generaloneof.GeneralOneOf_Fmsg:
			if x.Fmsg.Name != constName {
				b.Fatalf("Error: expcted name %s, but got %s", constName, x.Fmsg.Name)
			}
		case *generaloneof.GeneralOneOf_Smsg:
			continue
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
		// structName := string(bytes.TrimRight(overallData[:bufferSize], "\x00")[:])
		structName := string(overallData[:bufferSize])
		structSelf := overallData[bufferSize:]
		newInner := InnerStruct.InnerMessage{}
		err = proto.Unmarshal(structSelf, &newInner)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		if structName != constName {
			//			b.Fatalf("Error: expcted name %s, but got %s", constName, structName)
		}
	}
}

func BenchmarkGeneralLargeInnerWoAllocation(b *testing.B) {
	aInner := secondmessage.SecondMessage{
		Name:       constName,
		Id:         int32(100),
		SecondName: "secondRandomName",
		ThirdName:  "thirdRandomName",
		IdFloat:    float32(100.02),
		FourthName: "fourthRandomName",
		FifthName:  "fifthRandomName",
	}
	generalStruct := general.General{
		Smsg: &aInner,
	}
	data, err := proto.Marshal(&generalStruct)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	for i := 0; i < b.N; i++ {
		parser := general.General{}
		err = proto.Unmarshal(data, &parser)

		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if parser.Msg10 != nil {
		}
		if parser.Msg9 != nil {
		}
		if parser.Msg8 != nil {
		}
		if parser.Msg7 != nil {
		}
		if parser.Msg6 != nil {
		}
		if parser.Msg5 != nil {
		}
		if parser.Msg4 != nil {
		}
		if parser.Msg3 != nil {
		}
		if parser.Msg2 != nil {
		}
		if parser.Msg1 != nil {
		}
		if parser.Fmsg != nil {
		}
		if parser.Smsg != nil {
			if parser.Smsg.Name != constName {
				b.Fatalf("Error: expcted name %s, but got %s", constName, parser.Smsg.Name)
			}
		}
	}
}

func BenchmarkHeaderLargeInnerWoAllocation(b *testing.B) {
	aInner := secondmessage.SecondMessage{
		Name:       constName,
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

	hS := HeaderStruct.HeaderMessager{
		Type:    constName,
		Message: data,
	}
	headerData, err := proto.Marshal(&hS)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	for i := 0; i < b.N; i++ {
		headerParser := HeaderStruct.HeaderMessager{}
		err = proto.Unmarshal(headerData, &headerParser)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		headerMessage := secondmessage.SecondMessage{}
		err = proto.Unmarshal(headerParser.Message, &headerMessage)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		if headerParser.Type != constName {
			b.Fatalf("Error: expcted name %s, but got %s", constName, headerParser.Type)
		}
	}
}

func BenchmarkOneOfLargeInnerWoAllocation(b *testing.B) {
	aInner := &secondmessage.SecondMessage{
		Name:       constName,
		Id:         int32(100),
		SecondName: "secondRandomName",
		ThirdName:  "thirdRandomName",
		IdFloat:    float32(100.02),
		FourthName: "fourthRandomName",
		FifthName:  "fifthRandomName",
	}

	hS := generaloneof.GeneralOneOf{
		Msg: &generaloneof.GeneralOneOf_Smsg{
			Smsg: aInner,
		},
	}

	headerData, err := proto.Marshal(&hS)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	for i := 0; i < b.N; i++ {
		headerParser := generaloneof.GeneralOneOf{}
		err = proto.Unmarshal(headerData, &headerParser)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		// _ := secondmessage.SecondMessage{}
		switch x := headerParser.Msg.(type) {
		case *generaloneof.GeneralOneOf_Msg1:
			continue
		case *generaloneof.GeneralOneOf_Msg2:
			continue
		case *generaloneof.GeneralOneOf_Msg3:
			continue
		case *generaloneof.GeneralOneOf_Msg4:
			continue
		case *generaloneof.GeneralOneOf_Msg5:
			continue
		case *generaloneof.GeneralOneOf_Msg6:
			continue
		case *generaloneof.GeneralOneOf_Msg7:
			continue
		case *generaloneof.GeneralOneOf_Msg8:
			continue
		case *generaloneof.GeneralOneOf_Msg9:
			continue
		case *generaloneof.GeneralOneOf_Msg10:
			continue
		case *generaloneof.GeneralOneOf_Fmsg:
			continue
		case *generaloneof.GeneralOneOf_Smsg:
			if x.Smsg.Name != constName {
				b.Fatalf("Error: expcted name %s, but got %s", constName, x.Smsg.Name)
			}

		}

	}
}

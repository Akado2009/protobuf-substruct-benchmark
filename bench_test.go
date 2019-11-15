package main

import (
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"

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
	constName   = "AnyRandomStructure"
	bufferSize  = 20
	eventsCount = 100
	stringLimit = 20
	intLimit    = 100000
	nameBufferS = 20
	vBufferS    = 2
	sizeBufferS = 5
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randomInt(n int) int {
	return rand.Intn(n)
}

func generateRandomFirst() *firstmessage.FirstMessage {
	inner := &firstmessage.FirstMessage{
		Name:       randomString(stringLimit),
		Id:         int32(randomInt(intLimit)),
		SecondName: randomString(stringLimit),
	}
	return inner
}

func generateRandomSecond() *secondmessage.SecondMessage {
	aInner := &secondmessage.SecondMessage{
		Name:       randomString(stringLimit),
		Id:         int32(randomInt(intLimit)),
		SecondName: randomString(stringLimit),
		ThirdName:  randomString(stringLimit),
		IdFloat:    float32(randomInt(intLimit) / 3),
		FourthName: randomString(stringLimit),
		FifthName:  randomString(stringLimit),
	}
	return aInner
}

func generateWrappedBytes(key string) []byte {
	var data []byte
	var err error
	if key == "first" {
		data, err = proto.Marshal(generateRandomFirst())
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
	} else {
		data, err = proto.Marshal(generateRandomSecond())
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
	}
	size := len(data)
	version := "1"

	nameBuffer := make([]byte, nameBufferS, nameBufferS)
	copy(nameBuffer[:], constName)
	vBuffer := make([]byte, vBufferS, vBufferS)
	copy(vBuffer[:], version)
	sizeBuffer := make([]byte, sizeBufferS, sizeBufferS)
	copy(sizeBuffer[:], strconv.Itoa(size))
	overallData := append(nameBuffer, vBuffer...)
	overallData = append(overallData, sizeBuffer...)
	overallData = append(overallData, data...)

	return overallData
}

func generateWrappedArrays(key string) []byte {
	var data []byte
	var err error
	if key == "first" {
		data, err = proto.Marshal(generateRandomFirst())
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
	} else {
		data, err = proto.Marshal(generateRandomSecond())
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
	}
	return data
}

func generateWrappedGeneral(key string) []byte {
	generalStruct := general.General{}
	if key == "first" {
		generalStruct.Fmsg = generateRandomFirst()
	} else {
		generalStruct.Smsg = generateRandomSecond()
	}
	data, err := proto.Marshal(&generalStruct)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	return data
}

func generateEventsByte() []byte {
	b := make([]byte, 0)
	for i := 0; i < eventsCount/2; i++ {
		b = append(b, generateWrappedBytes("first")...)
		b = append(b, generateWrappedBytes("second")...)
	}
	return b
}

func generateEventsArrays() []byte {
	b := make([]byte, 0)

	for i := 0; i < eventsCount/2; i++ {
		b = append(b, generateWrappedArrays("first")...)
		b = append(b, generateWrappedArrays("second")...)
	}
	return b
}

func generateGeneralArrays() []byte {
	b := make([]byte, 0)

	for i := 0; i < eventsCount/2; i++ {
		b = append(b, generateWrappedGeneral("first")...)
		b = append(b, generateWrappedGeneral("second")...)
	}
	return b
}

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

func BenchmarkGeneralArray(b *testing.B) {
}

func BenchmarkByteArray(b *testing.B) {
}

func BenchmarkArrayArray(b *testing.B) {
}

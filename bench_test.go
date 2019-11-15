package main

import (
	"encoding/binary"
	"log"
	"math/rand"
	"testing"
	"time"

	firstmessage "github.com/Akado2009/protobuf-substruct-benchmark/test-proto/firstmessage"
	general "github.com/Akado2009/protobuf-substruct-benchmark/test-proto/general"
	generalarray "github.com/Akado2009/protobuf-substruct-benchmark/test-proto/generalarray"
	generaloptarray "github.com/Akado2009/protobuf-substruct-benchmark/test-proto/generaloptarray"
	secondmessage "github.com/Akado2009/protobuf-substruct-benchmark/test-proto/secondmessage"

	"github.com/golang/protobuf/proto"
)

const (
	constName   = "AnyRandomStructure"
	bufferSize  = 20
	eventsCount = 1000
	stringLimit = 20
	intLimit    = 100000
	nameBufferS = 20
	vBufferS    = 2
	sizeBufferS = 4
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func convertByteToInt(in []byte) int {
	return (int(in[0])<<24 | int(in[1])<<16 | int(in[2])<<8 | int(in[3]))
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
		Name:       constName,
		Id:         int32(randomInt(intLimit)),
		SecondName: randomString(stringLimit),
	}
	return inner
}

func generateRandomSecond() *secondmessage.SecondMessage {
	aInner := &secondmessage.SecondMessage{
		Name:       constName,
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
	nameBuffer := make([]byte, nameBufferS, nameBufferS)

	if key == "first" {
		data, err = proto.Marshal(generateRandomFirst())
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		copy(nameBuffer[:], "FirstMessage")
	} else {
		data, err = proto.Marshal(generateRandomSecond())
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		copy(nameBuffer[:], "SecondMessage")
	}
	size := len(data)
	version := "1"

	vBuffer := make([]byte, vBufferS, vBufferS)
	copy(vBuffer[:], version)
	sizeBuffer := make([]byte, sizeBufferS, sizeBufferS)
	binary.LittleEndian.PutUint32(sizeBuffer, uint32(size))
	overallData := append(nameBuffer, vBuffer...)
	overallData = append(overallData, sizeBuffer...)
	overallData = append(overallData, data...)

	return overallData
}

func generateWrappedGeneral(key string) *general.General {
	generalStruct := &general.General{}
	if key == "first" {
		generalStruct.Fmsg = generateRandomFirst()
	} else {
		generalStruct.Smsg = generateRandomSecond()
	}
	return generalStruct
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
	goa := generaloptarray.GeneralOptArrays{}
	fmsgs := make([]*firstmessage.FirstMessage, 0)
	smsgs := make([]*secondmessage.SecondMessage, 0)

	for i := 0; i < eventsCount/2; i++ {
		fmsgs = append(fmsgs, generateRandomFirst())
		smsgs = append(smsgs, generateRandomSecond())
	}
	goa.Fmsgs = fmsgs
	goa.Smsgs = smsgs
	data, err := proto.Marshal(&goa)
	if err != nil {
		log.Fatal("marshaling error")
	}
	return data
}

func generateGeneralArrays() []byte {
	gSubArray := make([]*general.General, 0)

	for i := 0; i < eventsCount/2; i++ {
		gSubArray = append(gSubArray, generateWrappedGeneral("first"))
		gSubArray = append(gSubArray, generateWrappedGeneral("second"))
	}
	ga := generalarray.GeneralArray{
		GenArray: gSubArray,
	}
	data, err := proto.Marshal(&ga)
	if err != nil {
		log.Fatal("error unarmshalling")
	}
	return data
}

// func BenchmarkAnySmallInnerWoAllocation(b *testing.B) {
// 	inner := &InnerStruct.InnerMessage{
// 		Name:       "randomName",
// 		Id:         int32(100),
// 		SecondName: "secondRandomName",
// 	}
// 	msg, err := ptypes.MarshalAny(inner)
// 	wrapper := UniversalParser.UniversalMessage{
// 		Name: constName,
// 		Msg:  msg,
// 	}
// 	data, err := proto.Marshal(&wrapper)
// 	if err != nil {
// 		log.Fatal("marshaling error: ", err)
// 	}
// 	for i := 0; i < b.N; i++ {
// 		parser := UniversalParser.UniversalMessage{}
// 		err = proto.Unmarshal(data, &parser)

// 		if err != nil {
// 			log.Fatal("unmarshaling error: ", err)
// 		}

// 		newInner := InnerStruct.InnerMessage{}
// 		err = proto.Unmarshal(parser.Msg.Value, &newInner)
// 		// err := ptypes.UnmarshalAny(parser.Msg, &newInner)

// 		if err != nil {
// 			log.Fatal("unmarshaling error: ", err)
// 		}
// 		if parser.Name != constName {
// 			b.Fatalf("Error: expcted name %s, but got %s", constName, parser.Name)
// 		}
// 	}
// }

// func BenchmarkByteSmallInnerWoAllocation(b *testing.B) {
// 	inner := InnerStruct.InnerMessage{
// 		Name:       "randomName",
// 		Id:         int32(100),
// 		SecondName: "secondRandomName",
// 	}
// 	data, err := proto.Marshal(&inner)
// 	if err != nil {
// 		log.Fatal("marshaling error: ", err)
// 	}
// 	buf := make([]byte, bufferSize, bufferSize)
// 	copy(buf[:], constName)
// 	overallData := append(buf, data...)
// 	for i := 0; i < b.N; i++ {
// 		structName := string(overallData[:bufferSize])
// 		structSelf := overallData[bufferSize:]
// 		newInner := InnerStruct.InnerMessage{}
// 		err = proto.Unmarshal(structSelf, &newInner)
// 		if err != nil {
// 			log.Fatal("unmarshaling error: ", err)
// 		}
// 		if structName != constName {
// 			//			b.Fatalf("Error: expcted name %s, but got %s", constName, structName)
// 		}
// 	}
// }

// func BenchmarkGeneralSmallInnerWoAllocation(b *testing.B) {
// 	inner := firstmessage.FirstMessage{
// 		Name:       constName,
// 		Id:         int32(100),
// 		SecondName: "secondRandomName",
// 	}
// 	generalStruct := general.General{
// 		Fmsg: &inner,
// 	}
// 	data, err := proto.Marshal(&generalStruct)
// 	if err != nil {
// 		log.Fatal("marshaling error: ", err)
// 	}
// 	for i := 0; i < b.N; i++ {
// 		parser := general.General{}
// 		err = proto.Unmarshal(data, &parser)
// 		if err != nil {
// 			log.Fatal("unmarshaling error: ", err)
// 		}
// 		if parser.Msg10 != nil {
// 		}
// 		if parser.Msg9 != nil {
// 		}
// 		if parser.Msg8 != nil {
// 		}
// 		if parser.Msg7 != nil {
// 		}
// 		if parser.Msg6 != nil {
// 		}
// 		if parser.Msg5 != nil {
// 		}
// 		if parser.Msg4 != nil {
// 		}
// 		if parser.Msg3 != nil {
// 		}
// 		if parser.Msg2 != nil {
// 		}
// 		if parser.Msg1 != nil {
// 		}
// 		if parser.Smsg != nil {
// 		}
// 		if parser.Fmsg != nil {
// 			if parser.Fmsg.Name != constName {
// 				b.Fatalf("Error: expcted name %s, but got %s", constName, parser.Fmsg.Name)
// 			}
// 		}
// 	}
// }

// func BenchmarkHeaderSmallInnerWoAllocation(b *testing.B) {
// 	inner := firstmessage.FirstMessage{
// 		Name:       constName,
// 		Id:         int32(100),
// 		SecondName: "secondRandomName",
// 	}
// 	data, err := proto.Marshal(&inner)
// 	if err != nil {
// 		log.Fatal("marshaling error: ", err)
// 	}

// 	hS := HeaderStruct.HeaderMessager{
// 		Type:    constName,
// 		Message: data,
// 	}
// 	headerData, err := proto.Marshal(&hS)
// 	if err != nil {
// 		log.Fatal("marshaling error: ", err)
// 	}
// 	for i := 0; i < b.N; i++ {
// 		headerParser := HeaderStruct.HeaderMessager{}
// 		err = proto.Unmarshal(headerData, &headerParser)
// 		if err != nil {
// 			log.Fatal("unmarshaling error: ", err)
// 		}

// 		headerMessage := firstmessage.FirstMessage{}
// 		err = proto.Unmarshal(headerParser.Message, &headerMessage)
// 		if err != nil {
// 			log.Fatal("unmarshaling error: ", err)
// 		}
// 		if headerParser.Type != constName {
// 			b.Fatalf("Error: expcted name %s, but got %s", constName, headerParser.Type)
// 		}
// 	}
// }

// func BenchmarkOneOfSmallInnerWoAllocation(b *testing.B) {
// 	inner := firstmessage.FirstMessage{
// 		Name:       constName,
// 		Id:         int32(100),
// 		SecondName: "secondRandomName",
// 	}

// 	hS := generaloneof.GeneralOneOf{
// 		Msg: &generaloneof.GeneralOneOf_Fmsg{
// 			Fmsg: &inner,
// 		},
// 	}

// 	headerData, err := proto.Marshal(&hS)
// 	if err != nil {
// 		log.Fatal("marshaling error: ", err)
// 	}
// 	for i := 0; i < b.N; i++ {
// 		headerParser := generaloneof.GeneralOneOf{}
// 		err = proto.Unmarshal(headerData, &headerParser)
// 		if err != nil {
// 			log.Fatal("unmarshaling error: ", err)
// 		}

// 		// _ := secondmessage.SecondMessage{}
// 		switch x := headerParser.Msg.(type) {
// 		case *generaloneof.GeneralOneOf_Msg1:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg2:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg3:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg4:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg5:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg6:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg7:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg8:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg9:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg10:
// 			continue
// 		case *generaloneof.GeneralOneOf_Fmsg:
// 			if x.Fmsg.Name != constName {
// 				b.Fatalf("Error: expcted name %s, but got %s", constName, x.Fmsg.Name)
// 			}
// 		case *generaloneof.GeneralOneOf_Smsg:
// 			continue
// 		}

// 	}
// }

// func BenchmarkAnyLargeInnerWoAllocation(b *testing.B) {
// 	aInner := AnotherInnerStruct.AnotherInnerMessage{
// 		Name:       "randomName",
// 		Id:         int32(100),
// 		SecondName: "secondRandomName",
// 		ThirdName:  "thirdRandomName",
// 		IdFloat:    float32(100.02),
// 		FourthName: "fourthRandomName",
// 		FifthName:  "fifthRandomName",
// 	}
// 	msg, err := ptypes.MarshalAny(&aInner)
// 	wrapper := UniversalParser.UniversalMessage{
// 		Name: constName,
// 		Msg:  msg,
// 	}
// 	data, err := proto.Marshal(&wrapper)
// 	if err != nil {
// 		log.Fatal("marshaling error: ", err)
// 	}
// 	for i := 0; i < b.N; i++ {
// 		parser := UniversalParser.UniversalMessage{}
// 		err = proto.Unmarshal(data, &parser)
// 		if err != nil {
// 			log.Fatal("unmarshaling error: ", err)
// 		}
// 		newAInner := AnotherInnerStruct.AnotherInnerMessage{}
// 		err = proto.Unmarshal(parser.Msg.Value, &newAInner)

// 		// err = ptypes.UnmarshalAny(parser.Msg, &newAInner)

// 		if err != nil {
// 			log.Fatal("unmarshaling error: ", err)
// 		}
// 		if parser.Name != constName {
// 			b.Fatalf("Error: expcted name %s, but got %s", constName, parser.Name)
// 		}
// 	}
// }

// func BenchmarkByteLargeInnerWoAllocation(b *testing.B) {
// 	aInner := AnotherInnerStruct.AnotherInnerMessage{
// 		Name:       "randomName",
// 		Id:         int32(100),
// 		SecondName: "secondRandomName",
// 		ThirdName:  "thirdRandomName",
// 		IdFloat:    float32(100.02),
// 		FourthName: "fourthRandomName",
// 		FifthName:  "fifthRandomName",
// 	}
// 	data, err := proto.Marshal(&aInner)
// 	if err != nil {
// 		log.Fatal("marshaling error: ", err)
// 	}
// 	buf := make([]byte, bufferSize, bufferSize)
// 	copy(buf[:], constName)
// 	overallData := append(buf, data...)
// 	for i := 0; i < b.N; i++ {
// 		// structName := string(bytes.TrimRight(overallData[:bufferSize], "\x00")[:])
// 		structName := string(overallData[:bufferSize])
// 		structSelf := overallData[bufferSize:]
// 		newInner := InnerStruct.InnerMessage{}
// 		err = proto.Unmarshal(structSelf, &newInner)
// 		if err != nil {
// 			log.Fatal("unmarshaling error: ", err)
// 		}

// 		if structName != constName {
// 			//			b.Fatalf("Error: expcted name %s, but got %s", constName, structName)
// 		}
// 	}
// }

// func BenchmarkGeneralLargeInnerWoAllocation(b *testing.B) {
// 	aInner := secondmessage.SecondMessage{
// 		Name:       constName,
// 		Id:         int32(100),
// 		SecondName: "secondRandomName",
// 		ThirdName:  "thirdRandomName",
// 		IdFloat:    float32(100.02),
// 		FourthName: "fourthRandomName",
// 		FifthName:  "fifthRandomName",
// 	}
// 	generalStruct := general.General{
// 		Smsg: &aInner,
// 	}
// 	data, err := proto.Marshal(&generalStruct)
// 	if err != nil {
// 		log.Fatal("marshaling error: ", err)
// 	}
// 	for i := 0; i < b.N; i++ {
// 		parser := general.General{}
// 		err = proto.Unmarshal(data, &parser)

// 		if err != nil {
// 			log.Fatal("unmarshaling error: ", err)
// 		}
// 		if parser.Msg10 != nil {
// 		}
// 		if parser.Msg9 != nil {
// 		}
// 		if parser.Msg8 != nil {
// 		}
// 		if parser.Msg7 != nil {
// 		}
// 		if parser.Msg6 != nil {
// 		}
// 		if parser.Msg5 != nil {
// 		}
// 		if parser.Msg4 != nil {
// 		}
// 		if parser.Msg3 != nil {
// 		}
// 		if parser.Msg2 != nil {
// 		}
// 		if parser.Msg1 != nil {
// 		}
// 		if parser.Fmsg != nil {
// 		}
// 		if parser.Smsg != nil {
// 			if parser.Smsg.Name != constName {
// 				b.Fatalf("Error: expcted name %s, but got %s", constName, parser.Smsg.Name)
// 			}
// 		}
// 	}
// }

// func BenchmarkHeaderLargeInnerWoAllocation(b *testing.B) {
// 	aInner := secondmessage.SecondMessage{
// 		Name:       constName,
// 		Id:         int32(100),
// 		SecondName: "secondRandomName",
// 		ThirdName:  "thirdRandomName",
// 		IdFloat:    float32(100.02),
// 		FourthName: "fourthRandomName",
// 		FifthName:  "fifthRandomName",
// 	}
// 	data, err := proto.Marshal(&aInner)
// 	if err != nil {
// 		log.Fatal("marshaling error: ", err)
// 	}

// 	hS := HeaderStruct.HeaderMessager{
// 		Type:    constName,
// 		Message: data,
// 	}
// 	headerData, err := proto.Marshal(&hS)
// 	if err != nil {
// 		log.Fatal("marshaling error: ", err)
// 	}
// 	for i := 0; i < b.N; i++ {
// 		headerParser := HeaderStruct.HeaderMessager{}
// 		err = proto.Unmarshal(headerData, &headerParser)
// 		if err != nil {
// 			log.Fatal("unmarshaling error: ", err)
// 		}

// 		headerMessage := secondmessage.SecondMessage{}
// 		err = proto.Unmarshal(headerParser.Message, &headerMessage)
// 		if err != nil {
// 			log.Fatal("unmarshaling error: ", err)
// 		}
// 		if headerParser.Type != constName {
// 			b.Fatalf("Error: expcted name %s, but got %s", constName, headerParser.Type)
// 		}
// 	}
// }

// func BenchmarkOneOfLargeInnerWoAllocation(b *testing.B) {
// 	aInner := &secondmessage.SecondMessage{
// 		Name:       constName,
// 		Id:         int32(100),
// 		SecondName: "secondRandomName",
// 		ThirdName:  "thirdRandomName",
// 		IdFloat:    float32(100.02),
// 		FourthName: "fourthRandomName",
// 		FifthName:  "fifthRandomName",
// 	}

// 	hS := generaloneof.GeneralOneOf{
// 		Msg: &generaloneof.GeneralOneOf_Smsg{
// 			Smsg: aInner,
// 		},
// 	}

// 	headerData, err := proto.Marshal(&hS)
// 	if err != nil {
// 		log.Fatal("marshaling error: ", err)
// 	}
// 	for i := 0; i < b.N; i++ {
// 		headerParser := generaloneof.GeneralOneOf{}
// 		err = proto.Unmarshal(headerData, &headerParser)
// 		if err != nil {
// 			log.Fatal("unmarshaling error: ", err)
// 		}

// 		// _ := secondmessage.SecondMessage{}
// 		switch x := headerParser.Msg.(type) {
// 		case *generaloneof.GeneralOneOf_Msg1:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg2:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg3:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg4:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg5:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg6:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg7:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg8:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg9:
// 			continue
// 		case *generaloneof.GeneralOneOf_Msg10:
// 			continue
// 		case *generaloneof.GeneralOneOf_Fmsg:
// 			continue
// 		case *generaloneof.GeneralOneOf_Smsg:
// 			if x.Smsg.Name != constName {
// 				b.Fatalf("Error: expcted name %s, but got %s", constName, x.Smsg.Name)
// 			}

// 		}

// 	}
// }

func BenchmarkGeneralArray(b *testing.B) {
	events := generateGeneralArrays()
	var err error
	for i := 0; i < b.N; i++ {
		ga := generalarray.GeneralArray{}
		err = proto.Unmarshal(events, &ga)
		if err != nil {
			log.Fatal("error unmarshaling")
		}
		for _, event := range ga.GenArray {
			if event.Msg10 != nil {
			}
			if event.Msg9 != nil {
			}
			if event.Msg8 != nil {
			}
			if event.Msg7 != nil {
			}
			if event.Msg6 != nil {
			}
			if event.Msg5 != nil {
			}
			if event.Msg4 != nil {
			}
			if event.Msg3 != nil {
			}
			if event.Msg2 != nil {
			}
			if event.Msg1 != nil {
			}
			if event.Fmsg != nil {
				if event.Fmsg.Name != constName {
					b.Fatalf("Error: expcted name %s, but got %s", constName, event.Fmsg.Name)
				}
			}
			if event.Smsg != nil {
				if event.Smsg.Name != constName {
					b.Fatalf("Error: expcted name %s, but got %s", constName, event.Smsg.Name)
				}
			}
		}
	}
}

func BenchmarkByteArray(b *testing.B) {
	events := generateEventsByte()
	// Чтобы не делать трим лишний раз на имени, будет якобы хранить ключи в nameBufferS буфере

	nameBuffer := make([]byte, nameBufferS, nameBufferS)
	copy(nameBuffer[:], "FirstMessage")
	firstName := string(nameBuffer)
	nameBuffer = make([]byte, nameBufferS, nameBufferS)
	copy(nameBuffer[:], "SecondMessage")
	secondName := string(nameBuffer)

	for i := 0; i < b.N; i++ {
		index := 0
		var err error
		for index < len(events) {
			// Get name
			structName := string(events[index : index+nameBufferS])
			index += nameBufferS
			if structName == firstName {
				index += vBufferS

				size := int(binary.LittleEndian.Uint32(events[index : index+sizeBufferS]))
				index += sizeBufferS
				fmsg := firstmessage.FirstMessage{}
				err = proto.Unmarshal(events[index:index+size], &fmsg)
				index += size
				if err != nil {
					log.Fatal("error unmarshling")
				}
				if fmsg.Name != constName {
					b.Fatalf("Error: expcted name %s, but got %s", constName, fmsg.Name)
				}
			}
			if structName == secondName {
				// version := string(events[i : i+vBufferS])
				index += vBufferS
				size := int(binary.LittleEndian.Uint32(events[index : index+sizeBufferS]))
				index += sizeBufferS
				smsg := secondmessage.SecondMessage{}
				err = proto.Unmarshal(events[index:index+size], &smsg)
				index += size + 1
				if err != nil {
					log.Fatal("error unmarshling")
				}
				if smsg.Name != constName {
					b.Fatalf("Error: expcted name %s, but got %s", constName, smsg.Name)
				}

			}
		}
	}
}

func BenchmarkArraysArray(b *testing.B) {
	data := generateEventsArrays()
	var err error
	for i := 0; i < b.N; i++ {
		goa := generaloptarray.GeneralOptArrays{}
		err = proto.Unmarshal(data, &goa)
		if err != nil {
			log.Fatal("error unmarshaling")
		}
		if goa.Fmsgs != nil {
			for _, msg := range goa.Fmsgs {
				if msg.Name != constName {
					b.Fatalf("Error: expcted name %s, but got %s", constName, msg.Name)
				}
			}
		}
		if goa.Smsgs != nil {
			for _, msg := range goa.Fmsgs {
				if msg.Name != constName {
					b.Fatalf("Error: expcted name %s, but got %s", constName, msg.Name)
				}
			}
		}
	}
}

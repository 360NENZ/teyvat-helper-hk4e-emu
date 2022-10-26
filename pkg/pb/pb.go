package pb

import (
	"log"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

type ProtoCommand uint16

type ProtoMessageType string

type ProtoMessage interface {
	protoreflect.ProtoMessage
	ProtoCommand() ProtoCommand
	ProtoMessageType() ProtoMessageType
}

type ProtoCommandNewFunc func() ProtoMessage

type protoCommandNewFuncMap map[ProtoCommand]ProtoCommandNewFunc

func newProtoCommandNewFuncMap(fs ...ProtoCommandNewFunc) protoCommandNewFuncMap {
	m := make(protoCommandNewFuncMap)
	m.Add(fs...)
	return m
}

func (m protoCommandNewFuncMap) New(cmd ProtoCommand) ProtoMessage {
	if f, ok := m[cmd]; ok {
		return f()
	}
	log.Println("unknown command:", cmd)
	return nil
}

func (m protoCommandNewFuncMap) Add(fs ...ProtoCommandNewFunc) {
	for _, f := range fs {
		v := f()
		if e, ok := m[v.ProtoCommand()]; ok {
			panic("duplicate command: " + v.ProtoMessageType() + " exists: " + e().ProtoMessageType())
		}
		m[v.ProtoCommand()] = f
	}
}

var ProtoCommandNewFuncMap = newProtoCommandNewFuncMap()

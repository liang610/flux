// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbsemantic

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Position struct {
	_tab flatbuffers.Struct
}

func (rcv *Position) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Position) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Position) Line() int32 {
	return rcv._tab.GetInt32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Position) MutateLine(n int32) bool {
	return rcv._tab.MutateInt32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Position) Column() int32 {
	return rcv._tab.GetInt32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *Position) MutateColumn(n int32) bool {
	return rcv._tab.MutateInt32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func CreatePosition(builder *flatbuffers.Builder, line int32, column int32) flatbuffers.UOffsetT {
	builder.Prep(4, 8)
	builder.PrependInt32(column)
	builder.PrependInt32(line)
	return builder.Offset()
}
package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Push byte
type BIPUSH struct {
	val int8
}

func (p *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	p.val = reader.ReadInt8()
}

func (p *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(p.val)
	frame.OperandStack().PushInt(i)
}

// Push short
type SIPUSH struct {
	val int16
}

func (p *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	p.val = reader.ReadInt16()
}

func (p *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(p.val)
	frame.OperandStack().PushInt(i)
}

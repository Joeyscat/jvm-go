package conversions

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Convert int to byte
type I2B struct {
	base.NoOperandsInstruction
}

func (i *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	b := int32(int8(val))
	stack.PushInt(b)
}

// Convert int to char
type I2C struct {
	base.NoOperandsInstruction
}

func (i *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	c := int32(uint16(val))
	stack.PushInt(c)
}

// Convert int to short
type I2S struct {
	base.NoOperandsInstruction
}

func (i *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	s := int32(uint16(val))
	stack.PushInt(s)
}

// Convert int to long
type I2L struct {
	base.NoOperandsInstruction
}

func (i *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	l := int64(val)
	stack.PushLong(l)
}

// Convert int to float
type I2F struct {
	base.NoOperandsInstruction
}

func (i *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	f := float32(val)
	stack.PushFloat(f)
}

// Convert int to double
type I2D struct {
	base.NoOperandsInstruction
}

func (i *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	d := float64(val)
	stack.PushDouble(d)
}

package conversions

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Convert double to float
type D2F struct {
	base.NoOperandsInstruction
}

func (d *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	f := float32(val)
	stack.PushFloat(f)
}

// Convert double to int
type D2I struct {
	base.NoOperandsInstruction
}

func (d *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	i := int32(val)
	stack.PushInt(i)
}

// Convert double to long
type D2L struct {
	base.NoOperandsInstruction
}

func (d *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	l := int64(val)
	stack.PushLong(l)
}

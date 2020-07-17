package conversions

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Convert long to double
type L2D struct {
	base.NoOperandsInstruction
}

func (l *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	d := float64(val)
	stack.PushDouble(d)
}

// Convert long to float
type L2F struct {
	base.NoOperandsInstruction
}

func (l *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	f := float32(val)
	stack.PushFloat(f)
}

// Convert long to int
type L2I struct {
	base.NoOperandsInstruction
}

func (l *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	i := int32(val)
	stack.PushInt(i)
}

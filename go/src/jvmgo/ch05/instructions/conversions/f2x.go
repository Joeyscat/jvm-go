package conversions

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Convert float to double
type F2D struct {
	base.NoOperandsInstruction
}

func (f *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	d := float64(val)
	stack.PushDouble(d)
}

// Convert float to int
type F2I struct {
	base.NoOperandsInstruction
}

func (d *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	i := int32(val)
	stack.PushInt(i)
}

// Convert float to long
type F2L struct {
	base.NoOperandsInstruction
}

func (d *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	l := int64(val)
	stack.PushLong(l)
}

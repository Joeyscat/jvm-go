package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Branch if int comparison with zero succeeds
type IFEQ struct {
	base.BranchInstruction
}

func (cond *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, cond.Offset)
	}
}

type IFNE struct {
	base.BranchInstruction
}

func (cond *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, cond.Offset)
	}
}

type IFLT struct {
	base.BranchInstruction
}

func (cond *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, cond.Offset)
	}
}

type IFLE struct {
	base.BranchInstruction
}

func (cond *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, cond.Offset)
	}
}

type IFGT struct {
	base.BranchInstruction
}

func (cond *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, cond.Offset)
	}
}

type IFGE struct {
	base.BranchInstruction
}

func (cond *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, cond.Offset)
	}
}

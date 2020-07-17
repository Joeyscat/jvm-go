package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Branch if int comparison succeeds
type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (cmp *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		base.Branch(frame, cmp.Offset)
	}
}

type IF_ICMPNE struct {
	base.BranchInstruction
}

func (cmp *IF_ICMPNE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		base.Branch(frame, cmp.Offset)
	}
}

type IF_ICMPLT struct {
	base.BranchInstruction
}

func (cmp *IF_ICMPLT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base.Branch(frame, cmp.Offset)
	}
}

type IF_ICMPLE struct {
	base.BranchInstruction
}

func (cmp *IF_ICMPLE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base.Branch(frame, cmp.Offset)
	}
}

type IF_ICMPGT struct {
	base.BranchInstruction
}

func (cmp *IF_ICMPGT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.Branch(frame, cmp.Offset)
	}
}

type IF_ICMPGE struct {
	base.BranchInstruction
}

func (cmp *IF_ICMPGE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		base.Branch(frame, cmp.Offset)
	}
}

func _icmpPop(frame *rtda.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}

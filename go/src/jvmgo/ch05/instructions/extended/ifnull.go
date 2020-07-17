package extended

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Branch uf reference is null
type IFNULL struct {
	base.BranchInstruction
}

func (cond *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, cond.Offset)
	}
}

// Branch uf reference not null
type IFNOTNULL struct {
	base.BranchInstruction
}

func (cond *IFNOTNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, cond.Offset)
	}
}

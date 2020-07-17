package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Increment local variable by constant
type IINC struct {
	Index uint
	Const int32
}

func (i *IINC) FetchOperands(reader *base.BytecodeReader) {
	i.Index = uint(reader.ReadInt8())
	i.Const = int32(reader.ReadInt8())
}

func (i *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(i.Index)
	val += i.Const
	localVars.SetInt(i.Index, val)
}

package base

import "jvmgo/ch05/rtda"

type Instruction interface {
	// 从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	// 执行指令逻辑
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (noi *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

type BranchInstruction struct {
	Offset int
}

func (bi *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	bi.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (i8i *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	i8i.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (i8i *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	i8i.Index = uint(reader.ReadUint16())
}

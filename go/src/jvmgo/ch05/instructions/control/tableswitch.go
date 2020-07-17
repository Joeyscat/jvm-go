package control

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

/*
tableswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4
jump offsets...
*/
// Access jump table by index and jump
type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (s *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	s.defaultOffset = reader.ReadInt32()
	s.low = reader.ReadInt32()
	s.high = reader.ReadInt32()
	jumpOffsetsCount := s.high - s.low + 1
	s.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (s *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= s.low && index <= s.high {
		offset = int(s.jumpOffsets[index-s.low])
	} else {
		offset = int(s.defaultOffset)
	}

	base.Branch(frame, offset)
}

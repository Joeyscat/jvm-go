package control

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

/*
lookupswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
npairs1
npairs2
npairs3
npairs4
match-offset pairs...
*/
// Access jump table by key match and jump
type LOOKUP_SWITCH struct {
	defaultOffet int32
	npairs       int32
	matchOffsets []int32
}

func (s *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	s.defaultOffet = reader.ReadInt32()
	s.npairs = reader.ReadInt32()
	s.matchOffsets = reader.ReadInt32s(s.npairs * 2)
}

func (s *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < s.npairs*2; i += 2 {
		if s.matchOffsets[i] == key {
			offset := s.matchOffsets[i]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(s.defaultOffet))
}

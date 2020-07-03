package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

// 读取u1类型数据
func (cr *ClassReader) readUint8() unit8 {
	val := cr.data[0]
	cr.data = cr.data[1:]
	return val
}

// 读取u2类型数据
func (cr *ClassReader) readUint16() unit16 {
	val := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return val
}

// 读取u1类型数据
func (cr *ClassReader) readUint32() unit32 {
	val := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return val
}

func (cr *ClassReader) readUint64() unit64 {
	val := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return val
}

func (cr *ClassReader) readUint16s() []uint16 {
	n := cr.readUnit16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = cr.readUnit16()
	}
	return s
}

func (cr *ClassReader) readBytes(n unit32) []byte {
	bytes := cr.data[:n]
	cr.data = cr.data[n:]
	return bytes
}

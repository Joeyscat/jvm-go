package classfile

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (c *UnparsedAttribute) readInfo(reader *ClassReader) {
	c.info = reader.readBytes(c.length)
}

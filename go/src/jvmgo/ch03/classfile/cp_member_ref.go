package classfile

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}
type ConstantInterfacerefInfo struct {
	ConstantMemberrefInfo
}

func (c *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	c.classIndex = reader.readUint16()
	c.nameAndTypeIndex = reader.readUint16()
}

func (c *ConstantMemberrefInfo) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}

func (c *ConstantMemberrefInfo) NameAndDescriptor() string {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}

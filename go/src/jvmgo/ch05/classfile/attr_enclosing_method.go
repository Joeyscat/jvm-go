package classfile

/*
EnclosingMethod_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 class_index;
    u2 method_index;
}
*/
type EnclosingMethodAttribute struct {
	cp          ConstantPool
	classIndex  uint16
	methodIndex uint16
}

func (ema *EnclosingMethodAttribute) readInfo(reader *ClassReader) {
	ema.classIndex = reader.readUint16()
	ema.methodIndex = reader.readUint16()
}

func (ema *EnclosingMethodAttribute) ClassName() string {
	return ema.cp.getClassName(ema.classIndex)
}

func (ema *EnclosingMethodAttribute) MethodNameAndDescriptor() (string, string) {
	if ema.methodIndex > 0 {
		return ema.cp.getNameAndType(ema.methodIndex)
	} else {
		return "", ""
	}
}

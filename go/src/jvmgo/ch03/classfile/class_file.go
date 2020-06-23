package classfile

import "fmt"

type ClassFile struct {
	// magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMember(reader, self.constantPool)
	self.methods = readMember(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

// 检查magic
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.class.ClassFormatError: magic!")
	}
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// getter
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

// getter
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

// getter
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

// getter
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

// getter
func (self *ClassFile) Fields() []*MemberInfo {
	return self.Fields()
}

// getter
func (self *ClassFile) Methods() *MemberInfo {
	return self.Methods()
}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" // 只有java.lang.Onject没有超类
}
func (self *ClassFile) InterfaceNames() []string {
	interfaceName := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceName[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceName
}

/**
相比Java语言，Go的访问控制非常简单：只有公开和私有两种。
所有首字母大写的类型、结构体、字段、变量、函数、方法等都是公开的，
可供其他包使用。首字母小写则是私有的，只能在包内部使用。

在本书的代码中，尽量只公开必要的变量、字段、函数和方法等
。但是为了提高代码可读性，所有的结构体都是公开的，也就是首字母是大写的。
*/

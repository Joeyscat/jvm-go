package classfile

import "fmt"

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
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

func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constantPool = readConstantPool(reader)
	cf.accessFlags = reader.readUint16()
	cf.thisClass = reader.readUint16()
	cf.superClass = reader.readUint16()
	cf.interfaces = reader.readUint16s()
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool)
}

// 检查magic
func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.class.ClassFormatError: magic!")
	}
}

func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()
	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// getter
func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

// getter
func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

// getter
func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

// getter
func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

// getter
func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}

// getter
func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	return "" // 只有java.lang.Onject没有超类
}
func (cf *ClassFile) InterfaceNames() []string {
	interfaceName := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		interfaceName[i] = cf.constantPool.getClassName(cpIndex)
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

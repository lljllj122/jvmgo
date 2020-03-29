package classfile

const (
	Utf8Tag               = 1
	IntegerTag            = 3
	FloatTag              = 4
	LongTag               = 5
	DoubleTag             = 6
	ClassTag              = 7
	StringTag             = 8
	FieldRefTag           = 9
	MethodRefTag          = 10
	NameAndTypeTag        = 12
	InterfaceMethodRefTag = 11
	MethodHandleTag       = 15
	MethodTypeTag         = 16
	InvokeDyanmicTag      = 18
)

// Interface for ConstantInfo item
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, pool *ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	constantInfo := newConstantInfo(tag, pool)
	constantInfo.readInfo(reader)
	return constantInfo
}

func newConstantInfo(tag u1, pool *ConstantPool) ConstantInfo {
	switch tag {
	case IntegerTag:
		return &ConstantIntegerInfo{}
	case FloatTag:
		return &ConstantFloatInfo{}
	case LongTag:
		return &ConstantFloatInfo{}
	case DoubleTag:
		return &ConstantDoubleInfo{}
	case Utf8Tag:
		return &ConstantUtf8Info{}
	case StringTag:
		return &ConstantStringInfo{pool: pool}
	case ClassTag:
		return &ConstantClassInfo{pool: pool}
	case FieldRefTag:
		return &ConstantMemberRefInfo{pool: pool}
	case MethodRefTag:
		return &ConstantMemberRefInfo{pool: pool}
	case InterfaceMethodRefTag:
		return &ConstantMemberRefInfo{pool: pool}
	case NameAndTypeTag:
		return &ConstantNameAndTypeInfo{}
	case MethodTypeTag:
		return &ConstantMethodTypeInfo{}
	case MethodHandleTag:
		return &ConstantMethodHandleInfo{}
	case InvokeDyanmicTag:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: invalid constant pool tag.")
	}
}

package classfile

type ConstantPool struct {
	constantInfos []ConstantInfo
}

func readConstantPool(reader *ClassReader) *ConstantPool {
	constantPoolCount := int(reader.readUint16())
	pool := &ConstantPool{
		constantInfos: make([]ConstantInfo, constantPoolCount),
	}
	// i start from 1; 0 is not a valid index in constant pool
	for i := 1; i < constantPoolCount; i++ {
		pool.constantInfos[i] = readConstantInfo(reader, pool)
		switch pool.constantInfos[i].(type) {
		// Long and Double constant info occoupy 2 slots in pool
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return pool
}

func (pool *ConstantPool) Size() int {
	return len(pool.constantInfos)
}

// get a ConstantInfo item by index from the pool
func (pool *ConstantPool) getConstantInfo(index u2) ConstantInfo {
	if info := pool.constantInfos[index]; info != nil {
		return info
	}
	panic("Invalid constant pool index.")
}

func (pool *ConstantPool) getNameAndType(index u2) (_name string, _type string) {
	nameAndTypeInfo := pool.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	// name and type are stored as items in constant pool
	_name = pool.getUtf8(nameAndTypeInfo.nameIndex)
	_type = pool.getUtf8(nameAndTypeInfo.descriptorIndex)
	return
}

func (pool *ConstantPool) getUtf8(index u2) string {
	utf8Info := pool.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

func (pool *ConstantPool) getClassName(index u2) string {
	nameIndex := pool.getConstantInfo(index).(*ConstantClassInfo).nameIndex
	return pool.getUtf8(nameIndex)
}

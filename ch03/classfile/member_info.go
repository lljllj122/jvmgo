package classfile

// MemberInfo :  a struct to represents both field and method items in class
type MemberInfo struct {
	constantPool *ConstantPool
	// all fields are in sequetial order
	accessFlag      u2
	nameIndex       u2 // Index of the member name in constant pool;
	descriptorIndex u2 // Index of the descriptor in constant pool; A descriptor is a string representing the type of a field or method.
	attributes      []*AttributeInfo
}

func readMembers(reader *ClassReader, constantPool *ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, constantPool)
	}
	return members
}

func readMember(reader *ClassReader, constantPool *ConstantPool) *MemberInfo {
	m := MemberInfo{
		constantPool:    constantPool,
		accessFlag:      reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, constantPool),
	}
	return &m
}

func (m *MemberInfo) AccessFlag() u2 {
	return m.accessFlag
}

func (m *MemberInfo) Name() string {
	return m.constantPool.getUtf8(m.nameIndex)
}

func (m *MemberInfo) Descriptor() string {
	return m.constantPool.getUtf8(m.descriptorIndex)
}

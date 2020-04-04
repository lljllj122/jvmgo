package base

type BytecodeReader struct {
	code []byte
	pc   int
}

func (reader *BytecodeReader) Reset(code []byte, pc int) {
	reader.pc = pc
	reader.code = code
}

func (reader *BytecodeReader) ReadByte() byte {
	u1 := reader.code[reader.pc]
	reader.pc++
	return u1
}

func (reader *BytecodeReader) ReadUint16() uint16 {
	high := reader.ReadByte()
	low := reader.ReadByte()
	return uint16(high)<<8 | uint16(low)
}

func (reader *BytecodeReader) ReadInt16() int16 {
	return int16(reader.ReadUint16())
}

func (reader *BytecodeReader) ReadUint32() uint32 {
	high := reader.ReadUint16()
	low := reader.ReadUint16()
	return uint32(high)<<16 | uint32(low)
}

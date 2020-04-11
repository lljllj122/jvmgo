package classfile

import "math"

/*
CONSTANT_Integer_info {
   u1 tag;
   u4 bytes;
}
to store types of boolean, byte, short, char, and int (<= 32-bit)
*/
type ConstantIntegerInfo struct {
	value int32
}

func (info *ConstantIntegerInfo) Value() int32 {
	return info.value
}

func (info *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	info.value = int32(bytes)
}

/*
CONSTANT_Float_info {
   u1 tag;
   u4 bytes;
}
*/
type ConstantFloatInfo struct {
	value float32
}

func (info *ConstantFloatInfo) Value() float32 {
	return info.value
}

func (info *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	info.value = math.Float32frombits(bytes)
}

/*
CONSTANT_Long_info {
   u1 tag;
   u4 high_bytes;
   u4 low_bytes;
}
*/
type ConstantLongInfo struct {
	value int64
}

func (info *ConstantLongInfo) Value() int64 {
	return info.value
}

func (info *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	info.value = int64(bytes)
}

/*
CONSTANT_Double_info {
   u1 tag;
   u4 high_bytes;
   u4 low_bytes;
}
*/
type ConstantDoubleInfo struct {
	// IEEE 754 Floating-Point Arithmetic.
	// equivalent to float64 in Go
	value float64
}

func (info *ConstantDoubleInfo) Value() float64 {
	return info.value
}

func (info *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	info.value = math.Float64frombits(bytes)
}

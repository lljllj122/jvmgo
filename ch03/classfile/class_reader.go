package classfile

import "encoding/binary"

// JVM defines 3 types to represent data
type u1 = uint8  // one byte
type u2 = uint16 // two bytes
type u4 = uint32 // four bytes

// ClassReader : a entity to read class in stream of 8-bit bytes
// byte is an alias for uint8 and is equivalent to uint8 in all ways.
type ClassReader struct {
	data []byte
}

// u1
func (cr *ClassReader) readUint8() uint8 {
	val := cr.data[0]
	cr.data = cr.data[1:]
	return val
}

// u2
func (cr *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(cr.data[:2])
	cr.data = cr.data[2:]
	return val
}

// u4
func (cr *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(cr.data[:4])
	cr.data = cr.data[4:]
	return val
}

// not actually defined in JVM types
func (cr *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(cr.data[:8])
	cr.data = cr.data[8:]
	return val
}

// Read a talbe of u2
func (cr *ClassReader) readUint16Table() []uint16 {
	// the header of table defines number of items in table
	n := cr.readUint16()
	u2s := make([]u2, n)
	for i := range u2s {
		u2s[i] = cr.readUint16()
	}
	return u2s
}

// Read n bytes
func (cr *ClassReader) readBytes(n uint32) []byte {
	bytes := cr.data[:n]
	cr.data = cr.data[n:]
	return bytes
}

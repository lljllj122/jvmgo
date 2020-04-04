package rtda

import "math"

/*
Slot is used to represent the local variable element.
A single local variable can hold a value of type
boolean, byte, char, short, int, float, reference, or returnAddress.
A pair of local variables can hold a value of type long or double.

We use uint32 to store the pimitive data types.
*/
type Slot struct {
	// the size of uint Golang is determined by the OS.
	bits uint32
	ref  *Object
}

type LocalVars []Slot

func newLocalVars(maxLocals uint8) LocalVars {
	// size of local variable is defined during compiling.
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (localVars LocalVars) SetInt(index uint8, val int32) {
	localVars[index].bits = uint32(val)
}

func (localVars LocalVars) GetInt(index uint8) int32 {
	return int32(localVars[index].bits)
}

func (localVars LocalVars) SetFloat(index uint8, val float32) {
	// convert the float value to 32-bits
	bits := math.Float32bits(val)
	localVars[index].bits = bits
}

func (localVars LocalVars) GetFloat(index uint8) float32 {
	bits := uint32(localVars[index].bits)
	return math.Float32frombits(bits)
}

func (localVars LocalVars) SetLong(index uint8, val int64) {
	// use Big endian
	localVars[index].bits = uint32(val >> 32) // high bits
	localVars[index+1].bits = uint32(val)     // low bits

}

func (localVars LocalVars) GetLong(index uint8) int64 {
	highBits := localVars[index].bits
	lowBits := localVars[index+1].bits
	return int64(highBits)<<32 | int64(lowBits)
}

func (localVars LocalVars) SetDouble(index uint8, val float64) {
	bits := math.Float64bits(val)
	// use Big endian
	localVars[index].bits = uint32(bits >> 32) // high bits
	localVars[index+1].bits = uint32(bits)     // low bits
}

func (localVars LocalVars) GetDouble(index uint8) float64 {
	highBits := localVars[index].bits
	lowBits := localVars[index+1].bits
	return math.Float64frombits(uint64(highBits)<<32 | uint64(lowBits))
}

func (localVars LocalVars) SetRef(index uint8, ref *Object) {
	localVars[index].ref = ref
}

func (localVars LocalVars) GetRef(index uint8) *Object {
	return localVars[index].ref
}

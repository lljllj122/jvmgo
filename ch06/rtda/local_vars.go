package rtda

import (
	"jvmgo/ch06/rtda/heap"
	"math"
)

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	// size of local variable is defined during compiling.
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (localVars LocalVars) SetInt(index uint, val int32) {
	localVars[index].bits = uint32(val)
}

func (localVars LocalVars) GetInt(index uint) int32 {
	return int32(localVars[index].bits)
}

func (localVars LocalVars) SetFloat(index uint, val float32) {
	// convert the float value to 32-bits
	bits := math.Float32bits(val)
	localVars[index].bits = bits
}

func (localVars LocalVars) GetFloat(index uint) float32 {
	bits := uint32(localVars[index].bits)
	return math.Float32frombits(bits)
}

func (localVars LocalVars) SetLong(index uint, val int64) {
	// use Big endian
	localVars[index].bits = uint32(val >> 32) // high bits
	localVars[index+1].bits = uint32(val)     // low bits

}

func (localVars LocalVars) GetLong(index uint) int64 {
	highBits := localVars[index].bits
	lowBits := localVars[index+1].bits
	return int64(highBits)<<32 | int64(lowBits)
}

func (localVars LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	// use Big endian
	localVars[index].bits = uint32(bits >> 32) // high bits
	localVars[index+1].bits = uint32(bits)     // low bits
}

func (localVars LocalVars) GetDouble(index uint) float64 {
	highBits := localVars[index].bits
	lowBits := localVars[index+1].bits
	return math.Float64frombits(uint64(highBits)<<32 | uint64(lowBits))
}

func (localVars LocalVars) SetRef(index uint, ref *heap.Object) {
	localVars[index].ref = ref
}

func (localVars LocalVars) GetRef(index uint) *heap.Object {
	return localVars[index].ref
}

package heap

import "math"

type Slot struct {
	// the size of uint Golang is determined by the OS.
	bits uint32
	ref  *Object
}

type Slots []Slot

func newSlots(varCount uint) Slots {
	if varCount > 0 {
		return make([]Slot, varCount)
	}
	return nil
}

func (slots Slots) SetInt(index uint, val int32) {
	slots[index].bits = uint32(val)
}

func (slots Slots) GetInt(index uint) int32 {
	return int32(slots[index].bits)
}

func (slots Slots) SetFloat(index uint, val float32) {
	// convert the float value to 32-bits
	bits := math.Float32bits(val)
	slots[index].bits = bits
}

func (slots Slots) GetFloat(index uint) float32 {
	bits := uint32(slots[index].bits)
	return math.Float32frombits(bits)
}

func (slots Slots) SetLong(index uint, val int64) {
	// use Big endian
	slots[index].bits = uint32(val >> 32) // high bits
	slots[index+1].bits = uint32(val)     // low bits

}

func (slots Slots) GetLong(index uint) int64 {
	highBits := slots[index].bits
	lowBits := slots[index+1].bits
	return int64(highBits)<<32 | int64(lowBits)
}

func (slots Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	// use Big endian
	slots[index].bits = uint32(bits >> 32) // high bits
	slots[index+1].bits = uint32(bits)     // low bits
}

func (slots Slots) GetDouble(index uint) float64 {
	highBits := slots[index].bits
	lowBits := slots[index+1].bits
	return math.Float64frombits(uint64(highBits)<<32 | uint64(lowBits))
}

func (slots Slots) SetRef(index uint, ref *Object) {
	slots[index].ref = ref
}

func (slots Slots) GetRef(index uint) *Object {
	return slots[index].ref
}

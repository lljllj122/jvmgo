package rtda

import "math"

/*
The maximum depth of the operand stack of a frame is determined at compile-time.
*/
type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}

	return nil
}

func (stack *OperandStack) PushInt(val int32) {
	stack.slots[stack.size].bits = uint32(val)
	stack.size++
}

func (stack *OperandStack) PopInt() int32 {
	stack.size--
	res := int32(stack.slots[stack.size].bits)
	return res
}

func (stack *OperandStack) PushFloat(val float32) {
	stack.slots[stack.size].bits = math.Float32bits(val)
	stack.size++
}

func (stack *OperandStack) PopFloat() float32 {
	stack.size--
	res := math.Float32frombits(stack.slots[stack.size].bits)
	return res
}

func (stack *OperandStack) PushLong(val int64) {
	// use Big endian
	bits := uint64(val)
	stack.slots[stack.size].bits = uint32(bits >> 32) // high bits
	stack.slots[stack.size+1].bits = uint32(bits)     // low bits
	stack.size += 2
}

func (stack *OperandStack) PopLong() int64 {
	highBits := stack.slots[stack.size-2].bits
	lowBits := stack.slots[stack.size-1].bits
	stack.size -= 2
	return int64(highBits)<<32 | int64(lowBits)
}

func (stack *OperandStack) PushDouble(val float64) {
	// use Big endian
	bits := math.Float64bits(val)
	stack.slots[stack.size].bits = uint32(bits >> 32) // high bits
	stack.slots[stack.size+1].bits = uint32(bits)     // low bits
	stack.size += 2
}

func (stack *OperandStack) PopDouble() float64 {
	highBits := stack.slots[stack.size-2].bits
	lowBits := stack.slots[stack.size-1].bits
	stack.size -= 2
	return math.Float64frombits(uint64(highBits)<<32 | uint64(lowBits))
}

func (stack *OperandStack) PushRef(ref *Object) {
	stack.slots[stack.size].ref = ref
	stack.size++
}

func (stack *OperandStack) PopRef() (ref *Object) {
	stack.size--
	ref = stack.slots[stack.size].ref
	stack.slots[stack.size].ref = nil
	return
}

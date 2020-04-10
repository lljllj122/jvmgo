package rtda

import "jvmgo/ch06/rtda/heap"

/*
Slot is used to represent the elements in localVars and OpStack
A single local variable can hold a value of type
boolean, byte, char, short, int, float, reference, or returnAddress.
A pair of local variables can hold a value of type long or double.

We use uint32 to store the primitive data types.
*/
type Slot struct {
	// the size of uint Golang is determined by the OS.
	bits uint32
	ref  *heap.Object
}

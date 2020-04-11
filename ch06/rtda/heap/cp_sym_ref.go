package heap

type SymbolRef struct {
	constantPool *ConstantPool
	className    string
	class        *Class
}

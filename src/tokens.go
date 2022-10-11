package main

type TokenId uint64

const (
	Unknown TokenId = iota
	// Declaritors
	VariableDeclaritor
	ConstantDeclaritor
	FunctionDeclaritor
	ObjectDeclaritor
	// Primitive types
	Bit
	Intager8
	Intager16
	Intager32
	Intager64
	Intager128
	Intager256
	Unsigned8
	Unsigned16
	Unsigned32
	Unsigned64
	Unsigned128
	Unsigned256
	FloatingPoint16
	FloatingPoint32
	FloatingPoint64
	FloatingPoint128
	Character
	MemoryAddress
	// Mathematical operators
	Additon
	AddAndSet
	Subtraction
	SubtractAndSet
	Multiplication
	MultiplyAndSet
	Division
	DivideAndSet
	Exponentiation
	ExponentiateAndSet
	RootExtraction
	RootExtractAndSet
	Increment
	Decrement
	// Conditional operators
	Equals
	NotEquals
	LessThan
	GreaterThan
	LessThanOrEquals
	GreaterThanOrEquals
	Not
	And
	Or
	Xor
	Nand
	Nor
	Xnor
	// Binary operators
	LeftShift
	RightShift
	LeftRotate
	RightRotate
	BinaryNot
	BinaryAnd
	BinaryOr
	BinaryXor
	BinaryNand
	BinaryNor
	BinaryXnor
	// Other Operators
	OpenGroup
	CloseGroup
	// Todo: Add more tokens
)

type Token struct {
	id   TokenId
	data []uint64
}

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
	Type
	Void
	// Other builtin types
	Array
	ArraySize
	List
	// Mathematical operators
	Set
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
	NotAndSet
	And
	AndAndSet
	Or
	OrAndSet
	Xor
	XorAndSet
	Nand
	NandAndSet
	Nor
	NorAndSet
	Xnor
	XnorAndSet
	// Binary operators
	LeftShift
	LeftShiftAndSet
	RightShift
	RightShiftAndSet
	LeftRotate
	LeftRotateAndSet
	RightRotate
	RightRotateAndSet
	BinaryNot
	BinaryNotAndSet
	BinaryAnd
	BinaryAndAndSet
	BinaryOr
	BinaryOrAndSet
	BinaryXor
	BinaryXorAndSet
	BinaryNand
	BinaryNandAndSet
	BinaryNor
	BinaryNorAndSet
	BinaryXnor
	BinaryXnorAndSet
	// Other Operators
	OpenGroup
	CloseGroup
	// Usages
	VariableUsage
	ConstantUsage
	FunctionUsage
	// Functions
	Parameter
	ReturnType
	Return
	// Objects
	Public
	Single
	Base
	Final
	Self
	Extends
	Constructor
	Destructor
	// Loops
	Loop
	Skip
	Break
	// Conditional logic
	If
	Other
	Match
	// Multithreading
	Thread
	ThreadNumber
	Volitile
	// Compiler builtins
	CompilerFunctionTypeSet
	CompilerFunctionTypeGet
	// Scoping
	EnterScopeLevel
	LeaveScopeLevel
	Parent
	// Goto
	Goto
	GotoLabel
	// Other
	Delimiter
	Entry
	InsertLLVM
)

type Token struct {
	id   TokenId
	data []uint64
}

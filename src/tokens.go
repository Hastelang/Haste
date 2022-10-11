package main

type TokenId uint64

const (
	Unknown TokenId = iota
	// Todo: Add tokens
)

type Token struct {
	id TokenId
	data []uint64
}
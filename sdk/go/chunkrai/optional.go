// Code generated by Fern. DO NOT EDIT.

package chunkrai

import (
	core "github.com/gitslim/chunkr-ai-sdk/sdk/go/chunkrai/core"
)

// Optional initializes an optional field.
func Optional[T any](value T) *core.Optional[T] {
	return &core.Optional[T]{
		Value: value,
	}
}

// Null initializes an optional field that will be sent as
// an explicit null value.
func Null[T any]() *core.Optional[T] {
	return &core.Optional[T]{
		Null: true,
	}
}

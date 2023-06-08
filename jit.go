package jit

import (
	"reflect"

	"github.com/x-research-team/jit/internal/ccall"
)

var (
	TypeInt       = &Type{ccall.TypeGoInt}
	TypeFloat32   = &Type{ccall.TypeFloat32}
	TypeFloat64   = &Type{ccall.TypeFloat64}
	TypeString    = &Type{ccall.TypeGoString}
	TypeInterface = &Type{ccall.TypeGoInterface}
)

func ReflectTypeToType(t reflect.Type) *Type {
	return &Type{ccall.ReflectTypeToType(t)}
}

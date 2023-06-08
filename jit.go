package jit

import (
	"github.com/x-research-team/jit/internal/ccall"
)

var (
	TypeInt = &Type{ccall.TypeGoInt}
)

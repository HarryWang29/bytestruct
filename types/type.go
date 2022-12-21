package types

import (
    "github.com/HarryWang29/bytestruct/runtime"
    "unsafe"
)

type Codec interface {
    DecodeStream(*Stream, unsafe.Pointer) error
    EncodeStream(*Stream, unsafe.Pointer) error
}

type emptyInterface struct {
    typ *runtime.Type
    ptr unsafe.Pointer
}

const (
    nul                   = '\000'
    maxDecodeNestingDepth = 10000
)

//nolint:golint
//go:linkname unsafe_New reflect.unsafe_New
func unsafe_New(*runtime.Type) unsafe.Pointer

func UnsafeNew(t *runtime.Type) unsafe.Pointer {
    return unsafe_New(t)
}

package types

import (
    "github.com/HarryWang29/bytestruct/runtime"
    "unsafe"
)

type anonymousFieldType struct {
    structType *runtime.Type
    offset     uintptr
    dec        Codec
}

func newAnonymousFieldDecoder(structType *runtime.Type, offset uintptr, dec Codec) *anonymousFieldType {
    return &anonymousFieldType{
        structType: structType,
        offset:     offset,
        dec:        dec,
    }
}

func (d *anonymousFieldType) DecodeStream(s *Stream, p unsafe.Pointer) error {
    if *(*unsafe.Pointer)(p) == nil {
        *(*unsafe.Pointer)(p) = unsafe_New(d.structType)
    }
    p = *(*unsafe.Pointer)(p)
    return d.dec.DecodeStream(s, unsafe.Pointer(uintptr(p)+d.offset))
}

func (d *anonymousFieldType) EncodeStream(s *Stream, p unsafe.Pointer) error {
    if *(*unsafe.Pointer)(p) == nil {
        return nil
    }
    p = *(*unsafe.Pointer)(p)
    return d.dec.EncodeStream(s, unsafe.Pointer(uintptr(p)+d.offset))
}

package types

import (
    "unsafe"
)

type structFieldSet struct {
    dec         Codec
    offset      uintptr
    isTaggedKey bool
    fieldIdx    int
    err         error
}

type structType struct {
    fieldVec           []*structFieldSet
    fieldUniqueNameNum int
    stringDecoder      *stringType
    structName         string
    fieldName          string
    isTriedOptimize    bool
    keyBitmapUint8     [][256]uint8
    keyBitmapUint16    [][256]uint16
    sortedFieldSets    []*structFieldSet
}

func newStructDecoder(structName, fieldName string, fieldVec []*structFieldSet) *structType {
    return &structType{
        fieldVec:      fieldVec,
        stringDecoder: newStringDecoder(structName, fieldName),
        structName:    structName,
        fieldName:     fieldName,
    }
}

func (d *structType) DecodeStream(s *Stream, p unsafe.Pointer) error {
    for _, v := range d.fieldVec {
        s.reset()
        if err := v.dec.DecodeStream(s, unsafe.Pointer(uintptr(p)+v.offset)); err != nil {
            return err
        }
    }
    return nil
}

func (d *structType) EncodeStream(s *Stream, p unsafe.Pointer) error {
    for _, v := range d.fieldVec {
        s.reset()
        if err := v.dec.EncodeStream(s, *(*unsafe.Pointer)(unsafe.Pointer(uintptr(p) + v.offset))); err != nil {
            return err
        }
    }
    return nil
}

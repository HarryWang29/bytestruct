package types

import (
    "encoding/hex"
    "fmt"
    "github.com/HarryWang29/bytestruct/errors"
    "github.com/HarryWang29/bytestruct/runtime"
    "reflect"
    "unsafe"
)

type ptrType struct {
    dec        Codec
    typ        *runtime.Type
    structName string
    fieldName  string
}

func newPtrDecoder(dec Codec, typ *runtime.Type, structName, fieldName string) *ptrType {
    return &ptrType{
        dec:        dec,
        typ:        typ,
        structName: structName,
        fieldName:  fieldName,
    }
}

func (d *ptrType) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
    return &errors.UnmarshalTypeError{
        Value:  fmt.Sprintf("ptr %s", hex.EncodeToString(buf)),
        Type:   reflect.TypeOf(""),
        Offset: offset,
        Struct: d.structName,
        Field:  d.fieldName,
    }
}

func (d *ptrType) contentDecoder() Codec {
    dec, ok := d.dec.(*ptrType)
    if !ok {
        return d.dec
    }
    return dec.contentDecoder()
}

func (d *ptrType) DecodeStream(s *Stream, p unsafe.Pointer) error {
    var newptr unsafe.Pointer
    if *(*unsafe.Pointer)(p) == nil {
        newptr = unsafe_New(d.typ)
        *(*unsafe.Pointer)(p) = newptr
    } else {
        newptr = *(*unsafe.Pointer)(p)
    }
    err := d.dec.DecodeStream(s, newptr)
    if err != nil {
        return err
    }
    return nil
}

func (d *ptrType) EncodeStream(stream *Stream, p unsafe.Pointer) error {
    if *(*unsafe.Pointer)(p) == nil {
        return nil
    }
    return d.dec.EncodeStream(stream, *(*unsafe.Pointer)(p))
}

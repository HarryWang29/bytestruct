package types

import (
    "encoding/hex"
    "fmt"
    "github.com/HarryWang29/bytestruct/errors"
    "reflect"
    "unsafe"
)

type stringType struct {
    structName string
    fieldName  string
}

func newStringDecoder(structName, fieldName string) *stringType {
    return &stringType{
        structName: structName,
        fieldName:  fieldName,
    }
}

func (d *stringType) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
    return &errors.UnmarshalTypeError{
        Value:  fmt.Sprintf("string %s", hex.EncodeToString(buf)),
        Type:   reflect.TypeOf(""),
        Offset: offset,
        Struct: d.structName,
        Field:  d.fieldName,
    }
}

func (d *stringType) DecodeStream(s *Stream, p unsafe.Pointer) error {
    bytes := s.buf[s.cursor:]
    str, c, err := ParseString(s.Option, bytes)
    if err != nil {
        return d.typeError(bytes, s.totalOffset())
    }
    s.cursor += c
    **(**string)(unsafe.Pointer(&p)) = *(*string)(unsafe.Pointer(&str))
    s.reset()
    return nil
}

func (d *stringType) EncodeStream(stream *Stream, pointer unsafe.Pointer) error {
    return PutString(stream.Option, &stream.w, *(*string)(pointer))
}

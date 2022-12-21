package types

import (
    "unsafe"
)

type boolType struct {
    structName string
    fieldName  string
}

func newBoolDecoder(structName, fieldName string, ) *boolType {
    return &boolType{
        structName: structName,
        fieldName:  fieldName,
    }
}

func (d *boolType) DecodeStream(s *Stream, p unsafe.Pointer) error {
    bytes := s.buf[s.cursor:]
    u, c, err := ParseBool(s.Option, bytes)
    if err != nil {
        return err
    }
    s.cursor += c
    if u == 0 {
        **(**bool)(unsafe.Pointer(&p)) = false
    } else {
        **(**bool)(unsafe.Pointer(&p)) = true
    }
    return nil
}

func (d *boolType) EncodeStream(stream *Stream, pointer unsafe.Pointer) error {
    return PutBool(stream.Option, &stream.w, *(*bool)(pointer))
}

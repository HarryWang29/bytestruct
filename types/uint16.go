package types

import (
    "encoding/hex"
    "fmt"
    "github.com/HarryWang29/bytestruct/errors"
    "reflect"
    "unsafe"
)

type uint16Type struct {
    structName string
    fieldName  string
}

func newUint16Decoder(structName, fieldName string) *uint16Type {
    return &uint16Type{
        structName: structName,
        fieldName:  fieldName,
    }
}

func (d *uint16Type) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
    return &errors.UnmarshalTypeError{
        Value:  fmt.Sprintf("uint16 %s", hex.EncodeToString(buf)),
        Type:   reflect.TypeOf(uint16(0)),
        Offset: offset,
    }
}

func (d *uint16Type) DecodeStream(s *Stream, p unsafe.Pointer) error {
    bytes := s.buf[s.cursor:]
    u16, c, err := ParseUint16(s.Option, bytes)
    if err != nil {
        return d.typeError(bytes, s.totalOffset())
    }
    s.cursor += c
    *(*uint16)(p) = u16
    return nil
}

func (d *uint16Type) EncodeStream(stream *Stream, pointer unsafe.Pointer) error {
    return PutUint16(stream.Option, &stream.w, *(*uint16)(pointer))
}

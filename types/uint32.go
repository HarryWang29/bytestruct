package types

import (
    "encoding/hex"
    "fmt"
    "github.com/HarryWang29/bytestruct/errors"
    "reflect"
    "unsafe"
)

type uint32Type struct {
    structName string
    fieldName  string
}

func newUint32Decoder(structName, fieldName string) *uint32Type {
    return &uint32Type{
        structName: structName,
        fieldName:  fieldName,
    }
}

func (d *uint32Type) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
    return &errors.UnmarshalTypeError{
        Value:  fmt.Sprintf("uint32 %s", hex.EncodeToString(buf)),
        Type:   reflect.TypeOf(uint32(0)),
        Offset: offset,
    }
}

func (d *uint32Type) DecodeStream(s *Stream, p unsafe.Pointer) error {
    bytes := s.buf[s.cursor:]
    u32, c, err := ParseUint32(s.Option, bytes)
    if err != nil {
        return d.typeError(bytes, s.totalOffset())
    }
    s.cursor += c
    *(*uint32)(p) = u32
    return nil
}

func (d *uint32Type) EncodeStream(stream *Stream, pointer unsafe.Pointer) error {
    return PutUint32(stream.Option, &stream.w, *(*uint32)(pointer))
}

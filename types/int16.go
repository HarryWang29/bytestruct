package types

import (
    "encoding/hex"
    "fmt"
    "github.com/HarryWang29/bytestruct/errors"
    "reflect"
    "unsafe"
)

type int16Type struct {
    structName string
    fieldName  string
}

func newInt16Decoder(structName, fieldName string) *int16Type {
    return &int16Type{
        structName: structName,
        fieldName:  fieldName,
    }
}

func (d *int16Type) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
    return &errors.UnmarshalTypeError{
        Value:  fmt.Sprintf("int %s", hex.EncodeToString(buf)),
        Type:   reflect.TypeOf(int16(0)),
        Offset: offset,
    }
}

func (d *int16Type) DecodeStream(s *Stream, p unsafe.Pointer) error {
    bytes := s.buf[s.cursor:]
    u, c, err := ParseInt16(s.Option, bytes)
    if err != nil {
        return d.typeError(bytes, s.totalOffset())
    }
    s.cursor += c
    *(*int16)(p) = u
    return nil
}

func (d *int16Type) EncodeStream(stream *Stream, pointer unsafe.Pointer) error {
    return PutInt16(stream.Option, &stream.w, *(*int16)(pointer))
}

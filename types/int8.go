package types

import (
    "encoding/hex"
    "fmt"
    "github.com/HarryWang29/bytestruct/errors"
    "reflect"
    "unsafe"
)

type int8Type struct {
    structName string
    fieldName  string
}

func newInt8Decoder(structName, fieldName string) *int8Type {
    return &int8Type{
        structName: structName,
        fieldName:  fieldName,
    }
}

func (d *int8Type) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
    return &errors.UnmarshalTypeError{
        Value:  fmt.Sprintf("int %s", hex.EncodeToString(buf)),
        Type:   reflect.TypeOf(int8(0)),
        Offset: offset,
    }
}

func (d *int8Type) DecodeStream(s *Stream, p unsafe.Pointer) error {
    bytes := s.buf[s.cursor:]
    u, c, err := ParseInt8(s.Option, bytes)
    if err != nil {
        return d.typeError(bytes, s.totalOffset())
    }
    s.cursor += c
    *(*int8)(p) = u
    return nil
}

func (d *int8Type) EncodeStream(stream *Stream, pointer unsafe.Pointer) error {
    return PutInt8(stream.Option, &stream.w, *(*int8)(pointer))
}

package types

import (
    "encoding/hex"
    "fmt"
    "github.com/HarryWang29/bytestruct/errors"
    "reflect"
    "unsafe"
)

type int32Type struct {
    structName string
    fieldName  string
}

func newInt32Decoder(structName, fieldName string) *int32Type {
    return &int32Type{
        structName: structName,
        fieldName:  fieldName,
    }
}

func (d *int32Type) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
    return &errors.UnmarshalTypeError{
        Value:  fmt.Sprintf("int %s", hex.EncodeToString(buf)),
        Type:   reflect.TypeOf(int32(0)),
        Offset: offset,
    }
}

func (d *int32Type) DecodeStream(s *Stream, p unsafe.Pointer) error {
    bytes := s.buf[s.cursor:]
    u, c, err := ParseInt32(s.Option, bytes)
    if err != nil {
        return d.typeError(bytes, s.totalOffset())
    }
    s.cursor += c
    *(*int32)(p) = u
    return nil
}

func (d *int32Type) EncodeStream(stream *Stream, pointer unsafe.Pointer) error {
    return PutInt32(stream.Option, &stream.w, *(*int32)(pointer))
}

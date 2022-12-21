package types

import (
    "encoding/hex"
    "fmt"
    "github.com/HarryWang29/bytestruct/errors"
    "reflect"
    "unsafe"
)

type int64Type struct {
    structName string
    fieldName  string
}

func newInt64Decoder(structName, fieldName string) *int64Type {
    return &int64Type{
        structName: structName,
        fieldName:  fieldName,
    }
}

func (d *int64Type) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
    return &errors.UnmarshalTypeError{
        Value:  fmt.Sprintf("int %s", hex.EncodeToString(buf)),
        Type:   reflect.TypeOf(int64(0)),
        Offset: offset,
    }
}

func (d *int64Type) DecodeStream(s *Stream, p unsafe.Pointer) error {
    bytes := s.buf[s.cursor:]
    u, c, err := ParseInt64(s.Option, bytes)
    if err != nil {
        return d.typeError(bytes, s.totalOffset())
    }
    s.cursor += c
    *(*int64)(p) = u
    return nil
}

func (d *int64Type) EncodeStream(stream *Stream, pointer unsafe.Pointer) error {
    return PutInt64(stream.Option, &stream.w, *(*int64)(pointer))
}

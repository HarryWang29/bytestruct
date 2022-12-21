package types

import (
    "encoding/hex"
    "fmt"
    "github.com/HarryWang29/bytestruct/errors"
    "reflect"
    "unsafe"
)

type bytesType struct {
    structName string
    fieldName  string
}

func newBytesDecoder(structName, fieldName string) *bytesType {
    return &bytesType{
        structName: structName,
        fieldName:  fieldName,
    }
}

func (d *bytesType) typeError(buf []byte, offset int64) *errors.UnmarshalTypeError {
    return &errors.UnmarshalTypeError{
        Value:  fmt.Sprintf("bytes %s", hex.EncodeToString(buf)),
        Type:   reflect.TypeOf(""),
        Offset: offset,
        Struct: d.structName,
        Field:  d.fieldName,
    }
}

func (d *bytesType) DecodeStream(s *Stream, p unsafe.Pointer) error {
    bytes, c, err := ParseBytes(s.Option, s.buf[s.cursor:])
    if err != nil {
        return d.typeError(bytes, s.totalOffset())
    }
    s.cursor += c
    *(*[]byte)(p) = bytes
    s.reset()
    return nil
}

func (d *bytesType) EncodeStream(stream *Stream, pointer unsafe.Pointer) error {
    return PutBytes(stream.Option, &stream.w, *(*[]byte)(pointer))
}

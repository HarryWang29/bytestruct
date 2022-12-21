package types

import (
    "bytes"
    "encoding/binary"
    "fmt"
)

var (
    PutUint8   = put
    PutUint16  = put
    PutUint32  = put
    PutUint64  = put
    PutInt8    = put
    PutInt16   = put
    PutInt32   = put
    PutInt64   = put
    PutFloat32 = put
    PutFloat64 = put
    PutBool    = put
    PutString  = putString
    PutBytes   = putBytes
)

func put(opt *Option, buf *bytes.Buffer, data interface{}) error {
    return binary.Write(buf, opt.Order, data)
}

func putString(opt *Option, buf *bytes.Buffer, data string) error {
    if opt.Flags&OptionFlagsString2Zero != 0 {
        return putBytes2Zero(opt, buf, []byte(data))
    }
    return putLenBytes(opt, buf, []byte(data))
}

func putLenBytes(opt *Option, buf *bytes.Buffer, data []byte) error {
    if err := put(opt, buf, len(data)); err != nil {
        return fmt.Errorf("put len err: %v", err)
    }
    if err := binary.Write(buf, opt.Order, data); err != nil {
        return fmt.Errorf("put string err: %v", err)
    }
    return nil
}

func putBytes2Zero(opt *Option, buf *bytes.Buffer, data []byte) error {
    if err := binary.Write(buf, opt.Order, data); err != nil {
        return fmt.Errorf("put string err: %v", err)
    }
    if err := binary.Write(buf, opt.Order, uint8(0)); err != nil {
        return fmt.Errorf("put zero err: %v", err)
    }
    return nil
}

func putBytes(opt *Option, buf *bytes.Buffer, data []byte) error {
    if opt.Flags&OptionFlagsBytes2Zero != 0 {
        return putBytes2Zero(opt, buf, data)
    }
    return putLenBytes(opt, buf, data)
}

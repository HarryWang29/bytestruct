package bytestruct

import (
    "bytes"
    "encoding/binary"
    "encoding/hex"
    "github.com/HarryWang29/bytestruct/runtime"
    "github.com/HarryWang29/bytestruct/types"
    "io"
    "unsafe"
)

type ByteStruct struct {
    s     *types.Stream
    order binary.ByteOrder
    debug bool
}

func NewFromBytes(data []byte, order binary.ByteOrder, debug bool) *ByteStruct {
    return NewReader(bytes.NewReader(data), order, debug)
}

func NewFromHexString(data string, order binary.ByteOrder, debug bool) (*ByteStruct, error) {
    decodeString, err := hex.DecodeString(data)
    if err != nil {
        return nil, err
    }
    return NewReader(bytes.NewReader(decodeString), order, debug), nil
}

func NewWriter(order binary.ByteOrder, debug bool) *ByteStruct {
    bs := &ByteStruct{
        s:     types.NewWriter(),
        order: order,
        debug: debug,
    }
    bs.s.Option = &types.Option{
        Order: order,
    }
    return bs
}

func NewReader(r io.ReadSeeker, order binary.ByteOrder, debug bool) *ByteStruct {
    bs := &ByteStruct{
        s:     types.NewReader(r),
        order: order,
        debug: debug,
    }
    bs.s.Option = &types.Option{
        Order: order,
    }
    return bs
}

func (r *ByteStruct) SetDebug(debug bool) {
    r.debug = debug
}

type emptyInterface struct {
    typ *runtime.Type
    ptr unsafe.Pointer
}

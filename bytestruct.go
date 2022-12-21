package bytestruct

import (
    "encoding/binary"
    "github.com/HarryWang29/bytestruct/types"
)

func MarshalLE(v interface{}) ([]byte, error) {
    return Marshal(binary.LittleEndian, v)
}

func UnmarshalLE(data []byte, v interface{}) error {
    return Unmarshal(data, binary.LittleEndian, v)
}

func MarshalBE(v interface{}) ([]byte, error) {
    return Marshal(binary.BigEndian, v)
}

func UnmarshalBE(data []byte, v interface{}) error {
    return Unmarshal(data, binary.BigEndian, v)
}

func Marshal(order binary.ByteOrder, v interface{}) ([]byte, error) {
    return NewWriter(order, false).Marshal(v)
}

func Unmarshal(data []byte, order binary.ByteOrder, v interface{}) error {
    return NewFromBytes(data, order, false).Unmarshal(v)
}

func SetParseUint8(f func(*types.Option, []byte) (uint8, int64, error)) {
    types.ParseUint8 = f
}
func SetParseUint16(f func(*types.Option, []byte) (uint16, int64, error)) {
    types.ParseUint16 = f
}
func SetParseUint32(f func(*types.Option, []byte) (uint32, int64, error)) {
    types.ParseUint32 = f
}
func SetParseUint64(f func(*types.Option, []byte) (uint64, int64, error)) {
    types.ParseUint64 = f
}
func SetParseInt8(f func(*types.Option, []byte) (int8, int64, error)) {
    types.ParseInt8 = f
}
func SetParseInt16(f func(*types.Option, []byte) (int16, int64, error)) {
    types.ParseInt16 = f
}
func SetParseInt32(f func(*types.Option, []byte) (int32, int64, error)) {
    types.ParseInt32 = f
}
func SetParseInt64(f func(*types.Option, []byte) (int64, int64, error)) {
    types.ParseInt64 = f
}
func SetParseBool(f func(*types.Option, []byte) (uint8, int64, error)) {
    types.ParseBool = f
}
func SetParseString(f func(*types.Option, []byte) ([]byte, int64, error)) {
    types.ParseString = f
}
func SetParseBytes(f func(*types.Option, []byte) ([]byte, int64, error)) {
    types.ParseBytes = f
}

type OptionFunc func(option *types.Option)

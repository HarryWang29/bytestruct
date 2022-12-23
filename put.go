package bytestruct

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/HarryWang29/bytestruct/types"
)

func toI64(data interface{}) (int64, error) {
	var i64 int64
	switch data.(type) {
	case int8:
		i64 = int64(data.(int8))
	case int16:
		i64 = int64(data.(int16))
	case int32:
		i64 = int64(data.(int32))
	case int64:
		i64 = data.(int64)
	case uint8:
		i64 = int64(data.(uint8))
	case uint16:
		i64 = int64(data.(uint16))
	case uint32:
		i64 = int64(data.(uint32))
	case uint64:
		i64 = int64(data.(uint64))
	default:
		return 0, errors.New("error type")
	}
	return i64, nil
}

func toU64(data interface{}) (uint64, error) {
	var u64 uint64
	switch data.(type) {
	case int8:
		u64 = uint64(data.(int8))
	case int16:
		u64 = uint64(data.(int16))
	case int32:
		u64 = uint64(data.(int32))
	case int64:
		u64 = uint64(data.(int64))
	case uint8:
		u64 = uint64(data.(uint8))
	case uint16:
		u64 = uint64(data.(uint16))
	case uint32:
		u64 = uint64(data.(uint32))
	case uint64:
		u64 = data.(uint64)
	default:
		return 0, errors.New("error type")
	}
	return u64, nil
}

func PutVarint(opt *types.Option, buf *bytes.Buffer, data interface{}) error {
	i64, err := toI64(data)
	if err != nil {
		return fmt.Errorf("PutVarintZigzag error: %w", err)
	}
	bs := make([]byte, 8)
	n := binary.PutVarint(bs, i64)
	buf.Write(bs[:n])
	return nil
}

func PutUvarint(opt *types.Option, buf *bytes.Buffer, data interface{}) error {
	u64, err := toU64(data)
	if err != nil {
		return fmt.Errorf("PutUvarintZigzag error: %w", err)
	}

	bs := make([]byte, 8)
	n := binary.PutUvarint(bs, u64)
	buf.Write(bs[:n])
	return nil
}

func PutVarintZigzag(opt *types.Option, buf *bytes.Buffer, data interface{}) error {
	i64, err := toI64(data)
	if err != nil {
		return fmt.Errorf("PutVarintZigzag error: %w", err)
	}

	i64 = (i64 << 1) ^ (i64 >> 63)
	return PutVarint(opt, buf, i64)
}

func PutUvarintZigzag(opt *types.Option, buf *bytes.Buffer, data interface{}) error {
	u64, err := toU64(data)
	if err != nil {
		return fmt.Errorf("PutUvarintZigzag error: %w", err)
	}
	u64 = (u64 << 1) ^ (u64 >> 63)
	return PutUvarint(opt, buf, u64)
}

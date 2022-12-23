package bytestruct

import (
	"encoding/binary"
	"fmt"
	"github.com/HarryWang29/bytestruct/types"
	"io"
)

func ParseUvarint(opt *types.Option, buf io.Reader) (uint64, error) {
	uvarint, err := binary.ReadUvarint(buf.(io.ByteReader))
	if err != nil && err != io.EOF {
		return 0, fmt.Errorf("parse uint32 error: %w", err)
	}
	return uvarint, nil
}

func ParseVarint(opt *types.Option, buf io.Reader) (uint64, error) {
	varint, err := binary.ReadVarint(buf.(io.ByteReader))
	if err != nil && err != io.EOF {
		return 0, fmt.Errorf("parse uint32 error: %w", err)
	}
	return uint64(varint), nil
}

func ParseUvarintZigzag(opt *types.Option, buf io.Reader) (uint64, error) {
	uvarint, err := ParseUvarint(opt, buf)
	if err != nil {
		return 0, err
	}
	zig := int64(uint64(uvarint)>>1) ^ int64(-(uint64(uvarint) & 1))
	return uint64(zig), nil
}

func ParseVarintZigzag(opt *types.Option, buf io.Reader) (int64, error) {
	varint, err := ParseVarint(opt, buf)
	if err != nil {
		return 0, err
	}
	zig := int64(uint64(varint)>>1) ^ int64(-(uint64(varint) & 1))
	return int64(zig), nil
}

package types

import (
	"context"
	"encoding/binary"
)

const (
	OptionFlagsString2Zero = 1 << iota // 1 for string to zero, 0 for len string
	OptionFlagsBytes2Zero              // 1 for bytes to zero, 0 for len string
)

func SetString2Zero(option *Option) {
	option.Flags |= OptionFlagsString2Zero
}

func SetBytes2Zero(option Option) {
	option.Flags |= OptionFlagsBytes2Zero
}

type OptionFlags uint8

type Option struct {
	Flags   OptionFlags
	Context context.Context
	Order   binary.ByteOrder
}

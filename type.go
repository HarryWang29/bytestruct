package bytestruct

import (
	"encoding/binary"
	"github.com/HarryWang29/bytestruct/runtime"
	"github.com/HarryWang29/bytestruct/types"
	"unsafe"
)

type ByteStruct struct {
	s      *types.Stream
	order  binary.ByteOrder
	option *types.Option
}

func NewWriter(order binary.ByteOrder, optionFunc ...OptionFunc) *ByteStruct {
	bs := &ByteStruct{
		order: order,
		option: &types.Option{
			Order: order,
		},
	}
	for _, o := range optionFunc {
		o(bs.option)
	}
	return bs
}

func NewReader(order binary.ByteOrder, optionFunc ...OptionFunc) *ByteStruct {
	bs := &ByteStruct{
		order: order,
		option: &types.Option{
			Order: order,
		},
	}
	for _, o := range optionFunc {
		o(bs.option)
	}
	return bs
}

type emptyInterface struct {
	typ *runtime.Type
	ptr unsafe.Pointer
}

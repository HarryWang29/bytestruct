//go:build !race
// +build !race

package types

import (
    "errors"
    "github.com/HarryWang29/bytestruct/runtime"
    "unsafe"
)

func CompileToGetDecoder(typ *runtime.Type) (Codec, error) {
    typeptr := uintptr(unsafe.Pointer(typ))
    if typeptr > typeAddr.MaxTypeAddr {
        return nil, errors.New("type address is too large")
    }

    index := (typeptr - typeAddr.BaseTypeAddr) >> typeAddr.AddrShift
    if dec := cachedDecoder[index]; dec != nil {
        return dec, nil
    }

    dec, err := compileHead(typ, map[uintptr]Codec{})
    if err != nil {
        return nil, err
    }
    cachedDecoder[index] = dec
    return dec, nil
}

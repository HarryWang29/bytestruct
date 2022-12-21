//go:build race
// +build race

package types

import (
    "errors"
    "github.com/HarryWang29/bytestruct/runtime"
    "sync"
    "unsafe"
)

var decMu sync.RWMutex

func CompileToGetDecoder(typ *runtime.Type) (Codec, error) {
    typeptr := uintptr(unsafe.Pointer(typ))
    if typeptr > typeAddr.MaxTypeAddr {
        return nil, errors.New("type address is too large")
    }

    index := (typeptr - typeAddr.BaseTypeAddr) >> typeAddr.AddrShift
    decMu.RLock()
    if dec := cachedDecoder[index]; dec != nil {
        decMu.RUnlock()
        return dec, nil
    }
    decMu.RUnlock()

    dec, err := compileHead(typ, map[uintptr]Codec{})
    if err != nil {
        return nil, err
    }
    decMu.Lock()
    cachedDecoder[index] = dec
    decMu.Unlock()
    return dec, nil
}

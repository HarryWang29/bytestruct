package types

import (
    "github.com/HarryWang29/bytestruct/errors"
    "math"
)

var (
    ParseUint8   = parseUint8
    ParseUint16  = parseUint16
    ParseUint32  = parseUint32
    ParseUint64  = parseUint64
    ParseInt8    = parseInt8
    ParseInt16   = parseInt16
    ParseInt32   = parseInt32
    ParseInt64   = parseInt64
    ParseFloat32 = parseFloat32
    ParseFloat64 = parseFloat64
    ParseBool    = parseUint8
    ParseString  = parseString
    ParseBytes   = parseBytes
)

func parseUint8(opt *Option, buf []byte) (uint8, int64, error) {
    var size int64 = 1
    if len(buf) < int(size) {
        return 0, 0, errors.ErrUnexpectedEndOfJSON("number(unsigned integer)", 0)
    }
    u := buf[0]

    return u, size, nil
}

func parseUint16(opt *Option, buf []byte) (uint16, int64, error) {
    var size int64 = 2
    if len(buf) < int(size) {
        return 0, 0, errors.ErrUnexpectedEndOfJSON("number(unsigned integer)", 0)
    }
    u := opt.Order.Uint16(buf[:size])

    return u, size, nil
}

func parseUint32(opt *Option, buf []byte) (uint32, int64, error) {
    var size int64 = 4
    if len(buf) < int(size) {
        return 0, 0, errors.ErrUnexpectedEndOfJSON("number(unsigned integer)", 0)
    }
    u := opt.Order.Uint32(buf[:size])

    return u, size, nil
}

func parseUint64(opt *Option, buf []byte) (uint64, int64, error) {
    var size int64 = 8
    if len(buf) < int(size) {
        return 0, 0, errors.ErrUnexpectedEndOfJSON("number(unsigned integer)", 0)
    }
    u := opt.Order.Uint64(buf[:size])

    return u, size, nil
}

func parseInt8(opt *Option, buf []byte) (int8, int64, error) {
    u, size, err := parseUint8(opt, buf)
    return int8(u), size, err
}
func parseInt16(opt *Option, buf []byte) (int16, int64, error) {
    u, size, err := parseUint16(opt, buf)
    return int16(u), size, err
}

func parseInt32(opt *Option, buf []byte) (int32, int64, error) {
    u, size, err := parseUint32(opt, buf)
    return int32(u), size, err
}

func parseInt64(opt *Option, buf []byte) (int64, int64, error) {
    u, size, err := parseUint64(opt, buf)
    return int64(u), size, err
}

func parseFloat32(opt *Option, buf []byte) (float32, int64, error) {
    u, size, err := parseUint32(opt, buf)
    return math.Float32frombits(u), size, err
}

func parseFloat64(opt *Option, buf []byte) (float64, int64, error) {
    u, size, err := parseUint64(opt, buf)
    return math.Float64frombits(u), size, err
}

func parseString(opt *Option, buf []byte) ([]byte, int64, error) {
    if opt.Flags&OptionFlagsString2Zero != 0 {
        return parseString2Zero(opt, buf)
    }
    return parseLenString(opt, buf)
}

func parseLenString(opt *Option, buf []byte) ([]byte, int64, error) {
    strLen, size, err := parseUint8(opt, buf)
    if err != nil {
        return nil, 0, err
    }
    return buf[size : size+int64(strLen)], size + int64(strLen), nil
}

func parseString2Zero(opt *Option, buf []byte) ([]byte, int64, error) {
    var i int64
    for i = 0; i < int64(len(buf)); i++ {
        if buf[i] == 0 {
            return buf[:i], i + 1, nil
        }
    }
    return nil, 0, errors.ErrUnexpectedEndOfJSON("string", 0)
}

func parseBytes(opt *Option, buf []byte) ([]byte, int64, error) {
    if opt.Flags&OptionFlagsBytes2Zero != 0 {
        return parseString2Zero(opt, buf)
    }
    return parseLenString(opt, buf)
}

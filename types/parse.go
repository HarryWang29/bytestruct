package types

import (
	"io"
	"math"
)

var (
	ParseUint8     = parseUint8
	ParseUint16    = parseUint16
	ParseUint32    = parseUint32
	ParseUint64    = parseUint64
	ParseInt8      = parseInt8
	ParseInt16     = parseInt16
	ParseInt32     = parseInt32
	ParseInt64     = parseInt64
	ParseFloat32   = parseFloat32
	ParseFloat64   = parseFloat64
	ParseBool      = parseUint8
	ParseString    = parseString
	ParseStringLen = parseUint32
	ParseSliceLen  = parseUint32
	ParseMapLen    = parseUint32
	ParseBytes     = parseBytes
	ReadBytes      = readBytes
)

func readBytes(r io.Reader, n int64) ([]byte, int64, error) {
	b := make([]byte, n)
	an, err := io.ReadFull(r, b)
	if err != nil && err != io.EOF {
		return nil, 0, err
	}
	return b, int64(an), nil
}

func parseUint8(opt *Option, buf io.Reader) (uint64, error) {
	bytes, _, err := readBytes(buf, 1)
	if err != nil {
		return 0, err
	}
	u := bytes[0]
	return uint64(u), nil
}

func parseUint16(opt *Option, buf io.Reader) (uint64, error) {
	bytes, _, err := readBytes(buf, 2)
	if err != nil {
		return 0, err
	}
	u := opt.Order.Uint16(bytes)

	return uint64(u), nil
}

func parseUint32(opt *Option, buf io.Reader) (uint64, error) {
	bytes, _, err := readBytes(buf, 4)
	if err != nil {
		return 0, err
	}
	u := opt.Order.Uint32(bytes)

	return uint64(u), nil
}

func parseUint64(opt *Option, buf io.Reader) (uint64, error) {
	bytes, _, err := readBytes(buf, 8)
	if err != nil {
		return 0, err
	}
	u := opt.Order.Uint64(bytes)

	return u, nil
}

func parseInt8(opt *Option, buf io.Reader) (uint64, error) {
	u, err := parseUint8(opt, buf)
	return u, err
}
func parseInt16(opt *Option, buf io.Reader) (uint64, error) {
	u, err := parseUint16(opt, buf)
	return u, err
}

func parseInt32(opt *Option, buf io.Reader) (uint64, error) {
	u, err := parseUint32(opt, buf)
	return u, err
}

func parseInt64(opt *Option, buf io.Reader) (uint64, error) {
	u, err := parseUint64(opt, buf)
	return u, err
}

func parseFloat32(opt *Option, buf io.Reader) (float32, error) {
	u, err := parseUint32(opt, buf)
	return math.Float32frombits(uint32(u)), err
}

func parseFloat64(opt *Option, buf io.Reader) (float64, error) {
	u, err := parseUint64(opt, buf)
	return math.Float64frombits(u), err
}

func parseString(opt *Option, buf io.Reader) (string, error) {
	if opt.Flags&OptionFlagsString2Zero != 0 {
		return parseString2Zero(opt, buf)
	}
	return parseLenString(opt, buf)
}

func parseLenString(opt *Option, buf io.Reader) (string, error) {
	bytes, err := parseLenBytes(opt, buf)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseString2Zero(opt *Option, buf io.Reader) (string, error) {
	zero, err := parseBytes2Zero(opt, buf)
	if err != nil {
		return "", err
	}
	return string(zero), nil
}

func parseBytes(opt *Option, buf io.Reader) ([]byte, error) {
	if opt.Flags&OptionFlagsBytes2Zero != 0 {
		return parseBytes2Zero(opt, buf)
	}
	return parseLenBytes(opt, buf)
}

func parseLenBytes(opt *Option, buf io.Reader) ([]byte, error) {
	strLen, err := ParseStringLen(opt, buf)
	if err != nil {
		return nil, err
	}
	bytes, _, err := readBytes(buf, int64(strLen))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func parseBytes2Zero(opt *Option, buf io.Reader) ([]byte, error) {
	ret := make([]byte, 0)
	b := make([]byte, 1)
	for {
		_, err := buf.Read(b)
		if err != nil {
			return nil, err
		}
		if b[0] == 0 {
			return ret, nil
		}
		ret = append(ret, b[0])
	}
}

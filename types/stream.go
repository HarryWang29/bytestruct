package types

import (
    "bytes"
    "io"
    "unsafe"
)

const (
    initBufSize = 512
)

type Stream struct {
    buf                   []byte
    bufSize               int64
    length                int64
    r                     io.Reader
    w                     bytes.Buffer
    offset                int64
    cursor                int64
    filledBuffer          bool
    allRead               bool
    UseNumber             bool
    DisallowUnknownFields bool
    Option                *Option
}

func NewWriter() *Stream {
    return &Stream{
        w:       bytes.Buffer{},
        bufSize: initBufSize,
        buf:     make([]byte, initBufSize),
        Option:  &Option{},
    }
}

func NewReader(r io.Reader) *Stream {
    return &Stream{
        r:       r,
        bufSize: initBufSize,
        buf:     make([]byte, initBufSize),
        Option:  &Option{},
    }
}

func (s *Stream) TotalOffset() int64 {
    return s.totalOffset()
}

func (s *Stream) Buffered() io.Reader {
    buflen := int64(len(s.buf))
    for i := s.cursor; i < buflen; i++ {
        if s.buf[i] == nul {
            return bytes.NewReader(s.buf[s.cursor:i])
        }
    }
    return bytes.NewReader(s.buf[s.cursor:])
}

func (s *Stream) totalOffset() int64 {
    return s.offset + s.cursor
}

func (s *Stream) char() byte {
    return s.buf[s.cursor]
}

func (s *Stream) stat() ([]byte, int64, unsafe.Pointer) {
    return s.buf, s.cursor, (*sliceHeader)(unsafe.Pointer(&s.buf)).data
}

func (s *Stream) bufptr() unsafe.Pointer {
    return (*sliceHeader)(unsafe.Pointer(&s.buf)).data
}

func (s *Stream) statForRetry() ([]byte, int64, unsafe.Pointer) {
    s.cursor-- // for retry ( because caller progress cursor position in each loop )
    return s.buf, s.cursor, (*sliceHeader)(unsafe.Pointer(&s.buf)).data
}

func (s *Stream) Reset() {
    s.reset()
    s.bufSize = int64(len(s.buf))
}

func (s *Stream) reset() {
    s.offset += s.cursor
    s.buf = s.buf[s.cursor:]
    s.length -= s.cursor
    s.cursor = 0
}

func (s *Stream) readBuf() []byte {
    if s.filledBuffer {
        s.bufSize *= 2
        remainBuf := s.buf
        s.buf = make([]byte, s.bufSize)
        copy(s.buf, remainBuf)
    }
    remainLen := s.length - s.cursor
    remainNotNulCharNum := int64(0)
    for i := int64(0); i < remainLen; i++ {
        if s.buf[s.cursor+i] == nul {
            break
        }
        remainNotNulCharNum++
    }
    s.length = s.cursor + remainNotNulCharNum
    return s.buf[s.cursor+remainNotNulCharNum:]
}

func (s *Stream) read() bool {
    if s.allRead {
        return false
    }
    buf := s.readBuf()
    last := len(buf) - 1
    buf[last] = nul
    n, err := s.r.Read(buf[:last])
    s.length += int64(n)
    if n == last {
        s.filledBuffer = true
    } else {
        s.filledBuffer = false
    }
    if err == io.EOF {
        s.allRead = true
    } else if err != nil {
        return false
    }
    return true
}

func (s *Stream) PrepareForDecode() error {
    if s.read() {
        return nil
    }
    return io.EOF
}

func (s *Stream) Bytes() []byte {
    return s.w.Bytes()
}

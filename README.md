# bytestruct
此项目受到[binstruct](https://github.com/ghostiam/binstruct)与[go-json](https://github.com/goccy/go-json)启发，目标是使用`go-json`的模式进行效率的提升

通过`marshal`，可以将任意数据结构转换为字节流，然后再通过`unmarshal`，将字节流转换为原来的数据结构。

在`unmarshal`中使用了[go-json](https://github.com/goccy/go-json)的方案，采用unsafe指针赋值并且缓存将反射的类型进行缓存。

## 用法
### 结构体转换

```go
package main

import (
    "encoding/binary"
    "fmt"
    "github.com/HarryWang29/bytestruct"
)

type A struct {
    A int8
}
type B struct {
    B int16
}
type C struct {
    A  A
    AA A
    B  B
    T  int32
    Tt int32 `bytestruct:"-"`
    C  string
}

func main() {
    var v C
    s := bytestruct.NewReader(binary.LittleEndian)
    err := s.UnmarshalFromHexString("0203616201000000020000006162", &v)
    if err != nil {
        panic(err)
    }
    if v.A.A != 2 || v.AA.A != 3 || v.B.B != 25185 || v.T != 1 || v.Tt != 0 || v.C != "ab" {
        panic(fmt.Sprintf( "decode struct error: %+v", v))
    }
}
```
### 使用`varint` 或者 `varint+zigzag`
```go
package main

import (
    "encoding/binary"
    "fmt"
    "github.com/HarryWang29/bytestruct"
)

type A struct {
        A int8
    }
    type B struct {
        B int16
    }
    type C struct {
        A  A
        AA A
        B  B
        T  int32
        Tt int32 `bytestruct:"-"`
        C  string
    }

    bytestruct.SetParseStringLen(bytestruct.ParseUvarintZigzag)
    var v C
    s := bytestruct.NewReader(binary.LittleEndian)
    err := s.UnmarshalFromHexString("0203616201000000046162", &v)
    if err != nil {
       t.Fatal(err)
    }
    if v.A.A != 2 || v.AA.A != 3 || v.B.B != 25185 || v.T != 1 || v.Tt != 0 || v.C != "ab" {
        t.Fatal("decode struct error: ", v)
    }
```

### 可自定义解析列表
#### Decode
* SetParseUint8
* SetParseUint16
* SetParseUint32
* SetParseUint64
* SetParseInt8
* SetParseInt16
* SetParseInt32
* SetParseInt64
* SetParseBool
* SetParseString
* SetParseStringLen
* SetParseMapLen
* SetParseSliceLen
* SetParseBytes

#### Encode
* SetPutUint8
* SetPutUint16
* SetPutUint32
* SetPutUint64
* SetPutInt8
* SetPutInt16
* SetPutInt32
* SetPutInt64
* SetPutBool
* SetPutString
* SetPutLen
* SetPutBytes

### **特别注意**:
1. 此项目对`struct`中字段的顺序有要求，必须和字节中一一对应。
2. 基础类型中不支持`int`与`uint`类型，根据golang中描述`int`与`uint`都属于不确定长度的字段，因此不支持。
    > // int is a signed integer type that is at least 32 bits in size. It is a
    distinct type, however, and not an alias for, say, int32.
   > 
    > // uint is an unsigned integer type that is at least 32 bits in size. It is a
    distinct type, however, and not an alias for, say, uint32.
3. 目前存在一个已知bug，在解析为`map`类型时，必须将`map make`一下，否则无法赋值，但是不会影响解析后续字节，只是无法将解析出来的数字赋值到`map`中。
### tag
目前只支持最简单的`bytestruct:"-"`，表示忽略该字段。

## Features
[ ] 修复map未被make时，无法赋值的问题

[ ] 增加性能测试

[ ] 增加更多的测试用例

## 性能
todo
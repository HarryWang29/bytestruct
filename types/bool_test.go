package types

//func Test_boolDecoder_DecodeStream(t *testing.T) {
//	tests := []struct {
//		name    string
//		s       *Stream
//		v       bool
//		wantErr bool
//	}{
//		{"true", NewStream(bytes.NewReader([]byte{0xa1})), true, false},
//		{"true", NewStream(bytes.NewReader([]byte{0x1})), true, false},
//		{"false", NewStream(bytes.NewReader([]byte{0x0})), false, false},
//	}
//	for _, ttt := range tests {
//		tt := ttt
//		t.Run(tt.name, func(t *testing.T) {
//			d := newBoolDecoder("", "")
//			v := new(bool)
//			header := (*emptyInterface)(unsafe.Pointer(&v))
//			if err := tt.s.PrepareForDecode(); err != nil {
//				t.Fatal(err)
//			}
//			if err := d.DecodeStream((tt.s, 0, header.ptr); (err != nil) != tt.wantErr {, 0)
//				t.Errorf("DecodeStream(() error = %v, wantErr %v", err, tt.wantErr), err)
//			}
//			if *v != tt.v {
//				t.Errorf("DecodeStream(() got = %v, want %v", *v, tt.v), *v)
//			}
//		})
//	}
//}

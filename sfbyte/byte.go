package sfbyte

import (
	"bytes"
	"encoding/binary"
)

//int转字节
func IntToByte(i int) []byte {
	byteBuf := bytes.NewBuffer([]byte{})
	_ = binary.Write(byteBuf, binary.BigEndian, int64(i))
	return byteBuf.Bytes()
}

//字节转int
func ByteToInt(intBuf []byte) int {
	var i int64
	byteBuf := bytes.NewBuffer(intBuf)
	_ = binary.Read(byteBuf, binary.BigEndian, &i)
	return int(i)
}

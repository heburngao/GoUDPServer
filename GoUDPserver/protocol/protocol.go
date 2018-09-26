package protocol

import (
	"bytes"
	"encoding/binary"
)

const (
	Header         = "DB"
	HeaderLen   = 2
	BufferLen = 4
)

//封包
func Packet(message []byte) []byte {
	return append(append([]byte(Header), IntToBytes(len(message))...), message...)
}

//解包
func Unpack(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)

	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+HeaderLen+BufferLen {
			break
		}
		if string(buffer[i:i+HeaderLen]) == Header {
			messageLength := BytesToInt(buffer[i+HeaderLen : i+HeaderLen+BufferLen])
			if length < i+HeaderLen+BufferLen+messageLength {
				break
			}
			data := buffer[i+HeaderLen+BufferLen : i+HeaderLen+BufferLen+messageLength]
			readerChannel <- data

			i += HeaderLen + BufferLen + messageLength - 1
		}
	}

	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}


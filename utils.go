package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	ret := new(bytes.Buffer)
	err := binary.Write(ret, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return ret.Bytes()
}

// ReverseBytes reserves a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

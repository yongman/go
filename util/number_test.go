//
// number_test.go
// Copyright (C) 2018 YanMing <yming0221@gmail.com>
//
// Distributed under terms of the MIT license.
//

package util

import (
	"fmt"
	"testing"
)

func TestInt64ToBytes(t *testing.T) {
	var n int64 = 1234

	v, err := Int64ToBytes(n)
	fmt.Println(v, err)
}

func TestBytesToInt64(t *testing.T) {
	var n int64 = -1234

	v, err := Int64ToBytes(n)
	fmt.Println(v, err)

	nn, err := BytesToInt64(v)
	fmt.Println(nn, err)
}

func TestStrByteToInt64(t *testing.T) {
	var n string = "-123"

	nn, err := StrBytesToInt64([]byte(n))
	fmt.Println(nn, err)
}

func TestUint32ToBytes(t *testing.T) {
	var n uint32 = 1234
	nn, err := Uint32ToBytes(n)
	mm, _ := BytesToUint32(nn)
	fmt.Println(nn, mm, err)
}

func TestBytesToUint32(t *testing.T) {
	num := "1234"
	numBytes := []byte(num)
	nn, err := BytesToUint32(numBytes)
	fmt.Println(nn, err)
}

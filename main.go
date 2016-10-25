package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// key := [4]byte{0x12, 0x34, 0x56, 0x78}
	// cipher, _ := NewCipher(key[:])
	// p := [6]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}
	// c := make([]byte, 6)
	// cipher.Encrypt(c, p[:])
	// fmt.Printf("%02X %02X %02X %02X %02X %02X\n", c[0], c[1], c[2], c[3], c[4], c[5])

	p := [6]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB}
	q := [6]byte{0xFB, 0xD0, 0x40, 0xD6, 0xDB, 0x9C}
	c := make([]byte, 6)
	key := make([]byte, 4)
	i := uint32(0) // Around 1040000000

	for {
		if i%10000000 == 0 {
			fmt.Printf("%v/4294967295\n", i)
		}

		binary.BigEndian.PutUint32(key[:], i)
		cipher, _ := NewCipher(key)
		cipher.Encrypt(c, p[:])
		if bytes.Equal(c, q[:]) {
			break
		}

		i++
	}

	fmt.Printf("%02X %02X %02X %02X\n", key[0], key[1], key[2], key[3]) // 3E 76 AC 4B
}

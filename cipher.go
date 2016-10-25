// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto/cipher"
	"encoding/binary"
	"strconv"
)

// BlockSize is the block size in bytes
const BlockSize = 6

// A cipher is an instance of the encryption using a particular key.
type symmetricCipher struct {
	rk []byte
}

// KeySizeError is the error shown when the key size is not 32 bits
type KeySizeError int

func (k KeySizeError) Error() string {
	return "crypto/aes: invalid key size " + strconv.Itoa(int(k))
}

// NewCipher creates and returns a new cipher that can be used for encryption/decryption
func NewCipher(key []byte) (cipher.Block, error) {
	k := len(key)

	if k != 4 {
		return nil, KeySizeError(k)
	}

	rk := expandKey(binary.BigEndian.Uint32(key))
	c := symmetricCipher{rk}

	return &c, nil
}

func (c *symmetricCipher) BlockSize() int { return BlockSize }

func (c *symmetricCipher) Encrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic("crypto/aes: input not full block")
	}
	if len(dst) < BlockSize {
		panic("crypto/aes: output not full block")
	}
	encryptBlock(c.rk, dst, src)
}

func (c *symmetricCipher) Decrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic("crypto/aes: input not full block")
	}
	if len(dst) < BlockSize {
		panic("crypto/aes: output not full block")
	}
	// decryptBlock(c.rk, dst, src)
}

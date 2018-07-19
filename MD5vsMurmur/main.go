package main

import (
	"encoding/hex"

	"crypto/md5"

	"github.com/spaolacci/murmur3"
)

func MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func Murmur(data []byte) string {
	h := murmur3.New128()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

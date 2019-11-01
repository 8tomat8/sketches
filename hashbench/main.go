package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/minio/blake2b-simd"

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

func MinioBlake2b(data []byte) string {
	h := blake2b.New256()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func SHA256(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
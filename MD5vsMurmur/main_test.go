package main

import (
	"testing"

	"io/ioutil"

	"github.com/satori/go.uuid"
)

var result []string

func mkinput(n int, inputSize int) [][]byte {
	rv := make([][]byte, n)
	for i := 0; i < n; i++ {
		for j := 0; j < inputSize; j++ {
			rv[i] = append(rv[i], uuid.Must(uuid.NewV4()).Bytes()...)
		}
	}
	return rv
}

func benchmarkMD5(b *testing.B, inputSize int) {
	//input := mkinput(b.N, inputSize)
	input, err := ioutil.ReadFile("./milky_way_starry_sky_galaxy_119519_3840x2160.jpg")
	if err != nil {
		b.Fatal(err)
	}
	output := make([]string, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		output[i] = MD5(input)
	}
	result = output
}

func benchmarkMurmur(b *testing.B, inputSize int) {
	//input := mkinput(b.N, inputSize)
	input, err := ioutil.ReadFile("./milky_way_starry_sky_galaxy_119519_3840x2160.jpg")
	if err != nil {
		b.Fatal(err)
	}
	output := make([]string, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		output[i] = Murmur(input)
	}
	result = output
}

func BenchmarkMD5_16(b *testing.B)   { benchmarkMD5(b, 1) }
func BenchmarkMD5_160(b *testing.B)  { benchmarkMD5(b, 10) }
func BenchmarkMD5_1600(b *testing.B) { benchmarkMD5(b, 100) }

//func BenchmarkMD5_16000(b *testing.B) { benchmarkMD5(b, 1000) }

func BenchmarkMurmur_16(b *testing.B)   { benchmarkMurmur(b, 1) }
func BenchmarkMurmur_160(b *testing.B)  { benchmarkMurmur(b, 10) }
func BenchmarkMurmur_1600(b *testing.B) { benchmarkMurmur(b, 100) }

//func BenchmarkMurmur_16000(b *testing.B) { benchmarkMurmur(b, 1000) }

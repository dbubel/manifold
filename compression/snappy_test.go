package compression

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
)

func TestSnappyCompressor(t *testing.T) {
	sc := NewSnappy()
	data := []byte(payload)

	// Test CompressIOPool
	compressedData, err := sc.CompressIOPool(data)
	if err != nil {
		t.Errorf("Unexpected error in CompressIOPool: %v", err)
	}
	if bytes.Equal(data, compressedData) {
		t.Error("CompressIOPool did not compress data")
	}

	// Test DecompressIOPool
	decompressedData, err := sc.DecompressIOPool(compressedData)
	if err != nil {
		t.Errorf("Unexpected error in DecompressIOPool: %v", err)
	}
	if !bytes.Equal(data, decompressedData) {
		t.Error("DecompressIOPool did not correctly decompress data")
	}

	// Test CompressIOBufPool
	//compressedDataBuf, err := sc.CompressIOBufPool(data)
	//if err != nil {
	//	t.Errorf("Unexpected error in CompressIOBufPool: %v", err)
	//}
	//if bytes.Equal(data, compressedDataBuf) {
	//	t.Error("CompressIOBufPool did not compress data")
	//}
	//
	//// Test DecompressIOBufPool
	//decompressedDataBuf, err := sc.DecompressIOBufPool(compressedDataBuf)
	//if err != nil {
	//	t.Errorf("Unexpected error in DecompressIOBufPool: %v", err)
	//}
	//if !bytes.Equal(data, decompressedDataBuf) {
	//	t.Error("DecompressIOBufPool did not correctly decompress data")
	//}
	//
	//// Test CompressBufPool
	//compressedDataBuffer, err := sc.CompressBufPool(data)
	//if err != nil {
	//	t.Errorf("Unexpected error in CompressBufPool: %v", err)
	//}
	//if bytes.Equal(data, compressedDataBuffer) {
	//	t.Error("CompressBufPool did not compress data")
	//}
	//
	//// Test DecompressBufPool
	//decompressedDataBuffer, err := sc.DecompressBufPool(compressedDataBuffer)
	//if err != nil {
	//	t.Errorf("Unexpected error in DecompressBufPool: %v", err)
	//}
	//if !bytes.Equal(data, decompressedDataBuffer) {
	//	t.Error("DecompressBufPool did not correctly decompress data")
	//}
}

// characters to use to generate random string
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// GenerateRandomString generates a random string of the given length.
func GenerateRandomString(length int) []byte {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return []byte(string(s))
}

var lenth = 100000

//var data = []byte(payload)

//func BenchmarkCompressIOPool(b *testing.B) {
//	sc := NewSnappy()
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		_, _ = sc.CompressIOPool(data)
//	}
//}

func BenchmarkDecompressIOPool(b *testing.B) {
	sc := NewSnappy()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		compressedData, _ := sc.CompressIOPool(GenerateRandomString(lenth))
		_, _ = sc.DecompressIOPool(compressedData)
	}
}

//func BenchmarkCompressIOBufPool(b *testing.B) {
//	sc := NewSnappy()
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		_, _ = sc.CompressIOBufPool(data)
//	}
//}

//func BenchmarkDecompressIOBufPool(b *testing.B) {
//	sc := NewSnappy()
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		compressedData, _ := sc.CompressIOBufPool(GenerateRandomString(lenth))
//		_, _ = sc.DecompressIOBufPool(compressedData)
//	}
//}

//func BenchmarkCompressBufPool(b *testing.B) {
//	sc := NewSnappy()
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		_, _ = sc.CompressBufPool(data)
//	}
//}

//func BenchmarkDecompressBufPool(b *testing.B) {
//	sc := NewSnappy()
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		compressedData, _ := sc.CompressBufPool(GenerateRandomString(lenth))
//		_, _ = sc.DecompressBufPool(compressedData)
//	}
//}

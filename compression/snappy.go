package compression

import (
	"bytes"
	"io"
	"sync"

	"github.com/golang/snappy"
)

type SnappyCompressor struct {
	writerPool sync.Pool
	readerPool sync.Pool
	//bufferPool sync.Pool
}

func NewSnappy() *SnappyCompressor {
	return &SnappyCompressor{
		writerPool: sync.Pool{
			New: func() interface{} {
				return snappy.NewBufferedWriter(nil)
			},
		},
		readerPool: sync.Pool{
			New: func() interface{} {
				return snappy.NewReader(nil)
			},
		},
	}
}

func (sc *SnappyCompressor) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := sc.writerPool.Get().(*snappy.Writer)
	writer.Reset(&buf)

	if _, err := writer.Write(data); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	sc.writerPool.Put(writer)
	return buf.Bytes(), nil
}

func (sc *SnappyCompressor) Decompress(data []byte) ([]byte, error) {
	b := bytes.NewBuffer(data)
	reader := sc.readerPool.Get().(*snappy.Reader)
	reader.Reset(b)

	uncompressed, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	sc.readerPool.Put(reader)
	return uncompressed, nil
}

//
//func (sc *SnappyCompressor) CompressIOBufPool(data []byte) ([]byte, error) {
//	// Get a writer and buffer from the pools
//	writer := sc.writerPool.Get().(*snappy.Writer)
//	buffer := sc.bufferPool.Get().(*bytes.Buffer)
//
//	// Reset the writer and buffer to clean state
//	writer.Reset(buffer)
//	buffer.Reset()
//
//	// CompressIOPool the data
//	_, err := writer.Write(data)
//	if err != nil {
//		return nil, err
//	}
//
//	// Close the writer to ensure it has finished writing
//	err = writer.Close()
//	if err != nil {
//		return nil, err
//	}
//
//	// Put the writer back in the pool
//	sc.writerPool.Put(writer)
//
//	// Extract the compressed data from the buffer
//	compressed := buffer.Bytes()
//
//	// Put the buffer back in the pool
//	sc.bufferPool.Put(buffer)
//
//	return compressed, nil
//}
//
//func (sc *SnappyCompressor) DecompressIOBufPool(data []byte) ([]byte, error) {
//	// Get a reader from the pool
//	reader := sc.readerPool.Get().(*snappy.Reader)
//
//	// Reset the reader to read from the new data
//	reader.Reset(bytes.NewBuffer(data))
//
//	// Create a buffer to hold the decompressed data
//	var buf bytes.Buffer
//
//	// Create a byte slice as an intermediate buffer
//	tmp := make([]byte, 256)
//
//	for {
//		// Read the decompressed data
//		n, err := reader.Read(tmp)
//		if err != nil && err != io.EOF {
//			return nil, err
//		}
//		if n > 0 {
//			buf.Write(tmp[:n])
//		}
//		if err == io.EOF {
//			break
//		}
//	}
//
//	// Put the reader back in the pool
//	sc.readerPool.Put(reader)
//
//	// Return the decompressed data
//	return buf.Bytes(), nil
//}
//
//func (sc *SnappyCompressor) CompressBufPool(data []byte) ([]byte, error) {
//	buf := sc.bufferPool.Get().(*bytes.Buffer)
//	buf.Reset()
//	writer := snappy.NewBufferedWriter(buf)
//
//	if _, err := writer.Write(data); err != nil {
//		return nil, err
//	}
//	if err := writer.Close(); err != nil {
//		return nil, err
//	}
//
//	compressedData := make([]byte, buf.Len())
//	copy(compressedData, buf.Bytes())
//
//	sc.bufferPool.Put(buf)
//	return compressedData, nil
//}
//
//func (sc *SnappyCompressor) DecompressBufPool(data []byte) ([]byte, error) {
//	reader := snappy.NewReader(bytes.NewReader(data))
//
//	uncompressed, err := io.ReadAll(reader)
//	if err != nil {
//		return nil, err
//	}
//
//	return uncompressed, nil
//}

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

func (sc *SnappyCompressor) CompressIOPool(data []byte) ([]byte, error) {
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

func (sc *SnappyCompressor) DecompressIOPool(data []byte) ([]byte, error) {
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

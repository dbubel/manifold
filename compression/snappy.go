package compression

import (
	"bytes"
	"github.com/dbubel/manifold/linked_list"
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

func (sc *SnappyCompressor) Compress(data *linked_list.Element) error {
	buf := bytes.NewBuffer(data.Value)
	writer := sc.writerPool.Get().(*snappy.Writer)
	writer.Reset(buf)

	if _, err := writer.Write(buf.); err != nil {
		return err
	}
	if err := writer.Close(); err != nil {
		return err
	}

	sc.writerPool.Put(writer)
	return nil
}

func (sc *SnappyCompressor) Decompress(element *linked_list.Element) error {
	b := bytes.NewBuffer(element.Value)
	reader := sc.readerPool.Get().(*snappy.Reader)
	reader.Reset(b)

	uncompressed, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	element.Value = uncompressed
	sc.readerPool.Put(reader)
	return nil
}

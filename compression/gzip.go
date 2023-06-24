package compression

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"sync"
)

type GZIPCompressor struct {
	writerPool sync.Pool
	readerPool sync.Pool
}

func New() *GZIPCompressor {
	return &GZIPCompressor{
		writerPool: sync.Pool{
			New: func() interface{} {
				g, _ := gzip.NewWriterLevel(nil, gzip.BestSpeed)

				return g
			},
		},
		readerPool: sync.Pool{
			New: func() interface{} {
				// gzip.Reader requires an existing compressed data stream to be created,
				// so we can't create it here as we did with gzip.Writer.
				// Instead, we will create it on-the-fly in the DecompressIOPool method.
				return new(gzip.Reader)
			},
		},
	}
}

func (sc *GZIPCompressor) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := sc.writerPool.Get().(*gzip.Writer)
	gz.Reset(&buf)

	if _, err := gz.Write(data); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}

	sc.writerPool.Put(gz)
	return buf.Bytes(), nil
}

func (sc *GZIPCompressor) Decompress(data []byte) ([]byte, error) {
	b := bytes.NewBuffer(data)
	r := sc.readerPool.Get().(*gzip.Reader)
	if err := r.Reset(b); err != nil {
		return nil, err
	}

	uncompressed, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if err := r.Close(); err != nil {
		return nil, err
	}

	sc.readerPool.Put(r)
	return uncompressed, nil
}

//
//type FlateCompressor struct {
//	pool sync.Pool
//}
//
//func NewFlate() *FlateCompressor {
//	return &FlateCompressor{
//		pool: sync.Pool{
//			New: func() interface{} {
//				// Buffer creation
//				return new(bytes.Buffer)
//			},
//		},
//	}
//}
//
//func (c *FlateCompressor) CompressIOPool(data []byte, level int) ([]byte, error) {
//	buf := c.pool.Get().(*bytes.Buffer)
//	buf.Reset() // Clear the buffer
//	defer c.pool.Put(buf)
//
//	// New writer with the specified compression level
//	writer, err := flate.NewWriter(buf, level)
//	if err != nil {
//		return nil, err
//	}
//
//	if _, err := writer.Write(data); err != nil {
//		return nil, err
//	}
//
//	if err := writer.Close(); err != nil {
//		return nil, err
//	}
//
//	// Need to make a copy of the data to prevent data race
//	res := make([]byte, buf.Len())
//	copy(res, buf.Bytes())
//
//	return res, nil
//}
//
//func (c *FlateCompressor) DecompressIOPool(data []byte) ([]byte, error) {
//	r := flate.NewReader(bytes.NewReader(data))
//	defer r.Close()
//
//	return ioutil.ReadAll(r)
//}

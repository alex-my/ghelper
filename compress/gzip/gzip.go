package gzip

import (
	"bytes"
	cgzip "compress/gzip"
	"io/ioutil"

	"github.com/alex-my/ghelper/compress"
)

type gzip struct {
	level int
}

// NewGZip ..
func NewGZip(level ...int) compress.Compress {
	l := cgzip.DefaultCompression
	if len(level) > 0 {
		l = level[0]
	}

	return &gzip{
		level: l,
	}
}

// Compress 压缩
func (gzip *gzip) Compress(in []byte) ([]byte, error) {
	var buffer bytes.Buffer

	writer, err := cgzip.NewWriterLevel(&buffer, gzip.level)
	if err != nil {
		return nil, err
	}

	if _, err = writer.Write(in); err != nil {
		return nil, err
	}

	if err := writer.Flush(); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// Uncompress 解压缩
func (gzip *gzip) Uncompress(in []byte) ([]byte, error) {
	reader, err := cgzip.NewReader(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}

	defer reader.Close()

	return ioutil.ReadAll(reader)
}

// Name 获取名称
func (gzip *gzip) Name() string {
	return "gzip"
}

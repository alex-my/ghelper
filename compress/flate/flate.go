package flate

import (
	"bytes"
	cflate "compress/flate"
	"io/ioutil"

	"github.com/alex-my/ghelper/compress"
)

type flate struct {
	level int
}

// NewFlate ..
func NewFlate(level ...int) compress.Compress {
	l := cflate.DefaultCompression
	if len(level) > 0 {
		l = level[0]
	}

	return &flate{
		level: l,
	}
}

// Compress 压缩
func (flate *flate) Compress(in []byte) ([]byte, error) {
	var buffer bytes.Buffer

	writer, err := cflate.NewWriter(&buffer, flate.level)
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
func (flate *flate) Uncompress(in []byte) ([]byte, error) {
	reader := cflate.NewReader(bytes.NewReader(in))

	defer reader.Close()

	return ioutil.ReadAll(reader)
}

// Name 获取名称
func (flate *flate) Name() string {
	return "flate"
}

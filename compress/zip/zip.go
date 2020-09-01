package zip

import (
	"bytes"
	czip "compress/zlib"
	"io/ioutil"

	"github.com/alex-my/ghelper/compress"
)

type zip struct {
	level int
}

// NewZip ..
func NewZip(level ...int) compress.Compress {
	l := czip.DefaultCompression
	if len(level) > 0 {
		l = level[0]
	}

	return &zip{
		level: l,
	}
}

// Compress 压缩
func (zip *zip) Compress(in []byte) ([]byte, error) {
	var buffer bytes.Buffer

	writer, err := czip.NewWriterLevel(&buffer, zip.level)
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
func (zip *zip) Uncompress(in []byte) ([]byte, error) {
	reader, err := czip.NewReader(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}

	defer reader.Close()

	return ioutil.ReadAll(reader)
}

// Name 获取名称
func (zip *zip) Name() string {
	return "zip"
}

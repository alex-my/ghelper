package codec

import (
	"encoding/json"

	"github.com/alex-my/ghelper/codec"
)

type jsonCodec struct{}

// NewJSONCodec JSON
func NewJSONCodec() codec.Codec {
	return &jsonCodec{}
}

// Encode 编码
func (*jsonCodec) Encode(in interface{}) ([]byte, error) {
	return json.Marshal(in)
}

// Decode 解码
func (*jsonCodec) Decode(in []byte, out interface{}) error {
	return json.Unmarshal(in, out)
}

// Name 名称
func (*jsonCodec) Name() string {
	return "json"
}

// MimeType 媒体类型
func (*jsonCodec) MimeType() string {
	return "application/json"
}

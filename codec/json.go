package codec

import "encoding/json"

type jsonCodec struct {
}

// NewJSONCodec JSON
func NewJSONCodec() Codec {
	return &jsonCodec{}
}

// Encode 编码
func (*jsonCodec) Encode(in interface{}) ([]byte, error) {
	return json.Marshal(in)
}

// Unmarshal 解码
func (*jsonCodec) Decode(in []byte, out interface{}) error {
	return json.Unmarshal(in, out)
}

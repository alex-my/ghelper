package codec

// Codec 编码与解码器接口
type Codec interface {
	Encode(in interface{}) ([]byte, error)
	Decode(in []byte, out interface{}) error
}

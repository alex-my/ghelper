// Package codec 消息的编码与解码，支持 json、protobuf，默认 json 格式
package codec

// Codec 编码与解码器
type Codec interface {
	// Encode 将数据转为 []byte
	Encode(in interface{}) ([]byte, error)

	// Decode 将 []byte 转为数据
	Decode(in []byte, out interface{}) error

	// Name 名称
	Name() string

	// MimeType 媒体类型
	MimeType() string
}

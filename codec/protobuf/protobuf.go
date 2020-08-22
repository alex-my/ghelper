package codec

import (
	"errors"

	"github.com/alex-my/ghelper/codec"
	"github.com/golang/protobuf/proto"
)

var (
	// ErrInvalidPBMessage 无效的 google protobuf 消息
	ErrInvalidPBMessage = errors.New("invalid pb message")
)

type protobufCodec struct{}

// NewProtobufCodec 创建默认的编码与解码器
func NewProtobufCodec() codec.Codec {
	return &protobufCodec{}
}

// Encode 编码
func (c *protobufCodec) Encode(in interface{}) ([]byte, error) {
	m, ok := in.(proto.Message)
	if !ok {
		return nil, ErrInvalidPBMessage
	}

	return proto.Marshal(m)
}

// Decode 解码
func (c *protobufCodec) Decode(in []byte, out interface{}) error {
	m, ok := out.(proto.Message)
	if !ok {
		return ErrInvalidPBMessage
	}

	return proto.Unmarshal(in, m)
}

// Name 名称
func (*protobufCodec) Name() string {
	return "protobuf"
}

// MimeType 媒体类型
func (*protobufCodec) MimeType() string {
	return "application/binary"
}

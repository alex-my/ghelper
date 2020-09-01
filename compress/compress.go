package compress

// Compress 压缩接口
type Compress interface {
	// Compress 压缩
	Compress(in []byte) ([]byte, error)

	// Uncompress 解压缩
	Uncompress(in []byte) ([]byte, error)

	// Name 获取名称
	Name() string
}

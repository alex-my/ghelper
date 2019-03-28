package graceful

// Package 添加服务器时的选项

// Option 添加额外选项
type Option func(hs *httpServer)

// WithName 给服务器添加名称
func WithName(name string) Option {
	return func(hs *httpServer) {
		hs.name = name
	}
}

// WithTimeout 给服务器单独指定超时时间
// 未额外指定的则使用公共的超时时
// 公共超时时间可以使用 SetShutdownTimeout 设置
func WithTimeout(timeout int) Option {
	return func(hs *httpServer) {
		hs.timeout = timeout
	}
}

// WithShutdownHandler 给服务器单独添加关闭函数
// 给所有服务器添加关闭函数，可以使用 RegisterShutdownHandler
func WithShutdownHandler(f func()) Option {
	return func(hs *httpServer) {
		hs.server.RegisterOnShutdown(f)
	}
}

# 说明

- 优雅的重启和关闭服务器
- 优雅的重启需要获取原进程监听套接字的文件描述符，所以不能直接使用`http.server`的`ListenAndServe`和`ListenAndServeTLS`
- 只要将`http.Server`替换为`ghelper/graceful/server`即可
- 优雅的关闭服务器:

  - 服务器会关闭监听，执行用户注册的清理函数，等待未处理完毕的连接继续处理
  - 服务器每隔 `500` 毫秒判断一次所有连接是否处理完毕，如果处理完毕，则关闭服务器
  - 如果直到超时，都未处理完，直接关闭服务器。如果没有额外设置，默认超时时间为 `5` 秒

- 优雅的重启服务器:
  - 新进程"继承"旧进程的 `os.Stdin, os.Stdout, os.Stderr` 三个文件描述符以及 `监听套接字` 的文件描述符
  - 旧进程优雅的关闭，监听新连接的工作由新进程接手

# 示例 1

本示例使用默认信号

- 代码

  ```go
  import (
      "net/http"
      "os"

      "github.com/alex-my/ghelper/graceful"
      "github.com/alex-my/ghelper/logger"
  )

  // http.Server 要求实现 ServeHTTP
  type testServer struct {}

  func (server *testServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    // 这里进行逻辑处理，比如按照路由进行处理
    res.Write([]byte("Hello world"))
  }

  func main() {
      addr := "127.0.0.1:8877"
      server := &testServer{}
      logger := logger.NewLogger()

      // 使用 gserver 而不是 http.Server
      gserver := graceful.NewServer(server, logger)

      logger.Infof("listen on: http://%s, pid: %d", addr, os.Getpid())

      // 监听信号
      graceful.ListenSignal()

      // 启动服务器，接受连接
      err := gserver.ListenAndServe(addr)
      if err != nil {
          if err == http.ErrServerClosed {
              logger.Info("server closed")
          } else {
              logger.Error(err.Error())
          }
      }
  }
  ```

* 操作
  - 关闭服务器: `ctrl+c|cmd+c 或者 kill {pid}`
  - 重启服务器: `kill -USR1 {pid} 或者 kill -USR2 {pid}`
  - 重启服务器后，我们发现，仍然可以访问网站。查看端口: `mac` 下 `lsof -i tcp:8877`

# 示例 2

在示例 1 中，使用了默认配置的信号，`SIGUSR1`，`SIGUSR2` 用于重启服务器，`SIGINT`，`SIGTERM` 用于关闭服务器。但在实际情况中，可能 `SIGUSR2` 被用户设置其它用途，比如`重载配置`。这里就需要用到自定义信号

- 代码:

  ```go
  import (
    "net/http"
    "os"

    "github.com/alex-my/ghelper/graceful"
    "github.com/alex-my/ghelper/logger"
  )

    // http.Server 要求实现 ServeHTTP
  type testServer struct {}

  func (server *testServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    // 这里进行逻辑处理，比如按照路由进行处理
    res.Write([]byte("Hello world"))
  }

  func main() {
    addr := "127.0.0.1:8877"
    server := &testServer{}
    logger := logger.NewLogger()

    // 使用 gserver 而不是 http.Server
    gserver := graceful.NewServer(server, logger)

    // 设置重启信号  kill -USR1 {pid}
    graceful.RegisterRestartSignal(syscall.SIGUSR1)

    // 设置关闭信号 ctrl+c 或者 cmd+c
    graceful.RegisterCloseSignal(syscall.SIGINT)

    // 设置退出超时时间，默认 5 秒
    graceful.SetShutdownTimeout(3)

    // 注册清理函数，比如资源回收，先注册的函数先调用
    graceful.RegisterShutdownHandler(func() {
        // close mysql ..
    })
     graceful.RegisterShutdownHandler(func() {
        // some other work
    })

    // 监听信号
    graceful.ListenSignal()

    // 启动服务器，接受连接
    err := gserver.ListenAndServe(addr)
    if err != nil {
        if err == http.ErrServerClosed {
            logger.Info("server closed")
        } else {
            logger.Error(err.Error())
        }
    }
  }
  ```

# API

- **func NewServer(handler http.Handler, logger logger.Log) \*Server**

  - 功能: 生成服务器，用来替代 `http.Server`
  - 参数:
    - **handler**: 与使用`http.Server`一致
    - **logger**: 日志

- **func (server \*Server) ListenAndServe(addr string) error**

  - 功能: 用于替代 `http.Server.ListenAndServe`
  - 参数:
    - **addr**: 监听地址，例如 `:8080`，`localhost:8080`

- **func (server \*Server) ListenAndServeTLS(addr, certFile, keyFile string) error**

  - 功能: 用于替代 `http.Server.ListenAndServeTLS`
  - 参数:
    - **addr**: 监听地址，例如 `:8080`，`localhost:8080`
    - **certFile**: 证书路径
    - **keyFile**: 私钥路径

- **func RegisterRestartSignal(sig ...os.Signal) error**

  - 功能: 设置重启服务器信号，不设置时默认为 `syscall.SIGUSR1, syscall.SIGUSR2`

- **func RegisterCloseSignal(sig ...os.Signal) error**

  - 功能: 设置关闭服务器信号，不设置时默认为 `syscall.SIGINT, syscall.SIGTERM`

- **func SetShutdownTimeout(timeout int)**

  - 功能: 设置超时时间，默认为 `5` 秒 服务器会每隔 500 毫秒检查一次所有连接是否都处理完毕，直到超时时间触发
  - 参数:
    - **timeout**: 超时时间，单位秒

- **func RegisterShutdownHandler(f func()) error**

  - 功能: 设置清理函数，服务器在关闭监听之后，关闭连接之前会调用这些函数，按照注册的顺序调用，先进先出

- **func ListenSignal()**

  - 功能: 监听关闭服务器和重启服务器的信号

- **func Close()**

  - 功能: 直接关闭服务器

- **func Shutdown()**

  - 功能: 优雅的关闭服务器

- **func Restart()**

  - 功能: 优雅的重启服务器

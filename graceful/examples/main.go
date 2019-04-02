package main

import (
	"net/http"
	"os"

	"github.com/alex-my/ghelper/graceful"
	"github.com/alex-my/ghelper/logger"
)

type testServer struct {
}

func (server *testServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello world 2"))
}

func main() {
	server := &testServer{}
	addr := "127.0.0.1:8877"
	logger := logger.NewLogger()

	gserver, err := graceful.NewServer(server, addr, logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Infof("listen on: http://%s, pid: %d", addr, os.Getpid())

	graceful.ListenSignal()

	// 启动，接受连接
	err = gserver.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			logger.Info("server closed")
		} else {
			logger.Error(err.Error())
		}
	}
}

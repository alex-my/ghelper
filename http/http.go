package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var (
	// defaultDisableKeepAlives
	// 默认为 true，当短时间大量连接，可通过 SetDisableKeepAlives 设置为 false
	defaultDisableKeepAlives = true

	// defaultMaxIdleConnsPerHost
	// 默认为 DefaultMaxIdleConnsPerHost，值为 2
	// 当对一个地址频繁连接的时候，可以通过 SetMaxIdleConnsPerHost 加大次数
	defaultMaxIdleConnsPerHost = 0
)

// SetDisableKeepAlives 设置 defaultDisableKeepAlives
func SetDisableKeepAlives(enable bool) {
	defaultDisableKeepAlives = enable
}

// SetMaxIdleConnsPerHost 设置 defaultMaxIdleConnsPerHost
func SetMaxIdleConnsPerHost(count int) {
	defaultMaxIdleConnsPerHost = count
}

// Get GET 请求，包含超时时间
// params 添加到 url 上
// timeout 超时时间，秒
func Get(url string, headers, params map[string]string, timeout int64) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	if params != nil {
		query := req.URL.Query()
		for k, v := range params {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}

	c := client(timeout)
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetRetry GET 请求，包含超时时间，重试次数
// timeout 超时时间，秒
// retryCount 重试次数
// sleep 重试之间 sleep 时间，毫秒
func GetRetry(url string, headers, params map[string]string, timeout int64, retryCount, sleep uint) ([]byte, error) {
	var i uint
	for i = 0; i < retryCount; i++ {
		body, err := Get(url, headers, params, timeout)
		if err != nil {
			time.Sleep(time.Duration(sleep) * time.Millisecond)
			continue
		}
		return body, nil
	}

	return nil, errors.New("timeout with retry")
}

// GetString GET 请求，包含超时时间，结果转为字符串
// params 添加到 url 上
// timeout 超时时间，秒
func GetString(url string, headers, params map[string]string, timeout int64) (string, error) {
	result, err := Get(url, headers, params, timeout)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// GetJSON GET 请求，包含超时时间，结果转为 JSON
// params 添加到 url 上
// timeout 超时时间，秒
// out 结果写入于此，可以是 nil，指针
func GetJSON(url string, headers, params map[string]string, timeout int64, out interface{}) error {
	result, err := Get(url, headers, params, timeout)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(result, out); err != nil {
		return err
	}
	return nil
}

// Post POST 请求(JSON 格式)，包含超时时间
// params 添加到 url 上
// body 添加到 body 上
// timeout 超时时间，秒
func Post(url string, headers, params, body map[string]string, timeout int64) ([]byte, error) {
	var jsonData []byte
	if body != nil {
		var err error
		jsonData, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	if params != nil {
		query := req.URL.Query()
		for k, v := range params {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}

	req.Header.Set("Content-type", "application/json;charset=utf-8")

	c := client(timeout)
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res != nil {
		defer res.Body.Close()
	}
	_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return _body, nil
}

// PostRetry POST 请求，包含超时时间，重试次数
// timeout 超时时间，秒
func PostRetry(url string, headers, params, body map[string]string, timeout int64, retryCount, sleep uint) ([]byte, error) {
	var i uint
	for i = 0; i < retryCount; i++ {
		body, err := Post(url, headers, params, body, timeout)
		if err != nil {
			time.Sleep(time.Duration(sleep) * time.Millisecond)
			continue
		}
		return body, nil
	}

	return nil, errors.New("timeout with retry")
}

// PostString POST 请求，包含超时时间，结果转为字符串
// params 添加到 url 上
// timeout 超时时间，秒
func PostString(url string, headers, params map[string]string, timeout int64) (string, error) {
	result, err := Get(url, headers, params, timeout)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// PostJSON POST 请求，包含超时时间，结果转为 JSON
// params 添加到 url 上
// timeout 超时时间，秒
// out 结果写入于此，可以是 nil，指针
func PostJSON(url string, headers, params map[string]string, timeout int64, out interface{}) error {
	result, err := Get(url, headers, params, timeout)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(result, out); err != nil {
		return err
	}
	return nil
}

// client 获取一个 client
// timeout 超时时间，秒
func client(timeout int64) *http.Client {
	transport := &http.Transport{
		DisableKeepAlives:   defaultDisableKeepAlives,
		MaxIdleConnsPerHost: defaultMaxIdleConnsPerHost,
	}

	if timeout > 0 {
		to := time.Second * time.Duration(timeout)
		// 建立 TCP 连接的时间
		transport.Dial = func(network, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(network, addr, to)
			if err != nil {
				return nil, err
			}
			conn.SetDeadline(time.Now().Add(to))
			return conn, nil
		}
		// 限制读取 response headers 的时间，不包括 body
		transport.ResponseHeaderTimeout = to
	}

	client := &http.Client{
		Transport: transport,
	}
	return client
}

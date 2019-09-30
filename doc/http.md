# 说明

封装 HTTP 请求

# API

- **func Get(url string, headers, params map[string]string, timeout int64) ([]byte, error)**

  - 功能: `GET`请求，包含超时时间
  - 参数:
    - **url**: 请求地址
    - **headers**: 信息头，没有请设置为 `nil`
    - **params**: 参数，没有请设置为 `nil`
    - **timeout**: 超时时间，单位`秒`
  - 示例:

    ```go
    url := "https://www.keylala.cn"
    params := map[string]string{"name": "Alex", "id": "1001"}
    Get(url, nil, params, 5)
    ```

- **func GetRetry(url string, headers, params map[string]string, timeout int64, retryCount, sleep uint) ([]byte, error)**

  - 功能: `GET`请求，如果请求失败，则会重试
  - 参数:
    - **url**: 请求地址
    - **headers**: 信息头，没有请设置为 `nil`
    - **params**: 参数，没有请设置为 `nil`
    - **timeout**: 超时时间，单位`秒`
    - **retryCount**: 重试次数
    - **sleep**: 每次重试之间的时间间隔，单位：毫秒
  - 示例:

    ```go
    url := "https://www.keylala.cn"
    // 每次请求超时时间为 5 秒
    // 如果请求失败，则重试，直到 3 次为止
    // 每次重试之间 Sleep 500 毫秒
    GetRetry(url, nil, nil, 5, 3, 500)
    ```

- **func GetString(url string, headers, params map[string]string, timeout int64) (string, error)**

  - 功能: `GET`请求，并将结果转为字符串
  - 参数:
    - **url**: 请求地址
    - **headers**: 信息头，没有请设置为 `nil`
    - **params**: 参数，没有请设置为 `nil`
    - **timeout**: 超时时间，单位`秒`
  - 示例:

    ```go
    url := "https://www.keylala.cn"
    GetString(url, nil, nil, 5)
    ```

- **func GetJSON(url string, headers, params map[string]string, timeout int64, out interface{}) error**
  - 功能: `GET`请求，并将结果转为 `JSON`
  - 参数:
    - **url**: 请求地址
    - **headers**: 信息头，没有请设置为 `nil`
    - **params**: 参数，没有请设置为 `nil`
    - **timeout**: 超时时间，单位`秒`
    - **out**: 转换结果写入于此

* **func Post(url string, headers, params map[string]string, timeout int64) ([]byte, error)**

  - 功能: `POST`请求，包含超时时间
  - 参数:
    - **url**: 请求地址
    - **headers**: 信息头，没有请设置为 `nil`
    - **params**: 参数，没有请设置为 `nil`
    - **timeout**: 超时时间，单位`秒`
  - 示例:

    ```go
    url := "https://www.keylala.cn"
    params := map[string]string{"name": "Alex", "id": "1001"}
    Post(url, nil, params, 5)
    ```

* **func PostRetry(url string, headers, params map[string]string, timeout int64, retryCount, sleep uint) ([]byte, error)**

  - 功能: `POST`请求，如果请求失败，则会重试
  - 参数:
    - **url**: 请求地址
    - **headers**: 信息头，没有请设置为 `nil`
    - **params**: 参数，没有请设置为 `nil`
    - **timeout**: 超时时间，单位`秒`
    - **retryCount**: 重试次数
    - **sleep**: 每次重试之间的时间间隔，单位：毫秒
  - 示例:

    ```go
    url := "https://www.keylala.cn"
    // 每次请求超时时间为 5 秒
    // 如果请求失败，则重试，直到 3 次为止
    // 每次重试之间 Sleep 500 毫秒
    PostRetry(url, nil, nil, 5, 3, 500)
    ```

* **func PostString(url string, headers, params map[string]string, timeout int64) (string, error)**

  - 功能: `POST`请求，并将结果转为字符串
  - 参数:
    - **url**: 请求地址
    - **headers**: 信息头，没有请设置为 `nil`
    - **params**: 参数，没有请设置为 `nil`
    - **timeout**: 超时时间，单位`秒`
  - 示例:

    ```go
    url := "https://www.keylala.cn"
    PostString(url, nil, nil, 5)
    ```

* **func PostJSON(url string, headers, params map[string]string, timeout int64, out interface{}) error**
  - 功能: `POST`请求，并将结果转为 `JSON`
  - 参数:
    - **url**: 请求地址
    - **headers**: 信息头，没有请设置为 `nil`
    - **params**: 参数，没有请设置为 `nil`
    - **timeout**: 超时时间，单位`秒`
    - **out**: 转换结果写入于此

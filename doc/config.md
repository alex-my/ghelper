# 说明

加载配置文件，支持`JSON`，`TOML`，`YAML`

# 示例

- 配置文件，`JSON`格式

  ```json
  // test.json

  {
    "framework": "gweb",
    "version": "1.0.0"
  }
  ```

* 读取配置

  ```go
  import (
      "github.com/alex-my/ghelper/config"
  )

  c := config.NewConfig()
  err := c.FileJSON("./test.json")
  if err != nil {
      return
  }

  // 使用配置的实例来获取配置
  framework, _ := c.Any("framework")

  // 如果只有一个配置实例，推荐使用 config.C, config.D 等辅助函数
  version, _ := config.C("framework")

  // 如果配置不存在，使用默认值
  addr := config.D("addr", "127.0.0.1:8080")

  ```

# API

- **func NewConfig() Config**

  - 功能: 生成一个配置文件实例，同时也会在赋予给全局的`defaultConfig`，这样可以直接使用`config.C`，`config.D`等函数获取数据

- **func (c \*config) LoadJSON(bytes []byte) error**

  - 功能: 从 `bytes` 中加载 `JSON` 数据

- **func (c \*config) LoadTOML(bytes []byte) error**

  - 功能: 从 `bytes` 中加载 `TOML` 数据

- **func (c \*config) LoadYAML(bytes []byte) error**

  - 功能: 从 `bytes` 中加载 `YAML` 数据

- **func (c \*config) FileJSON(path string) error**

  - 功能: 读取 `JSON` 文件中的数据

- **func (c \*config) FileTOML(path string) error**

  - 功能: 读取 `TOML` 文件中的数据

- **func (c \*config) FileYAML(path string) error**

  - 功能: 读取 `YAML` 文件中的数据

- **func (c \*config) Any(key string) (interface{}, error)**

  - 功能: 从配置中获取键 `key` 对应的配置数据

- **func Any(key string) (interface{}, error)**

  - 功能: 从全局配置 `defaultConfig` 获取键`key`对应的数据

- **func C(key string) (string, error)**

  - 功能: 从全局配置 `defaultConfig` 获取键`key`对应的数据，结果为 `string` 类型

- **func CB(key string) (bool, error)**

  - 功能: 从全局配置 `defaultConfig` 获取键`key`对应的数据，结果为 `bool` 类型

- **func CI(key string) (int, error)**

  - 功能: 从全局配置 `defaultConfig` 获取键`key`对应的数据，结果为 `int` 类型

- **func CI32(key string) (int32, error)**

  - 功能: 从全局配置 `defaultConfig` 获取键`key`对应的数据，结果为 `int32` 类型

- **func CI64(key string) (int64, error)**

  - 功能: 从全局配置 `defaultConfig` 获取键`key`对应的数据，结果为 `int64` 类型

- **func CF32(key string) (float32, error)**

  - 功能: 从全局配置 `defaultConfig` 获取键`key`对应的数据，结果为 `float32` 类型

- **func CF64(key string) (float64, error)**

  - 功能: 从全局配置 `defaultConfig` 获取键`key`对应的数据，结果为 `float64` 类型

- **func D(key, def string) string**

  - 功能: 从全局配置 `defaultConfig` 获取键`key`对应的数据，结果为 `string` 类型，如果不存在，则返回 `def`

- **func DB(key string, def bool) bool**

  - 功能: 从全局配置 `defaultConfig` 获取键`key`对应的数据，结果为 `bool` 类型，如果不存在，则返回 `def`

- **func DI(key string, def int) int**

  - 功能: 从全局配置 `defaultConfig` 获取键`key`对应的数据，结果为 `int` 类型，如果不存在，则返回 `def`

- **func DI32(key string, def int32) int32**

  - 功能: 从全局配置 `defaultConfig` 获取键`key`对应的数据，结果为 `int32` 类型，如果不存在，则返回 `def`

- **func DI64(key string, def int64) int64**

  - 功能: 从全局配置 `defaultConfig` 获取键`key`对应的数据，结果为 `int64` 类型，如果不存在，则返回 `def`

- **func DF32(key string, def float32) float32**

  - 功能: 从全局配置 `defaultConfig` 获取键`key`对应的数据，结果为 `float32` 类型，如果不存在，则返回 `def`

- **func DF64(key string, def float64) float64**
  - 功能: 从全局配置 `defaultConfig` 获取键`key`对应的数据，结果为 `float64` 类型，如果不存在，则返回 `def`

# 说明

封装随机值，唯一值

# API

- **func NewUUID() string**

  - 功能: 产生唯一标识符，12byte。基于时间戳，主机信息，进程 ID 生成，每次调用值加 1，具有原子性
  - 示例:

    ```go
    NewUUID()   // => 5cb840f90a5dcd71e779ba64
    NewUUID()   // => 5cb840f90a5dcd71e779ba65
    ```

- **func String(length int) string**

  - 功能: 获取指定长度的随机字符串

- **func Int(min, max int64) int64**

  - 功能: 获取指定范围内的整数

- **func Uint32() uint32**
  - 功能: 获取随机数，类型为 uint32

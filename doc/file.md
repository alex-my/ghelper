# 说明

文件(夹)相关

# API

- **func IsDir(path string) bool**

  - 功能: 判断路径是否是文件夹
  - 示例:

    ```go
    // /tmp 为实际存在的文件夹
    IsDir("/tmp") // => true

    // /tmp2 为实际不存在的文件夹
    IsDir("/tmp2") // => false

    // /tmp/test.txt 为实际存在的文件
    IsDir("/tmp/test.txt") // => false
    ```

- **func IsFile(path string) bool**

  - 功能: 判断路径是否是文件
  - 示例:

    ```go
    // /tmp 为实际存在的文件夹
    IsFile("/tmp") // => false

    // /tmp/test.txt 为实际存在的文件
    IsFile("/tmp/test.txt") // => true

    // /tmp/test2.txt 为实际不存在的文件
    IsFile("/tmp/test2.txt") // => false

    // /tmp/test_link.txt 为有效的链接文件 ln -s src dst
    IsFile("/tmp/test_link.txt") // => true
    ```

- **func IsExist(path string) bool**

  - 功能: 判断路径是否存在
  - 参数:
    - **path**: 文件/文件夹路径
  - 示例:

    ```go
    // /tmp 为实际存在的文件夹
    IsExist("/tmp") // => true

    // /tmp/test.txt 为实际存在的文件
    IsExist("/tmp/test.txt") // => true
    ```

- **func Name(path string) string**

  - 功能: 获取文件名称，包含后缀。不会判断文件是否存在
  - 示例:

    ```go
    Name("/tmp/test.txt") // => test.txt
    ```

- **func BaseName(path string) string**

  - 功能: 获取文件名称，不带后缀。不会判断文件是否存在
  - 示例:

    ```go
    BaseName("/tmp/test.txt") // => test
    ```

- **func ExtensionName(path string) string**

  - 功能: 获取文件后缀。不会判断文件是否存在
  - 示例:

    ```go
    ExtensionName("/tmp/test.txt") // => .txt
    ```

- **func NameRand(path string, sep ...string) string**

  - 功能: 生成一个随机文件名称
  - 参数:
    - **sep**: 拼接符号，默认为 `"_"`
  - 示例:

    ```go
    NameRand("test.txt") // => test_869mUEfWXOaB.txt
    NameRand("test.txt", "@") // => test@jQ0EQkDQ285x.txt
    ```

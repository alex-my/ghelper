# 说明

对数组进行处理

# API

- **func IntToString(i []int, sep ...string) string**

  - 功能: 将 int 数组转为字符串
  - 参数:
    - **sep**: 拼接符号，默认为 `","`
  - 示例:

    ```go
    IntToString([]int{1, 2, 3, 4})      // => "1,2,3,4"

    IntToString([]int{1, 2, 3, 4}, "+") // => "1+2+3+4"
    ```

- **func Int64ToString(i64 []int64, sep ...string) string**

  - 功能: 将 int64 数组转为字符串
  - 参数:
    - **sep**: 拼接符号，默认为 `","`
  - 示例:

    ```go
    Int64ToString([]int64{1, 2, 3, 4})      // => "1,2,3,4"

    Int64ToString([]int64{1, 2, 3, 4}, "+") // => "1+2+3+4"
    ```

- **func StringToString(str []string, sep ...string) string**

  - 功能: 将 string 数组转为字符串
  - 参数:
    - **sep**: 拼接符号，默认为 `","`
  - 示例:

    ```go
    StringToString([]string{"1", "2", "3", "4"})      // => "1,2,3,4"

    StringToString([]string{"1", "2", "3", "4"}, "+") // => "1+2+3+4"
    ```

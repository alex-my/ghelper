# 说明

- 封装 身份证验证/生成 请求
- 以下使用的身份证号码都为虚拟号码

# API

- **IDCard**

  - 功能: 身份证信息
  - 结构:

    ```go
    type IDCard struct {
        // Code 身份证号码
        Code string
        // Province 省份
        Province string
        // City 地级市
        City string
        // County 县
        County string
        // Year 出生年份
        Year int
        // Month 出生月份
        Month int
        // Day 出生天
        Day int
        // Sex 性别 0 女；1 男
        Sex int
        // SexName 性别名称 Male Female
        SexName string
    }
    ```

- **func IDAreas() map[string]string**

  - 功能: 获取全国 县以上行政区划代码
  - 示例:

    ```go
    ares := IDAreas()
    // ares
    // {
    // 	"110000": "北京市",
    // 	"110101": "东城区",
    // 	"110102": "西城区",
    // 	...
    // 	"440000": "广东省",
    // 	"440100": "广州市",
    // 	"440103": "荔湾区",
    // 	"440104": "越秀区",
    // 	"440105": "海珠区",
    // 	"440106": "天河区",
    // 	"440111": "白云区",
    // 	"440112": "黄埔区",
    // 	"440113": "番禺区",
    // 	"440114": "花都区",
    // 	"440115": "南沙区",
    // 	"440117": "从化区",
    // 	"440118": "增城区",
    // 	"440200": "韶关市",
    // 	"440203": "武江区",
    // 	...
    // }

    ```

- **func IDCheck(code string) bool**

  - 功能: 验证身份证是否正确
  - 参数:
    - **code**: 身份证号码
  - 返回值:
    - **true**: 身份证正确
    - **false**: 身份证错误

- **func IDInfo(code string) (\*IDCard, error)**

  - 功能: 解析出身份信息
  - 参数:
    - **code**: 身份证号码
  - 示例:

    ```go
    fmt.Printf("%+v\n", IDInfo("440106199910017896"))
    // => {Code:440106199910017896 Province:广东省 City:广州市 County:天河区 Year:1999 Month:10 Day:1 Sex:1 SexName:Male}
    ```

- **func IDGenerate(year, month, day, sex int, areaCode string, count int) ([]string, error)**

  - 功能: 生成虚拟的身份证信息
  - 参数:
    - **year**: 年份，例如 2019
    - **month**: 月份，1 ~ 12
    - **day**: 天数: 1-31
    - **sex**: 性别，0 女，1 男
    - **areaCode**: 区域编码，所有区域可以通过 `IDAreas()` 获取
    - **count int**: 生成身份证个数
  - 示例:

    ```go
    codes, err := IDGenerate(1999, 10, 1, 1, "440106", 5)
    // => [440106199910016594 440106199910012155 440106199910013238 440106199910013959 440106199910019074]
    ```

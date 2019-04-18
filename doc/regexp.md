# 说明

封装常用的正则表达式

# API

- **func IsChinesePhone(phone string) bool**

  - 功能: 判断是否是中国大陆手机号

- **func IsNickName(name string) bool**

  - 功能: 判断是否是合法的昵称，包含: 数字，大小写英文字母，\_, -, 汉字

- **func IsAccount(account string) bool**
  - 功能: 判断是否是合法的账号/用户名，包含：数字，大小写字母，_，—，其中，数字，_，- 不可以放在开头

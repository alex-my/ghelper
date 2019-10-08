## 阿里云邮箱

```go
import "github.com/alex-my/ghelper/email"

email.SetConfig(&email.Config{
	Host:     "smtp.mxhichina.com",
	Port:     465,
	Username: "阿里云邮箱地址",
	Password: "阿里云邮箱密码",
})

// 允许同时发送给多个目标
err := email.Send("xxx1@qq.com,xxx2@126.com", "这是标题", "<p>这是内容</p>")
```

## QQ 邮箱

```go
import "github.com/alex-my/ghelper/email"

email.SetConfig(&email.Config{
	Host:     "smtp.qq.com",
	Port:     465,
	Username: "QQ邮箱地址",
	Password: "QQ邮箱授权码",
})

err := email.Send("xxx@example.com", "这是标题", "<p>这是内容</p>")
```

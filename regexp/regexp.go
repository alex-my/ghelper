package regexp

import (
	"regexp"
)

const (
	// patternChinesePhone 中国大陆手机号，以1开头，11位
	patternChinesePhone = `^1[34578]{1}\d{9}$`

	// patternNickName 判断是否是合法的昵称，包含: 数字，大小写英文字母，_， -， 汉字
	patternNickName = `^[-_a-z0-9A-Z\p{Han}]+$`

	// patternAccount 判断是否是合法的用户名，包含：数字，大小写字母，_，—，其中，数字，_，- 不可以放在开头
	patternAccount = `^[a-zA-Z]+([-_a-z0-9A-Z]+)*?$`

	// patternEmail 邮件
	patternEmail = "^[a-zA-Z0-9_!#$%&'*+/=?`{|}~^.-]+@[a-zA-Z0-9.-]+$"
)

var (
	regexpChinesePhone = regexp.MustCompile(patternChinesePhone)
	regexpNickName     = regexp.MustCompile(patternNickName)
	regexpAccount      = regexp.MustCompile(patternAccount)
	regexpEmail        = regexp.MustCompile(patternEmail)
)

// IsChinesePhone 判断是否是中国大陆手机号
func IsChinesePhone(phone string) bool {
	if len(phone) != 11 {
		return false
	}

	return regexpChinesePhone.MatchString(phone)
}

// IsNickName 判断是否是合法的昵称
// 包含: 数字，大小写英文字母，_, -, 汉字
func IsNickName(name string) bool {
	if name == "" {
		return false
	}

	return regexpNickName.MatchString(name)
}

// IsAccount 判断是否是合法的账号/用户名
// 包含：数字，大小写字母，_，—，其中，数字，_，- 不可以放在开头
func IsAccount(account string) bool {
	if account == "" {
		return false
	}

	return regexpAccount.MatchString(account)
}

// IsEmail 判断是否是邮件
func IsEmail(email string) bool {
	if email == "" {
		return false
	}

	return regexpEmail.MatchString(email)
}

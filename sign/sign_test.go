package sign_test

import (
	"testing"

	"github.com/alex-my/ghelper/sign"
)

func TestValueWithSign(t *testing.T) {
	params := map[string]string{
		"timestamp": "1571747084",
		"username":  "root",
		"password":  "123456",
	}

	c1, _ := sign.Sign(params)

	params["sign"] = "2f9d60bd2032c8bac547c064a95363f2"

	c2, _ := sign.Sign(params)

	if c1 != c2 || c2 != params["sign"] {
		t.Error("all the same")
	}
}

func TestValueWithCalc(t *testing.T) {
	// 自定义签名函数
	calc := func(s string, k string) string {
		return s
	}

	sign.SetCalcFunc(calc)

	params := map[string]string{
		"timestamp": "1571747084",
		"username":  "root",
		"password":  "123456",
	}

	c, _ := sign.Sign(params)
	if c != "password=123456&timestamp=1571747084&username=root" {
		t.Error("sign failed")
	}
}

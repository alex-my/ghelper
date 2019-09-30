package human

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	_time "time"

	"github.com/alex-my/ghelper/random"
	"github.com/alex-my/ghelper/time"
)

// Package 中国身份证验证与生成

var (
	pattern = "^(\\d{6})(18|19|20)?(\\d{2})(0\\d|10|11|12)([012]\\d|3[01])(\\d{3})(\\d|X)?$"
	reg     = regexp.MustCompile(pattern)

	weight    = []int64{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	remainder = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
)

// IDAreas 获取全国 县以上行政区划代码
// 定期更新数据，具体见本目录下的 idcard.go
// eg:
// IDAreas() ->
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
// 	...
// }
func IDAreas() map[string]string {
	return areas
}

// IDCheck 验证身份证是否正确
// code: 身份证号码
// return: true 正确；false 错误
func IDCheck(code string) bool {
	idCard := &IDCard{Code: code}
	return idCard.check()
}

// IDInfo 解析出身份信息
// code: 身份证号码
// return: 身份证信息
// eg:
// (数据为虚构，如有巧合，实属意外)
// IDInfo("440106199910017896")
// -> {Code:440106199910017896 Province:广东省 City:广州市 County:天河区 Year:1999 Month:10 Day:1 Sex:1 SexName:Male}
func IDInfo(code string) (*IDCard, error) {
	idCard := &IDCard{Code: code}
	if err := idCard.parse(); err != nil {
		return nil, err
	}
	return idCard, nil
}

// IDGenerate 生成身份证信息
// year, month, day 出生年月日
// sex 性别，0 女，1 男
// areaCode 区域编码，所有区域可以通过 IDAreas() 获取
// count 生成身份证个数
// return 身份证信息
// eg:
// codes, err := IDGenerate(1999, 10, 1, 1, "440106", 5)
// -> [440106199910016594 440106199910012155 440106199910013238 440106199910013959 440106199910019074]
func IDGenerate(year, month, day, sex int, areaCode string, count int) ([]string, error) {
	// 验证地域信息是否存在
	if _, exist := areas[areaCode]; !exist {
		return nil, fmt.Errorf("Area code not found: %s", areaCode)
	}
	// 验证性别是否正确
	if sex != 0 && sex != 1 {
		return nil, errors.New("Sex error, must be 0 (female) or 1 (male)")
	}
	// 验证日期是否合法
	// 生成身份证信息不判断是否尚未出生，或者岁数过大的问题
	birth := fmt.Sprintf("%04d%02d%02d", year, month, day)
	_, err := _time.Parse("20060102150405", fmt.Sprintf("%s235959", birth))
	if err != nil {
		return nil, err
	}

	l := make([]string, 0, count)

	buf := bytes.Buffer{}

	for i := 0; i < count; i++ {
		// 三位数随机码，奇数分配给男性，偶数分配给女性
		randomCode := random.Int(100, 998)
		if sex == 0 && randomCode%2 != 0 {
			randomCode++
		} else if sex == 1 && randomCode%2 == 0 {
			randomCode++
		}
		// 生成校验码
		code17 := fmt.Sprintf("%s%s%d", areaCode, birth, randomCode)
		checkCode := calcCheckCode(fmt.Sprintf("%s0", code17))

		buf.Write([]byte(code17))
		buf.Write([]byte{checkCode})
		l = append(l, buf.String())
		buf.Reset()
	}

	return l, nil
}

// IDCard 身份证信息
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

func (idCard *IDCard) parse() error {
	code := idCard.Code
	if !reg.MatchString(code) {
		return errors.New("ID card format is incorrect")
	}
	// 前六位为地区编码
	areaCode := code[:6]
	if _, exist := areas[areaCode]; !exist {
		return fmt.Errorf("Area code not found: %s", areaCode)
	}
	if err := idCard.parseAreaName(areaCode); err != nil {
		return err
	}
	// 分析出生日期
	birth := code[6:14]
	_, err := _time.Parse("20060102", birth)
	if err != nil {
		return err
	}
	now := time.Date(time.YMD3)
	ibirth, _ := strconv.ParseInt(birth, 10, 32)
	inow, _ := strconv.ParseInt(now, 10, 32)
	// 不可以超过今天
	if ibirth > inow {
		return errors.New("Not yet born")
	}
	// 岁数太大，不正常
	if (inow-ibirth)/10000 > 150 {
		return errors.New("There is no such long-lived person")
	}
	idCard.Year, _ = strconv.Atoi(code[6:10])
	idCard.Month, _ = strconv.Atoi(code[10:12])
	idCard.Day, _ = strconv.Atoi(code[12:14])
	// 性别
	sex, _ := strconv.Atoi(code[14:17])
	if sex%2 == 0 {
		idCard.Sex = 0
		idCard.SexName = "Female"
	} else {
		idCard.Sex = 1
		idCard.SexName = "Male"
	}
	// 校验码检测
	if calcCheckCode(code) != code[len(code)-1] {
		return errors.New("Check code detection failed")
	}

	return nil
}

// 获取 省，市，县 名称
func (idCard *IDCard) parseAreaName(areaCode string) error {
	province, city := areaCode[:2], areaCode[2:4]
	provinceName, exist := areas[fmt.Sprintf("%s0000", province)]
	if !exist {
		return errors.New("Province code error")
	}
	idCard.Province = provinceName

	// 部分地区 City 这一级没有编码，比如北京，天津
	cityName, exist := areas[fmt.Sprintf("%s%s00", province, city)]
	if exist {
		idCard.City = cityName
	} else {
		idCard.City = provinceName
	}

	countyName, exist := areas[areaCode]
	if !exist {
		return errors.New("County code error")
	}
	idCard.County = countyName

	return nil
}

func (idCard *IDCard) check() bool {
	return idCard.parse() == nil
}

// calcCheckCode 计算校验码
// code 18位身份证号码
// return 校验码
func calcCheckCode(code string) byte {
	// 将前17位分别乘以不同的系数(weight)，并将结果相加
	// 将和求余 11，将结果取 code 中的数字，例如余数为 2，则校验码为 'X'
	var sum int64
	for i, char := range code[:len(code)-1] {
		ichar, _ := strconv.ParseInt(string(char), 10, 64)
		sum += ichar * weight[i]
	}
	return remainder[int(sum)%len(remainder)]
}

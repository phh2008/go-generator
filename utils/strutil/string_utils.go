package strutil

import (
	"regexp"
	"strings"
)

var regUnderLineLetter = regexp.MustCompile("_(\\w)")

var regFirstLowerCase = regexp.MustCompile("^[a-z]")

var regFirstUpperCase = regexp.MustCompile("^[A-Z]")

//下划线转驼峰
func UnderLineToCamelcase(src string) string {
	return regUnderLineLetter.ReplaceAllStringFunc(src, func(s string) string {
		return strings.ToUpper(string(s[1]))
	})
}

//首字母大写
func FirstLetterToUpper(src string) string {
	if src == "" {
		return src
	}
	return regFirstLowerCase.ReplaceAllString(src, strings.ToUpper(string(src[0])))
}

//首字母小写
func FirstLetterToLower(src string) string {
	if src == "" {
		return src
	}
	return regFirstUpperCase.ReplaceAllString(src, strings.ToLower(string(src[0])))
}

//匹配子串的位置
func IndexRune(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}
	return result
}

//最后一个匹配子串的位置
func LastIndexRune(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.LastIndex(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}
	return result
}

func SubStr1(str string, begin int) (substr string) {
	rs := []rune(str)
	lth := len(rs)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	return string(rs[begin:])
}

func SubStr2(str string, begin, length int) (substr string) {
	rs := []rune(str)
	lth := len(rs)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	return string(rs[begin:end])
}

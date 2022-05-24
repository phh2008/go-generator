package strutil

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func Test_regex_(t *testing.T) {
	reg := regexp.MustCompile("_(\\w)")

	text := "string_utils_test"

	list := reg.FindAllString(text, -1)

	for _, v := range list {
		fmt.Println(v)
	}
	fmt.Println("------------------------------")

	list2 := reg.FindAllStringSubmatch(text, -1)
	fmt.Println(list2)
	for _, v := range list2[0] {
		fmt.Println(v)
	}
	fmt.Println("------------------------------")
	list3 := reg.FindSubmatch([]byte(text))

	for _, v := range list3 {
		fmt.Println(string(v))
	}
}

//下划线转驼峰
func Test_underLine2Camelcase(t *testing.T) {
	reg := regexp.MustCompile("_(\\w)")

	text := "mem_string_utils_test_under"
	fmt.Println(text)

	text = reg.ReplaceAllStringFunc(text, func(s string) string {
		return strings.ToUpper(string(s[1]))
	})

	fmt.Println(text)
}

//首字母大写
func Test_firstLetterToUpper(t *testing.T) {
	reg := regexp.MustCompile("^[a-z]")
	text := "mem_string_utils_test_under"
	text = reg.ReplaceAllString(text, strings.ToUpper(string(text[0])))
	fmt.Println(text)
}

func TestFirstLetterToLower(t *testing.T) {
	text := "AMem_string_utils_test_under"
	fmt.Println(text)
	text = FirstLetterToLower(text)
	fmt.Println(text)
}

func TestFirstLetterToUpper(t *testing.T) {
	text := "mem_string_utils_test_under"
	fmt.Println(text)
	text = FirstLetterToUpper(text)
	fmt.Println(text)
}

func TestUnderLineToCamelcase(t *testing.T) {
	text := "mem_string_utils_test_under"
	fmt.Println(text)
	text = UnderLineToCamelcase(text)
	fmt.Println(text)
}

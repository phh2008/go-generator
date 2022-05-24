package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGetFileNameList(t *testing.T) {
	list, _ := GetFileNameList("../../resource/tpl")
	fmt.Println(list)
}

//创建目录
func Test_createDir(t *testing.T) {

	err := os.MkdirAll("./com/phh/dao", os.ModeDir)
	fmt.Println(err)
}

//创建文件
func Test_createFile(t *testing.T) {
	f, err := os.OpenFile("./com/phh/dao/xx.text", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.WriteString("新内容")

}

//压缩文件
func Test_compress_zip(t *testing.T) {
	dir := "./com"
	infoList, err := GetFileInfoList(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	var files []*os.File
	for _, info := range infoList {
		fmt.Println(">>>>>>>:", info.Path)
		file, err := os.Open(info.Path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		files = append(files, file)
	}
	outFile, err := os.Create("./out.zip")
	err = Compress(files, outFile)
	fmt.Println(err)
}

func Test_zip_2(t *testing.T) {
	dir := "./"
	outFile, err := os.Create("./out.zip")
	err = Zip(dir, outFile)
	fmt.Println(err)
}

func Test_xxx(t *testing.T) {
	r := filepath.Join("", "E:/com/cn")
	fmt.Println(filepath.Clean("./com/cn"))
	fmt.Println(r)
	fmt.Println(filepath.Base("./com/cn"))
}

package service

import (
	"com.phh/go-generator/utils/strutil"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"text/template"
)

type Column struct {
	Type string
	Name string
}

//测试模版
func Test_tpl_entity(t *testing.T) {

	dataMap := make(map[string]interface{})
	dataMap["package"] = "com.phh"
	dataMap["author"] = "phh"
	dataMap["comments"] = "测试"
	dataMap["datetime"] = "2019-08-06 15:06:00"
	dataMap["className"] = "User"
	var columns []Column
	columns = append(columns, Column{"String", "userName"})
	columns = append(columns, Column{"Integer", "age"})
	columns = append(columns, Column{"Date", "birthday"})
	fmt.Println(columns)
	dataMap["columns"] = columns

	tpl, err := template.ParseFiles("../resource/tpl/entity.java.html")
	if err != nil {
		panic(err)
	}
	tpl.Execute(os.Stdout, dataMap)
}

func Test_tpl_mapper_java(t *testing.T) {
	dataMap := make(map[string]interface{})
	dataMap["package"] = "com.phh"
	dataMap["author"] = "phh"
	dataMap["comments"] = "测试"
	dataMap["datetime"] = "2019-08-06 15:06:00"
	dataMap["className"] = "User"
	baseMapper := "com.phh.base.mapper.BaseMapper"
	dataMap["baseMapper"] = baseMapper
	fmt.Println(strutil.IndexRune(baseMapper, "."))
	fmt.Println(strutil.LastIndexRune(baseMapper, "."))
	dataMap["baseMapperName"] = strutil.SubStr1(baseMapper, strutil.LastIndexRune(baseMapper, ".")+1)
	tpl, err := template.ParseFiles("../resource/tpl/mapper.java.html")
	if err != nil {
		panic(err)
	}
	tpl.Execute(os.Stdout, dataMap)
}

func Test_getFiles(t *testing.T) {
	var files []string
	filepath.Walk("../resource/tpl/", func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, info.Name())
		}
		return nil
	})
	fmt.Println(files)
}

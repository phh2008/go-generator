package service

import (
	"com.phh/go-generator/dao"
	"com.phh/go-generator/domain"
	"com.phh/go-generator/utils/dateutil"
	"com.phh/go-generator/utils/fileutil"
	"com.phh/go-generator/utils/mysqlutil"
	"com.phh/go-generator/utils/strutil"
	"com.phh/go-generator/vo"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
	"unicode/utf8"
)

type GeneratorService interface {
	// QueryTableList 获取表列表
	QueryTableList(name string) []domain.TableName

	// Generate 生成代码模版文件
	Generate(gen vo.Gen) (filePath string, err error)
}

func NewGeneratorService(dao dao.GeneratorDao) GeneratorService {
	return &generatorService{
		genDao: dao,
	}
}

type generatorService struct {
	genDao dao.GeneratorDao
}

// QueryTableList 获取表列表
func (g *generatorService) QueryTableList(name string) []domain.TableName {
	return g.genDao.QueryTableList(name)
}

// Generate 生成代码模版文件
func (g *generatorService) Generate(gen vo.Gen) (filePath string, err error) {
	setDefaulSuffix(&gen)
	//模版函数
	funcMap := getFuncMap()
	tmplDir := "./resource/tpl/"
	files, err := fileutil.GetFileNameList(tmplDir)
	if err != nil {
		fmt.Println(err)
		log.Error().Msg(TEMPLATE_NOT_FOUND.Error())
		return "", TEMPLATE_NOT_FOUND
	}
	var filePaths []string
	for _, v := range files {
		filePaths = append(filePaths, tmplDir+v)
	}
	tmpl, err := template.New("goTmpl").Funcs(funcMap).ParseFiles(filePaths...)
	if err != nil {
		log.Error().Err(err).Msg(TEMPLATE_LOAD_ERROR.Error())
		return "", TEMPLATE_LOAD_ERROR
	}
	date := dateutil.FormatTime(time.Now(), "yyyy-MM-dd")
	//生成文件根目录
	rootDir := "./go_code_tmpl/"
	_ = os.RemoveAll(rootDir)
	for _, tableName := range gen.Tables {
		//模版参数
		dataMap := map[string]interface{}{}
		dataMap["gen"] = gen
		dataMap["date"] = date
		dataMap["primaryKeyName"] = "id" //主键实体映射统一为：id
		hasServiceInterface := gen.HasServiceInterface == "on"
		dataMap["hasServiceInterface"] = hasServiceInterface
		dataMap["hasTime"] = false
		dataMap["hasDecimal"] = false
		columns := g.genDao.GetTableColumnsByTableName(tableName)
		dataMap["columnNumber"] = len(columns)
		//转换列名与类型
		maxColumnLength := 0
		for i, col := range columns {
			//mysql类型转换为 Go 类型
			columns[i].GoType = mysqlutil.GetGoType(col.DataType)
			//mysql字段名称转换为符合 go 名称
			goName := strutil.UnderLineToCamelcase(col.Name)
			goName = strutil.FirstLetterToUpper(goName)
			columns[i].GoName = goName
			if columns[i].GoType == "time.Time" {
				dataMap["hasTime"] = true
			}
			if columns[i].GoType == "decimal.Decimal" {
				dataMap["hasDecimal"] = true
			}
			//主键
			if col.Key == "PRI" || i == 0 {
				//没有主键就是第一列
				dataMap["primaryKeyGoType"] = columns[i].GoType
				dataMap["primaryKeyColumn"] = col.Name
			}
			// 字段长度
			columnLength := utf8.RuneCountInString(goName)
			if columnLength > maxColumnLength {
				maxColumnLength = columnLength
			}
		}
		table := g.genDao.GetTableByTableName(tableName)
		dataMap["columns"] = columns
		dataMap["table"] = table
		dataMap["maxColumnLength"] = maxColumnLength
		//表名对应 go 名称：下划线转换为驼峰，首字母大写
		goName := tableName
		//是否生成模块名
		hasModule := gen.HasModule == "on"
		//模块名称
		mod := ""
		if hasModule {
			//解析第一个下划线前的词
			lineIdx := strutil.IndexRune(tableName, "_")
			if lineIdx <= 0 {
				hasModule = false
			} else if lineIdx < (len(tableName) - 1) {
				//if 确保下划线不是最后一个字符
				mod = strutil.SubStr2(tableName, 0, lineIdx)
				//截取第一个下划线之后部分为 go 名称
				goName = strutil.SubStr1(tableName, lineIdx+1)
			} else {
				hasModule = false
			}
		}
		// 去模块名的表名，用作生成 go 文件名称
		goNameUnderline := strings.ToLower(goName)
		//是否生成模块名
		dataMap["hasModule"] = hasModule
		//模块名称
		dataMap["mod"] = mod
		//把表名下划线转换为驼峰
		goName = strutil.UnderLineToCamelcase(goName)
		goName = strutil.FirstLetterToUpper(goName)
		dataMap["goName"] = goName
		//生成文件
		modPath := "/" + mod
		if !strings.HasSuffix(modPath, "/") {
			modPath = modPath + "/"
		}
		doDir := rootDir + gen.DoPkg + modPath
		daoDir := rootDir + gen.DaoPkg + modPath
		serviceDir := rootDir + gen.ServicePkg + modPath
		_ = os.MkdirAll(doDir, os.ModeDir|os.ModePerm)
		_ = os.MkdirAll(daoDir, os.ModeDir|os.ModePerm)
		_ = os.MkdirAll(serviceDir, os.ModeDir|os.ModePerm)
		filePath := ""
		for _, v := range files {
			// 注意：go 文件名全为小写
			if strings.Contains(v, "do.go") {
				filePath = doDir + spellGoFileName(goNameUnderline, gen.DoSuffix)
			} else if strings.Contains(v, "dao.go") {
				filePath = daoDir + spellGoFileName(goNameUnderline, gen.DaoSuffix)
			} else if strings.Contains(v, "service.go") {
				filePath = serviceDir + spellGoFileName(goNameUnderline, gen.ServiceSuffix)
			} else {
				//未启用的模版
				continue
			}
			file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
			if err != nil {
				fmt.Println(err)
				return "", FILE_OPEN_ERROR
			}
			err = tmpl.ExecuteTemplate(file, v, dataMap)
			if err != nil {
				fmt.Println(err)
				return "", TEMPLATE_RENDER_ERROR
			}
			file.Close()
		}
	}
	//打包zip
	fileName := "./go-code-tmpl-out.zip"
	zipFile, err := os.Create(fileName)
	fmt.Println(err)
	defer zipFile.Close()
	err = fileutil.Zip(rootDir, zipFile)
	if err != nil {
		fmt.Println(err)
		return "", errors.New(fmt.Sprintf("生成文件成功，但打包zip失败，可在目录(%s)找到文件", rootDir))
	}
	return fileName, nil
}

func spellGoFileName(fileName string, suffix string) string {
	end := suffix
	if suffix != "" {
		end = "_" + strings.ToLower(suffix)
	}
	return fileName + end + ".go"
}

func setDefaulSuffix(gen *vo.Gen) {
	//dao,do,service名称后辍
	if gen.DoSuffix == "" {
		gen.DoSuffix = ""
	}
	if gen.DaoSuffix == "" {
		gen.DaoSuffix = "DAO"
	}
	if gen.ServiceSuffix == "" {
		gen.ServiceSuffix = "Service"
	}
	if gen.MapperXmlSuffix == "" {
		gen.MapperXmlSuffix = "Mapper"
	}
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"FirstToLower": func(str string) string {
			return strutil.FirstLetterToLower(str)
		},
		"FirstToUpper": func(str string) string {
			return strutil.FirstLetterToUpper(str)
		},
		"ClassName": func(fullClass string) string {
			dotIdx := strutil.LastIndexRune(fullClass, ".")
			return strutil.SubStr1(fullClass, dotIdx+1)
		},
		"ToUpper": func(str string) string {
			return strings.ToUpper(str)
		},
		"ToLower": func(str string) string {
			return strings.ToLower(str)
		},
		"Add": func(a int, b int) int {
			return a + b
		},
		"Minus": func(a int, b int) int {
			return a - b
		},
		"In": func(src string, des string, sep string) bool {
			list := strings.Split(des, sep)
			for _, v := range list {
				if v == src {
					return true
				}
			}
			return false
		},
		"NotIn": func(src string, des string, sep string) bool {
			list := strings.Split(des, sep)
			for _, v := range list {
				if v == src {
					return false
				}
			}
			return true
		},
		"pkg": func(str string) string {
			slashIdx := strutil.LastIndexRune(str, "/")
			if slashIdx > 0 {
				return strutil.SubStr1(str, slashIdx+1)
			}
			return str
		},
		"FmtLen": func(str string, len int) string {
			var f = "%-" + strconv.Itoa(len) + "s"
			return fmt.Sprintf(f, str)
		},
	}
}

package controller

import (
	"com.phh/go-generator/domain"
	"com.phh/go-generator/service"
	"com.phh/go-generator/utils/dbutil"
	"com.phh/go-generator/vo"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"io"
	"io/ioutil"
	"os"
)

type IndexController struct {
	Ctx              iris.Context
	GeneratorService service.GeneratorService
}

// Get 首页
// localhost:8080/
func (c *IndexController) Get() mvc.Result {
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "首页",
			"navIndex": "layui-this",
		},
	}
}

// GetAbout 关于页
// localhost:8080/about
func (c *IndexController) GetAbout() mvc.Result {
	return mvc.View{
		Name: "about.html",
		Data: iris.Map{
			"Title":    "关于",
			"navAbout": "layui-this",
		},
	}
}

// GetTables 获取表列表
func (c *IndexController) GetTables() {
	name := c.Ctx.URLParam("name")
	var result = iris.Map{}
	var code, msg string
	var list []domain.TableName
	if dbutil.Db == nil {
		code = "0001"
		msg = "请连接数据库"
	} else {
		code = "0"
		msg = ""
		list = c.GeneratorService.QueryTableList(name)
	}
	result["code"] = code
	result["msg"] = msg
	result["count"] = len(list)
	result["data"] = list
	//c.Ctx.JSON(result)
	ret, _ := json.Marshal(result)
	c.Ctx.Write(ret)
}

// PostConnect 连接数据库
func (c *IndexController) PostConnect() {
	var conn vo.ConnectVO
	c.Ctx.ReadJSON(&conn)
	//加载数据库配置
	//"root:root@/demo?charset=utf8&parseTime=True&loc=Local"
	connStr := "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", fmt.Sprintf(connStr, conn.Username, conn.Password, conn.Url, conn.Db))
	if err != nil {
		fmt.Println(err)
		result := iris.Map{"code": "0001", "msg": "数据库连接失败"}
		//c.Ctx.JSON(result)
		ret, _ := json.Marshal(result)
		c.Ctx.Write(ret)
		return
	}
	db.SingularTable(true)
	db.LogMode(true)
	//连接成功，产闭原有连接
	if dbutil.Db != nil {
		_ = dbutil.Db.Close()
	}
	dbutil.Db = db
	result := iris.Map{"code": "0000", "msg": "success"}
	ret, _ := json.Marshal(result)
	//c.Ctx.JSON(result)
	c.Ctx.Write(ret)
}

// PostGen 生成代码文件
func (c *IndexController) PostGen() {
	var gen vo.Gen
	c.Ctx.ReadJSON(&gen)
	filePath, err := c.GeneratorService.Generate(gen)
	code := "0000"
	msg := ""
	if err != nil {
		code = "0001"
		msg = err.Error()
	}
	result := iris.Map{"code": code, "msg": msg, "data": filePath}
	ret, err := json.Marshal(result)
	//c.Ctx.JSON(result)
	c.Ctx.Write(ret)
}

// GetDownload 下载文件
func (c *IndexController) GetDownload() {
	file := c.Ctx.FormValue("file")
	f, err := os.Open(file)
	if err != nil {
		c.Ctx.WriteString("获取文件错误")
		return
	}
	c.Ctx.StreamWriter(func(w io.Writer) bool {
		bt, _ := ioutil.ReadAll(f)
		w.Write(bt)
		return false
	})
}

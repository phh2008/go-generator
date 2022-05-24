package main

import (
	"com.phh/go-generator/config"
	"com.phh/go-generator/dao"
	"com.phh/go-generator/service"
	"com.phh/go-generator/utils/strutil"
	"com.phh/go-generator/web/controller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"github.com/rs/zerolog/log"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func main() {
	//加载配置
	cfg := config.Cfg()
	//创建iris服务
	app := iris.New()
	app.Logger().SetLevel("debug")
	// 加载模版文件
	tmpl := iris.HTML("./web/templates", ".html").
		Layout("shared/layout.html").
		Reload(true)
	tmpl.AddFunc("fullPath", func(url string) string {
		root := cfg.ContextPath
		if strings.LastIndex(root, "/") == (len(root) - 1) {
			root = strutil.SubStr2(root, 0, len(root)-1)
		}
		return root + url
	})
	app.RegisterView(tmpl)
	//静态资源目录
	app.StaticWeb(cfg.StaticPath, "./web/res")
	app.Use(recover.New())
	app.Use(func(ctx context.Context) {
		log.Info().Str("url", ctx.RequestPath(false)).Msg("")
		ctx.ViewData("ctxPath", cfg.ContextPath)
		ctx.Next()
	})
	//错误处理
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("shared/error.html")
	})

	dao := dao.NewGeneratorDao()
	service := service.NewGeneratorService(dao)

	//root context-path
	//root := app.Party(cfg.ContextPath)
	root := mvc.New(app.Party(cfg.ContextPath))
	{
		//可以为controller 注入需要的对象
		root.Register(service)
		//controller 处理器注册
		root.Handle(new(controller.IndexController))
	}
	//打开浏览器
	url := "http://" + cfg.ServerAddress + cfg.ContextPath
	go openBrowser(url)
	//启动应用
	app.Run(
		// 启动端口
		iris.Addr(cfg.ServerAddress),
		//iris.WithoutVersionChecker,
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// 优化
		iris.WithOptimizations,
	)

}

func openBrowser(url string) {
	time.Sleep(time.Second * 2)
	//当前系统
	goos := runtime.GOOS
	var err error
	if goos == "windows" {
		cmd := exec.Command("cmd", "/c", "start", url)
		//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		err = cmd.Start()
	} else if goos == "darwin" {
		//mac系统
		err = exec.Command("open", url).Start()
	}
	if err != nil {
		log.Warn().Msg("调用浏览器失败，URL:" + url)
	}
}

package service

import "errors"

var (
	//模版文件未找到
	TEMPLATE_NOT_FOUND = errors.New("模版文件未找到")

	//模版加载错误
	TEMPLATE_LOAD_ERROR = errors.New("模版加载错误")

	//模版渲染错误
	TEMPLATE_RENDER_ERROR = errors.New("模版渲染错误")

	//目录创建错误
	DIR_CREATE_ERROR = errors.New("目录创建错误")

	//写入文件错误
	FILE_WRITE_ERROR = errors.New("写入文件错误")

	//打开文件错误
	FILE_OPEN_ERROR = errors.New("打开文件错误")
)

package vo

type Gen struct {
	//作者
	Author string `json:"author"`
	//dao父类
	DaoParentClass string `json:"daoParentClass"`
	//dao包名
	DaoPkg string `json:"daoPkg"`
	//dao名称后辍
	DaoSuffix string `json:"daoSuffix"`
	//do父类
	DoParentClass string `json:"doParentClass"`
	//do包名
	DoPkg string `json:"doPkg"`
	//do名称后辍
	DoSuffix string `json:"doSuffix"`
	//是否生成模块名:on/off
	HasModule string `json:"hasModule"`
	//service父类
	ServiceParentClass string `json:"serviceParentClass"`
	//service包名
	ServicePkg string `json:"servicePkg"`
	//service名称后辍
	ServiceSuffix string `json:"serviceSuffix"`
	//是否生成服务接口
	HasServiceInterface string `json:"hasServiceInterface"`
	//service接口父类
	ServiceInterfaceParentClass string `json:"serviceInterfaceParentClass"`
	//mapper.xml名称辍
	MapperXmlSuffix string `json:"mapperXmlSuffix"`
	//表名
	Tables []string `json:"tables"`
}

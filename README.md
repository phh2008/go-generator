![sample](https://github.com/phh2008/mybatis-generator-golang-impl/blob/master/logo/pic.png)

+ 模版中的变量|函数说明
+ 熟悉golang text/template可自行调整模版，使用已定义的变量
+ 注意区分大小写

+ 模版函数

    |名称| 参数类型 | 说明|  
    | :---- | :---- | :---- |  
    | FirstToLower |  string | 首字母小写 |  
    | FirstToUpper |  string | 首字母大写 |   
    | ClassName |  string | 从全类名提取类名 |   
    | ToUpper |  string | 转大写 |   
    | ToLower |  string | 转小写 |   
    | Add |  int,int | 加法：return a+b |   
    | Minus |  int,int | 减法：return a-b |   
    | In |  src string,des string,sep string | 是否包括字符串，如：src=state,des=`sn,state,valid`,sep=`,` 分隔des后，src是在des中的|   
    | NotIn |  src string,des string,sep string | 是否不包括字符串，与In相反 |   

+ .根属性(在模版以.attrName使用)

    |属性名| 类型 | 说明|
    | :---- | :---- | :---- |
    | date |  string | 日期(yyyy-MM-dd)，比如用于@date注释 | 
    | hasServiceInterface |  bool | 是否生成service接口，与.gen中属性一样，只是类型不一样 | 
    | hasDate |  bool | 是否有java中的Date类型 | 
    | hasBigDecimal |  bool | 是否有java中的BigDecimal类型 | 
    | columnNumber |  string | 数据表的列数量 | 
    | primaryKeyName |  string | 实体主键名，统一为：id | 
    | primaryKeyJdbcType |  string | 主键mysql类型 | 
    | primaryKeyJavaType |  string | 主键java类型 | 
    | primaryKeyColumn |  string | 主键列名 | 
    | hasModule |  bool | 是否有模块名 | 
    | mod |  string | 模块名称 |
    | javaName |  string | 表格对应的Java名称(去模块名并下划线转驼峰且首字母大写)如：mem_user_favorite对应 UserFavorite | 
    | serialVersionUID |  int | 序列号ID | 
    | -------- | --------- | ------------------------ |   
    | table |  对象 | 数据表信息 | 
    | columns | 集合 | 数据表的列信息集合 | 
    | gen |  对象 | 通用的模版设置，如下方的.gen变量说明 | 

+ .gen 变量属性(在模版需要以.gen.AttrName使用)  

    |属性名| 类型 | 说明|  
    | :---- | :---- | :---- |
    | Author|  string |  作者，用于@author注释 |      
    | DaoParentClass|  string | dao父全类名 | 
    | DaoPkg|  string | dao包名 | 
    | DaoSuffix |  string | dao类名后辍 如：DAO| 
    | DoParentClass |  string | 表映射实体父全类名 | 
    | DoPkg |  string | 实体类包名 | 
    | DoSuffix |  string | 实体类名后辍 | 
    | HasModule |  string | 是否生成模块名，值：on 表示开启，将解析表名前辍(第一个下划的位置)为模块名称 |                                       
    | ServiceParentClass |  string | service父全类名 |  
    | ServicePkg |  string | service包名 | 
    | ServiceSuffix |  string | service类名后辍 | 
    | HasServiceInterface |  string | 是否生成service接口类，值: on 表示开启， | 
    | ServiceInterfaceParentClass |  string | service接口父全类名 | 
    | MapperXmlSuffix |  string | mapper.xml文件名后辍 | 
    | Tables |  string | 表名列表 | 
    
+ .table 变量属性(在模版中以.table.AttrName使用)
    
    |属性名| 类型 | 说明|  
    | :---- | :---- | :---- | 
    | Name |  string | 数据表名 | 
    | Comment |  string | 数据表的注释 | 
    | CreateTime |  string | 数据表的创建时间 | 

+ .columns 变量属性(在模版中以.columns使用)
    
    |属性名| 类型 | 说明|  
    | :---- | :---- | :---- | 
    | Name |  string | 表字段名 | 
    | DataType |  string | 字段数据库类型(小写) | 
    | Comment |  string | 注释 | 
    | Key |  string | 键，如主键为：PRI | 
    | Extra |  string |  | 
    | JavaType |  string | 对应java类型名称(如：Integer,Long,Date,String) | 
    | JavaName |  string | 对应java名称(首字母小写) | 
    

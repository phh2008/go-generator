package mysqlutil

var mysqlGoTypeMap map[string]string

func init() {
	//mysql类型-java类型映射
	mysqlGoTypeMap = map[string]string{}
	//长整型
	mysqlGoTypeMap["bigint"] = "int64"
	//整型
	mysqlGoTypeMap["int"] = "int"
	mysqlGoTypeMap["tinyint"] = "int"
	mysqlGoTypeMap["smallint"] = "int"
	mysqlGoTypeMap["mediumint"] = "int"
	mysqlGoTypeMap["integer"] = "int"
	//小数
	mysqlGoTypeMap["float"] = "float"
	mysqlGoTypeMap["double"] = "float64"
	mysqlGoTypeMap["decimal"] = "decimal.Decimal"
	//bool
	mysqlGoTypeMap["bit"] = "bool"
	//字符串
	mysqlGoTypeMap["char"] = "string"
	mysqlGoTypeMap["varchar"] = "string"
	mysqlGoTypeMap["tinytext"] = "string"
	mysqlGoTypeMap["text"] = "string"
	mysqlGoTypeMap["mediumtext"] = "string"
	mysqlGoTypeMap["longtext"] = "string"
	//日期
	mysqlGoTypeMap["date"] = "time.Time"
	mysqlGoTypeMap["datetime"] = "time.Time"
	mysqlGoTypeMap["timestamp"] = "int64"
	//其它统一为字符串
}

func GetGoType(mysqlType string) string {
	goType := mysqlGoTypeMap[mysqlType]
	if goType == "" {
		return "string"
	}
	return goType
}

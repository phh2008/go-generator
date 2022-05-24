package dao

import (
	"com.phh/go-generator/domain"
	"com.phh/go-generator/utils/dbutil"
)

type GeneratorDao interface {
	// QueryTableList 获取表列表
	QueryTableList(name string) []domain.TableName
	// GetTableByTableName 根据表名获取表信息
	GetTableByTableName(tableName string) domain.TableName
	// GetTableColumnsByTableName 获取表的列
	GetTableColumnsByTableName(tableName string) []domain.Column
}

func NewGeneratorDao() GeneratorDao {
	return &generatorDaoMapper{}
}

type generatorDaoMapper struct {
}

func (g *generatorDaoMapper) QueryTableList(name string) []domain.TableName {
	var list []domain.TableName
	db := dbutil.Db
	db = db.Select("TABLE_NAME,TABLE_COMMENT,CREATE_TIME").
		Where("table_schema = (SELECT DATABASE())")
	if name != "" {
		db = db.Where("table_name LIKE ?", "%"+name+"%")
	}
	db.Find(&list)
	return list
}

func (g *generatorDaoMapper) GetTableByTableName(tableName string) domain.TableName {
	var table domain.TableName
	db := dbutil.Db
	db.Select("TABLE_NAME,TABLE_COMMENT,CREATE_TIME").
		Where("table_schema = (SELECT DATABASE()) AND table_name=?", tableName).
		Find(&table)
	return table
}

func (g *generatorDaoMapper) GetTableColumnsByTableName(tableName string) []domain.Column {
	var list []domain.Column
	db := dbutil.Db
	db.Select("COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT,COLUMN_KEY,EXTRA").
		Where("table_name = ? AND table_schema = (SELECT DATABASE())", tableName).
		Order("ordinal_position").
		Find(&list)
	return list
}

package domain

type TableName struct {
	Name    string `gorm:"column:TABLE_NAME"`
	Comment string `gorm:"column:TABLE_COMMENT"`
	CreateTime   string `gorm:"column:CREATE_TIME"`
}

func (TableName) TableName() string {
	return "information_schema.TABLES"
}
package domain

type Column struct {
	Name     string `gorm:"column:COLUMN_NAME"`
	DataType string `gorm:"column:DATA_TYPE"`
	Comment  string `gorm:"column:COLUMN_COMMENT"`
	Key      string `gorm:"column:COLUMN_KEY"`
	Extra    string `gorm:"column:EXTRA"`
	GoType   string `gorm:"-"`
	GoName   string `gorm:"-"`
}

func (Column) TableName() string {
	return "information_schema.COLUMNS"
}

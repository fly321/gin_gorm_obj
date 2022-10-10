package models

type CategoryBasic struct {
	Common
	Name     string `gorm:"column:name;type:varchar(255);" json:"name"`      //分类名称
	ParentId int    `gorm:"column:parent_id;type:int(11);" json:"parent_id"` //父级id
}

// TableName 配置表名
func (table *CategoryBasic) TableName() string {
	return "category_basic"
}

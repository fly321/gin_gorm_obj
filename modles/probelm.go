package modles

type Problem struct {
	Common
	CategoryId string `gorm:"column:category_id;type:varchar(255);" json:"category_id"` //问题分类id
	Title      string `gorm:"column:title;type:varchar(255);" json:"title"`             //问题标题
	Content    string `gorm:"column:content;type:varchar(255);" json:"content"`         //问题内容
	MaxMem     int    `gorm:"column:max_mem;type:int(11);" json:"max_mem"`              //最大内存
	MaxRuntime int    `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"`      //最大运行时长
}

// TableName 配置表名
func (table *Problem) TableName() string {
	return "problem"
}

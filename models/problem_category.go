package models

import "gorm.io/gorm"

type ProblemCategory struct {
	gorm.Model
	ProblemId     int            `gorm:"column:problem_id;type:int(11);" json:"problem_id"`   //问题id
	CategoryId    int            `gorm:"column:category_id;type:int(11);" json:"category_id"` //分类id
	CategoryBasic *CategoryBasic `gorm:"foreignKey:id;references:category_id"`                // 关联分类的基础信息表
}

func (table *ProblemCategory) TableName() string {
	return "problem_category"
}

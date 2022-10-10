package modles

import (
	"fmt"
	"gorm.io/gorm"
)

type ProblemBasic struct {
	Common
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id" json:"problem-categorys,omitempty"`                // 关联问题分类表
	CategoryId        string             `gorm:"column:category_id;type:varchar(255);" json:"category_id" json:"category-id,omitempty"` //问题分类id
	Title             string             `gorm:"column:title;type:varchar(255);" json:"title" json:"title,omitempty"`                   //问题标题
	Content           string             `gorm:"column:content;type:varchar(255);" json:"content" json:"content,omitempty"`             //问题内容
	MaxMem            int                `gorm:"column:max_mem;type:int(11);" json:"max_mem" json:"max-mem,omitempty"`                  //最大内存
	MaxRuntime        int                `gorm:"column:max_runtime;type:int(11);" json:"max_runtime" json:"max-runtime,omitempty"`      //最大运行时长
}

// TableName 配置表名
func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(keyword, categoryIdentity string) *gorm.DB {
	keywords := fmt.Sprintf("%%%s%%", keyword)
	fmt.Println(keywords, "key", "%"+keyword+"%")
	tx := Db.Model(new(ProblemBasic)).Preload("ProblemCategories").
		Preload("ProblemCategories.CategoryBasic").
		Where("problem_basic.title like ? or problem_basic.content like ?", keywords, keywords)

	if categoryIdentity != "" {
		tx.Joins("right join problem_category pc on pc.problem_id = problem_basic.id").
			Where("pc.category_id = (select cb.id from category_basic cb where cb.identity = ?)", categoryIdentity)
	}

	return tx
}

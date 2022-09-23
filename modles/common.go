package modles

import "gorm.io/gorm"

type Common struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"` //表的唯一标识
}

package models

type UserBasic struct {
	Common
	Username string `gorm:"column:username;type:varchar(100);" json:"name"`    //用户名
	Password string `gorm:"column:password;type:varchar(32);" json:"password"` //密码
	Mail     string `gorm:"column:mail;type:varchar(255);" json:"mail"`        //邮箱
	Phone    string `gorm:"column:phone;type:varchar(20);" json:"phone"`       //手机号
	IsAdmin  int    `gorm:"column:is_admin;type:int(1);" json:"is_admin"`      //是否是管理员
}

// TableName 配置表名
func (table *UserBasic) TableName() string {
	return "user_basic"
}

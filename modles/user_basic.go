package modles

type UserBasic struct {
	Common
	Name     string `gorm:"column:name;type:varchar(100);" json:"name"`        //用户名
	Password string `gorm:"column:password;type:varchar(32);" json:"password"` //密码
	Email    string `gorm:"column:email;type:varchar(255);" json:"email"`      //邮箱
	Phone    string `gorm:"column:phone;type:varchar(20);" json:"phone"`       //手机号
}

// TableName 配置表名
func (table *UserBasic) TableName() string {
	return "user"
}

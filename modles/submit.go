package modles

type Submit struct {
	Common
	UserIdentity    string `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`       //用户唯一标识
	ProblemIdentity string `gorm:"column:problem_identity;type:varchar(36);" json:"problem_identity"` //问题唯一标识
	Identity        string `gorm:"column:identity;type:varchar(36);" json:"identity"`                 //提交唯一标识
	Path            string `gorm:"column:path;type:varchar(255);" json:"path"`                        //提交文件路径
}

// TableName 配置表名
func (table *Submit) TableName() string {
	return "submit"
}

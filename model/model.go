package model

type User struct {
	ID       int    `gorm:"column:ID;type:int;primaryKey"`
	Name     string `gorm:"column:NAME;type:varchar(64);not null;comment:登录名称;uniqueIndex:udx_name"`
	Password string `gorm:"column:PASSWD;type:varchar(128);not null;comment:登录密码"`
	Nickname string `gorm:"column:NICKNAME;type:varchar(64);not null;comment:昵称"`
}

// 指定表名为“T_USER”
func (User) TableName() string {
	return "T_USER"
}

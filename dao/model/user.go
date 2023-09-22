package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"type:varchar(20);not null;comment:‘用户名’"`
	Password string `json:"password" form:"password" gorm:"type:varchar(100);not null;comment:‘密码’"`
	Nickname string `json:"nickname" form:"nickname" gorm:"type:varchar(20);not null;comment:‘昵称’"`
	Level    byte   `json:"level" gorm:"default:1;comment:‘用户级别’"`
	Phone    string `json:"phone" form:"phone" gorm:"type:varchar(15);default:'';comment:手机号"`
	Email    string `json:"email" form:"email" gorm:"type:varchar(20);default:'';comment:‘邮箱’"`
	Avatar   string `json:"avatar" form:"avatar" gorm:"type:varchar(100);default:'https://i2.hdslb.com/bfs/face/f6a42463ab0378b6ccfa9fadcee81866afe679ad.jpg@120w_120h_1c.webp';comment:‘头像’"`
}

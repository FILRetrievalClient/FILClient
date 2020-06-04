package user

import (
	"github.com/jinzhu/gorm"
)

// 用户数据模型
type User struct {
	gorm.Model
	Name		string	`gorm:"not null"`
	Username	string	`gorm:"unique"`
	Password	string	`gorm:"not null"`
}

// 返回接口
type Response struct {
	Status	bool		`json:"status"`
	Msg 	interface{}	`json:"msg"`
	Data 	interface{}	`json:"data"`
}


package user

import (
	"FILClient/models/db"
	"github.com/jinzhu/gorm"
)

type AuthToken struct {
	gorm.Model
	Token	string	`gorm:"not null default '' comment('Token')"`
	UserId 	uint	`gorm:"not null default '' comment('UserId')" `
	Secret 	string	`gorm:"not null default '' comment('Secret')"`
	ExpressIn 	int64	`gorm:"not null default 0 comment('是否是标准库')"`
	Revoked 	bool
}

type Token struct {
	Token 	string	`json:"token"`
}

func (at *AuthToken)AuthTokenCreate()(response Token)  {
	db.GetDB().Create(at)
	response = Token{at.Token}
	return
}


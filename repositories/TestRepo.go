package repositories

import (
	"FILClient/models/db"
	"github.com/jinzhu/gorm"
)

type TestRepositories struct {
	db *gorm.DB
}

func NewTestRepositories() *TestRepositories {
	return &TestRepositories{db: db.DB.Mysql}
}
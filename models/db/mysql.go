package db

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
)

// 判断返回值是否返回 ErrRecordNotFound,如果是则说明无相关记录
func IsNotFound(err error)  {
	if ok := errors.Is(err,gorm.ErrRecordNotFound);!ok && err!=nil{
		color.Red(fmt.Sprintf("error :%v \n",err))
	}
}

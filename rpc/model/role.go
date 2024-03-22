package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `gorm:"not null;unique"`
	Code string `gorm:"not null;unique"`
	//DefaultRouter string 登录页面先注释一下
	Remark string
}

package model

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Status    int64  `gorm:"default:1"` //0 禁用|1 正常使用
	Uuid      string `gorm:"not null"`
	Token     string `gorm:"not null"`
	ExpiredAt int64  `gorm:"not null"`
}

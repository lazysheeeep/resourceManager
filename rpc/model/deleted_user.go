package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeletedUser struct {
	gorm.Model
	Username       string `gorm:"unique;not null"`
	PasswordDigest string `gorm:"not null"`
	LoginStatus    int    `gorm:"default:0"` //0 禁用|1 正常
	Avatar         string
	Email          string
	Phone          string    `gorm:"type varchar(11)"`
	RoleId         int       `gorm:"default:1"` //0 管理员|1 普通用户
	UUID           uuid.UUID `gorm:"not null;unique"`
}

package model

import (
	"time"
)

type DeletedUser struct {
	Id             string    `gorm:"primaryKey;unique;not null"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt      time.Time
	Username       string `gorm:"not null"`
	PasswordDigest string `gorm:"not null"`
	LoginStatus    int    `gorm:"default:0"` //0 禁用|1 正常
	Avatar         string
	Email          string
	Phone          string `gorm:"type varchar(11)"`
	RoleId         int    `gorm:"default:1"` //0 管理员|1 普通用户
}

package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id             string    `gorm:"primaryKey;unique;not null"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Username       string    `gorm:"unique;not null"`
	PasswordDigest string    `gorm:"not null"`
	LoginStatus    int       `gorm:"default:1"` //0 禁用|1 正常
	Avatar         string
	Email          string `gorm:"unique"`
	Phone          string `gorm:"type varchar(11);unique"`
	RoleId         uint   `gorm:"not null"` //1 管理员|2 审核|3 普通用户
}

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

func CheckPassword(password, digit string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(digit), []byte(password))
	return err == nil
}

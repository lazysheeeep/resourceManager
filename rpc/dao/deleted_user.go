package dao

import (
	"gorm.io/gorm"
	"resourceManager/rpc/model"
)

type DeletedUserDao struct {
	*gorm.DB
}

func NewDeletedUserDao(db *gorm.DB) *DeletedUserDao {
	return &DeletedUserDao{db}
}

func (dao *DeletedUserDao) Create(user model.DeletedUser) error {
	return dao.Model(&model.DeletedUser{}).Create(&user).Error
}

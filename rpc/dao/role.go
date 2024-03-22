package dao

import (
	"gorm.io/gorm"
	"resourceManager/rpc/model"
)

type RoleDao struct {
	*gorm.DB
}

func NewRoleDao(db *gorm.DB) *RoleDao {
	return &RoleDao{db}
}

func (dao RoleDao) CreateRole(role model.Role) error {
	return dao.Model(&model.Role{}).Create(&role).Error
}

func (dao RoleDao) GerRoleById(id uint) (role model.Role, err error) {
	err = dao.Model(&model.Role{}).Where("id=?", id).Find(&role).Error
	return
}

package dao

import (
	"gorm.io/gorm"
	"resourceManager/rpc/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (dao *UserDao) Create(user model.User) error {
	return dao.Model(&model.User{}).Create(&user).Error

}

func (dao *UserDao) GetByUserName(name string) (user model.User, exist bool, err error) {
	var number int64
	err = dao.Model(&model.User{}).Where("username=?", name).Find(&user).Count(&number).Error
	if err != nil {
		return model.User{}, false, err
	}
	if number == 0 {
		return user, false, err
	}
	return user, true, nil
}

func (dao *UserDao) GetByUserId(id string) (user model.User, exist bool, err error) {
	var number int64
	err = dao.Model(&model.User{}).Where("id=?", id).Find(&user).Count(&number).Error
	if err != nil {
		return model.User{}, false, err
	}
	if number == 0 {
		return model.User{}, false, nil
	}
	return user, true, nil
}

func (dao *UserDao) UpdateUser(username string, user model.User) error {
	return dao.Model(&model.User{}).Where("username=?", username).Updates(&user).Error
}

func (dao *UserDao) DeleteUser(user model.User) error {
	return dao.Model(&model.User{}).Unscoped().Delete(&user).Error
}

func (dao *UserDao) GetUserList(page model.Page) (user []model.User, count int64, err error) {
	err = dao.Model(&model.User{}).Offset(int((page.PageNum - 1) * page.PageSize)).Limit(int(page.PageSize)).Find(&user).Count(&count).Error
	return
}

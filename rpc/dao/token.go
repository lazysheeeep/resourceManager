package dao

import (
	"gorm.io/gorm"
	"resourceManager/rpc/model"
)

type TokenDao struct {
	*gorm.DB
}

func NewTokenDao(db *gorm.DB) *TokenDao {
	return &TokenDao{db}
}

func (dao *TokenDao) CreateToken(token model.Token) error {
	return dao.Model(&model.Token{}).Create(&token).Error
}

func (dao *TokenDao) GetTokenById(id uint) (token model.Token, err error) {
	err = dao.Model(&model.Token{}).Where("id=?", id).Find(token).Error
	return
}

func (dao *TokenDao) GetLatestToken(uuid string) (token model.Token, err error) {
	err = dao.Model(&model.Token{}).Order("updated_at desc").Where("uuid=?", uuid).First(&token).Error
	return
}

func (dao *TokenDao) GetTokens(page model.Page) (tokens []model.Token, count int64, err error) {
	err = dao.Model(&model.Token{}).Offset(int((page.PageNum - 1) * page.PageSize)).Limit(int(page.PageSize)).Find(&tokens).Count(&count).Error
	return
}

func (dao *TokenDao) DeleteToken(tokens model.Token) error {
	return dao.Model(&model.Token{}).Delete(&tokens).Error
}

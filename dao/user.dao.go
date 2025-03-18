/*
 * @Author: lihuan
 * @Date: 2024-09-02 21:14:58
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-23 22:32:07
 * @Email: 17719495105@163.com
 */
package dao

import (
	"context"

	"github.com/smartgreeting/mini-go/models"
	"gorm.io/gorm"
)

type UserDao struct {
	db  *gorm.DB
	ctx context.Context
}

func NewUserDao(ctx context.Context, db *gorm.DB) *UserDao {
	return &UserDao{
		ctx: ctx,
		db:  db,
	}
}
func (u UserDao) FindUserInfoById(id int64) (*models.User, error) {

	var user models.User
	err := u.db.Where("id = ?", id).First(&user).Error

	return &user, err
}
func (u UserDao) DelById(id int64) error {
	err := u.db.Delete(&models.User{}, id).Error

	return err
}
func (u UserDao) CreatUserInfo(user *models.User) (*models.User, error) {
	r := u.db.First(&user, "openid= ?", user.Openid)
	if r.RowsAffected == 0 {
		err := u.db.Create(&user).Error
		return user, err
	}
	return user, nil
}

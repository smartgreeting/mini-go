/*
 * @Author: lihuan
 * @Date: 2024-09-02 21:14:58
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-03 21:40:17
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
	err := u.db.Where("id = ? AND deleted = ?", id, 0).First(&user).Error

	return &user, err
}

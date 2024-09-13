/*
 * @Author: lihuan
 * @Date: 2024-09-03 20:51:36
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-12 21:04:59
 * @Email: 17719495105@163.com
 */
package svc

import (
	"context"

	"github.com/smartgreeting/mini-go/dao"
	"github.com/smartgreeting/mini-go/database"
	"github.com/smartgreeting/mini-go/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SvcContext struct {
	Config  *utils.Conf
	DB      *gorm.DB
	UserDao *dao.UserDao
	RedisDB *database.RedisDB
}

var ctx = context.Background()

func NewSvcContext(c *utils.Conf) *SvcContext {

	db, err := gorm.Open(mysql.Open(c.MySql.Dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	rdb := database.NewRedisDB(ctx, c)
	return &SvcContext{
		Config:  c,
		DB:      db,
		RedisDB: rdb,
		UserDao: dao.NewUserDao(ctx, db.Debug()),
	}
}

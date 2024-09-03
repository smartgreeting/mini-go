/*
 * @Author: lihuan
 * @Date: 2024-09-03 20:51:36
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-03 22:15:15
 * @Email: 17719495105@163.com
 */
package svc

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/smartgreeting/mini-go/dao"
	"github.com/smartgreeting/mini-go/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SvcContext struct {
	Config  *utils.Conf
	DB      *gorm.DB
	UserDao *dao.UserDao
	RedisDB *redis.Client
}

func NewSvcContext(c *utils.Conf) *SvcContext {

	db, err := gorm.Open(mysql.Open(c.MySql.Dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Dns,
		Password: c.Redis.Pass,
		DB:       c.Redis.DB,
	})
	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return &SvcContext{
		Config:  c,
		DB:      db,
		RedisDB: rdb,
		UserDao: dao.NewUserDao(context.Background(), db),
	}
}

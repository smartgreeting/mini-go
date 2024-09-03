/*
 * @Author: lihuan
 * @Date: 2021-12-13 20:15:52
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-03 22:08:07
 * @Email: 17719495105@163.com
 */
package utils

import (
	"fmt"
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Conf struct {
	Application Application
	Token       Token
	Md5         Md5
	MySql       MySql
	Redis       Redis
}

type Token struct {
	Secret     string
	ExpireTime int `yaml:"expire_time"`
}
type Md5 struct {
	Secret string
}
type MySql struct {
	Dns string
}
type Redis struct {
	Dns  string
	Pass string
	DB   int
}

type Application struct {
	Address string
	Port    int
	Mode    string
}

var Cfg *Conf = nil

func InitConf(path string) (*Conf, error) {
	// 加载文件
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 将读取的yaml文件解析为响应的 struct
	err = yaml.Unmarshal(yamlFile, &Cfg)

	return Cfg, err
}

var (
	db   *gorm.DB
	lock sync.Mutex
)

func GetDB() *gorm.DB {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		_db, err := gorm.Open(mysql.Open(Cfg.MySql.Dns), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		db = _db
	}
	return db
}

/*
 * @Author: lihuan
 * @Date: 2024-08-27 20:42:11
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-18 21:07:29
 * @Email: 17719495105@163.com
 */
package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	Application Application
	Token       Token
	Md5         Md5
	MySql       MySql
	Redis       Redis
	WX          WX
	Log         Log
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
type WX struct {
	AppID     string `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
}
type Log struct {
	LogFilePath string `yaml:"log_file_path"`
	LogFileName string `yaml:"log_file_name"`
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

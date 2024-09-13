/*
 * @Author: lihuan
 * @Date: 2024-09-02 21:06:55
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-13 20:04:04
 * @Email: 17719495105@163.com
 */
package models

import "gorm.io/gorm"

type User struct {
	ID        uint
	Username  string
	Password  string
	Avatar    string
	Gender    int8
	Phone     string
	Email     string
	Address   string
	Hobbies   string
	CreatedAt uint
	UpdatedAt uint `json:"updatedAt"`
	DeletedAt gorm.DeletedAt
}

func (u User) TableName() string {
	return "hc_user"
}

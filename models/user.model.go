/*
 * @Author: lihuan
 * @Date: 2024-09-02 21:06:55
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-23 21:52:22
 * @Email: 17719495105@163.com
 */
package models

import "gorm.io/gorm"

type User struct {
	ID        uint           `json:"id"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Avatar    string         `json:"avatar"`
	Gender    int8           `json:"gender"`
	Phone     string         `json:"phone"`
	Email     string         `json:"email"`
	Address   string         `json:"address"`
	Hobbies   string         `json:"hobbies"`
	Openid    string         `json:"openid"`
	CreatedAt uint           `json:"createdAt"`
	UpdatedAt uint           `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (u User) TableName() string {
	return "hc_user"
}

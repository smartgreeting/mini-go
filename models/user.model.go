/*
 * @Author: lihuan
 * @Date: 2024-09-02 21:06:55
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-03 20:30:13
 * @Email: 17719495105@163.com
 */
package models

type User struct {
	ID        int64
	Username  string
	Password  string
	Avatar    string
	Gender    int32
	Phone     string
	Email     string
	Address   string
	Hobbies   string
	Deleted   int32
	CreatedAt int32
	UpdatedAt int32
}

func (u User) TableName() string {
	return "hc_user"
}

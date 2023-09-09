package models

import "gorm.io/gorm"

// 对于匿名字段，GORM 会将其字段包含在父结构体中
type User struct {
	gorm.Model
	Nickname  string
	Avatar    string
	Phone     string
	Email     string
	Wx        string
	PhoneType string
}

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}

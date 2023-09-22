package models

import "gorm.io/gorm"

type Proxy struct {
	gorm.Model
	Proxy string
	Port  string
	Type  string
}

func (Proxy) TableName() string {
	return "proxy"
}

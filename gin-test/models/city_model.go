package models

import "gorm.io/gorm"

type SubLevel struct {
	gorm.Model
	Name              string `json:"name"`
	SubLevelModelList []City `json:"subLevelModelList" gorm:"foreignKey:LevelType"`
}

type City struct {
	gorm.Model
	Name      string `json:"name"`
	Code      int    `json:"code"`
	LevelType uint   `json:"levelType"`
}

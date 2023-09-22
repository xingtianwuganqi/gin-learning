package models

import (
	"gorm.io/gorm"
	"time"
)

type Position struct {
	Code              float64        `json:"code" gorm:"primaryKey"`
	SubLevelModelList []SubLevelType `gorm:"foreignKey:PositionCode"`
	ParentCode        float64        `json:"parentCode"`
	Name              string         `json:"name"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type SubLevelType struct {
	Code              float64         `json:"code" gorm:"primaryKey"`
	SubLevelModelList []SubLevelModel `gorm:"foreignKey:SubLevelTypeCode"`
	PositionCode      float64         `json:"parentCode"`
	Name              string          `json:"name"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type SubLevelModel struct {
	Code              float64 `json:"code" gorm:"primaryKey"`
	SubLevelModelList *string `json:"subLevelModelList"`
	SubLevelTypeCode  float64 `json:"parentCode"`
	Name              string  `json:"name"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

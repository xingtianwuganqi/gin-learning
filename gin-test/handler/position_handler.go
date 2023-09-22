package handler

import (
	"gin-test/db"
	"gin-test/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ReadLocalTxtFile(c *gin.Context) {
	// 插入之前先判断数据库中是否已经有了
	var positions []models.Position
	result := db.DB.Find(&positions)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 402,
			"msg":  "fail",
		})
		return
	}
	if len(positions) > 0 {
		c.JSON(http.StatusCreated, gin.H{
			"code": 202,
			"msg":  "fail",
		})
		return
	}
	ReadLocalPosition(func(m []map[string]interface{}) {
		if len(m) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "fail",
			})
		}
		// 保存到数据库
		var positions []models.Position
		for i := 0; i < len(m); i++ {
			var subLevelArr []models.SubLevelType
			subLevels := m[i]["subLevelModelList"].([]interface{})
			for j := 0; j < len(subLevels); j++ {
				subInfo := subLevels[j].(map[string]interface{})
				var levelTypeArr []models.SubLevelModel
				var levelType = subInfo["subLevelModelList"].([]interface{})
				for k := 0; k < len(levelType); k++ {
					levelInfo := levelType[k].(map[string]interface{})
					var levelModel = models.SubLevelModel{
						Name:             levelInfo["name"].(string),
						Code:             levelInfo["code"].(float64),
						SubLevelTypeCode: levelInfo["parentCode"].(float64),
					}
					levelTypeArr = append(levelTypeArr, levelModel)
				}
				levelTypeModel := models.SubLevelType{
					Name:              subInfo["name"].(string),
					Code:              subInfo["code"].(float64),
					PositionCode:      subInfo["parentCode"].(float64),
					SubLevelModelList: levelTypeArr,
				}
				subLevelArr = append(subLevelArr, levelTypeModel)
			}
			position := models.Position{
				Name:              m[i]["name"].(string),
				Code:              m[i]["code"].(float64),
				ParentCode:        m[i]["parentCode"].(float64),
				SubLevelModelList: subLevelArr,
			}
			positions = append(positions, position)
		}
		result := db.DB.Model(&models.Position{}).Create(&positions)
		if result.Error != nil {
			log.Println(result.Error)
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "fail",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": positions,
		})
	})
}

func PositionSearch(c *gin.Context) {
	var position []models.Position
	result := db.DB.Find(&position)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": position,
		"msg":  "success",
	})
}

func SubLevelSearch(c *gin.Context) {
	parentCode := c.PostForm("code")
	var subLevel models.SubLevelType
	result := db.DB.Where("parent_code_info = ?", parentCode).Find(&subLevel)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": subLevel,
	})
}

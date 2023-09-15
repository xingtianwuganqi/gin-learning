package handler

import (
	"gin-test/db"
	"gin-test/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func FindCityHandler(c *gin.Context) {

	// 先查询是否已经插入了
	var subLevel []models.SubLevel
	subData := db.DB.Find(&subLevel)
	if subData.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "fails",
		})
		return
	}
	var citys []models.City
	cityData := db.DB.Find(&citys)
	if cityData.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "fails",
		})
		return
	}
	if subData.RowsAffected != 0 && cityData.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 300,
			"msg":  "fails",
		})
		return
	}
	GetCityJson(func(m []map[string]interface{}) {
		// 将m保持到数据库中
		// 根据map创建
		// 使用类型断言将其转换回 map[string]interface{}
		var subLevels []models.SubLevel
		for i := 0; i < len(m); i++ {
			subCitys, _ := m[i]["subLevelModelList"].([]map[string]interface{})
			var subs []models.City
			for j := 0; j < len(subCitys); j++ {
				sub := models.City{
					Code:      subCitys[j]["code"].(int),
					Name:      subCitys[j]["name"].(string),
					LevelType: uint(subCitys[j]["levelType"].(int)),
				}
				subs = append(subs, sub)
			}
			value := models.SubLevel{
				Name:              m[i]["name"].(string),
				SubLevelModelList: subs,
			}
			subLevels = append(subLevels, value)
		}
		result := db.DB.Model(&models.SubLevel{}).Create(&subLevels)
		if result.Error != nil {
			log.Println(result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "插入失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": m,
		})
	})
}

func GetSubInfo(c *gin.Context) {
	var sub []models.SubLevel
	result := db.DB.Find(&sub)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": sub,
		"msg":  "success",
	})
}

func GetSubCitys(c *gin.Context) {
	var subs []models.SubLevel

	result := db.DB.Preload("SubLevelModelList").Find(&subs)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": subs,
		"msg":  "success",
	})
}

func GetCityCode(c *gin.Context) {
	var name = c.Query("city")
	log.Println(name)
	var city models.City
	result := db.DB.Where("name = ?", name).Find(&city)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": city,
		"msg":  "success",
	})
}

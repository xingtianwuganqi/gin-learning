package handler

import (
	"gin-test/db"
	"gin-test/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserModel struct {
	NickName string `json:"nickname" form:"nickname"`
	Wx       string `json:"wx" form:"age"`
}

func UserRegister(c *gin.Context) {
	name := c.PostForm("username")
	age := c.PostForm("age")
	avator := c.PostForm("avator")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	user := models.User{
		Nickname:  name,
		Avatar:    avator,
		Phone:     phone,
		Email:     email,
		Wx:        age,
		PhoneType: "ios",
	}
	result := db.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user,
		"msg":  "success",
	})
}

func DeleteUser(c *gin.Context) {
	userId := c.PostForm("userId")
	// 删除1
	//num, err := strconv.ParseUint(userId, 10, 0)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{})
	//	return
	//}
	//user := models.User{Model: gorm.Model{ID: uint(num)}}
	//result := db.DB.Delete(&user)
	// 删除2
	//result := db.DB.Where("Id=?", userId).Delete(&models.User{})
	// 删除3
	result := db.DB.Delete(&models.User{}, userId)
	if result.Error != nil {
		c.JSON(http.StatusAccepted, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func FindUnDelete(c *gin.Context) {
	var users []models.User
	//// 查询未被删除的数据
	//result := db.DB.Where("deleted_at IS NULL").Find(&users)

	// 如果字段被删除，普通查询不会返回被删除的字段
	//result := db.DB.Where("nickname IS NOT NULL").Find(&users)

	//result := db.DB.Take(&users)

	//result := map[string]interface{}{}
	//db.DB.Model(&models.User{}).First(&result)

	//result := map[string]interface{}{}
	//db.DB.Model(&models.User{}).Last(&result)

	// 会查询到被删掉的数据
	//result := map[string]interface{}{}
	//db.DB.Table("users").Take(&result)

	// 查询单个id
	//result := db.DB.Find(&users, 5)
	//result := db.DB.First(&users, 4)

	// 查询指定id
	//result := db.DB.Find(&users, []int{4, 5})
	//result := db.DB.First(&users, "id = ?", "5")

	//user := models.User{Model: gorm.Model{ID: 5}}
	//result := db.DB.First(&user)

	//var user models.User
	//result := db.DB.First(&user)

	// 条件查询
	//result := db.DB.Where("nickname = ?", "Ethan").Find(&users)
	// 查询nickname != Ethan 的信息
	//result := db.DB.Where("nickname <> ?", "Ethan").First(&users)

	// 查询nickname是否在数组中
	//result := db.DB.Where("nickname IN ?", []string{"Ethan"}).Find(&users)
	// 查询nickname像
	//result := db.DB.Where("nickname LIKE ?", "%哈%").Find(&users)
	//result := db.DB.Where("nickname = ? AND wx > ?", "Ethan", 10).Find(&users)

	// 使用struct查询
	//result := db.DB.Where(&models.User{Nickname: "Ethan"}).Find(&users)
	// 使用map查询
	//result := db.DB.Where(map[string]interface{}{"wx": 34}).Find(&users)
	// 使用切片查询
	//result := db.DB.Where([]uint{4, 5}).Find(&users)

	// 内联查询
	//result := db.DB.Find(&users, "nickname = ?", "Ethan")

	// not 条件
	//result := db.DB.Not("nickname = ?", "Ethan").Find(&users)

	// not in
	/// 判断nickname不在[]string中的
	//result := db.DB.Not(map[string]interface{}{"nickname": []string{"哈哈哈哈"}}).Find(&users)
	// 使用struct查询
	//result := db.DB.Not(models.User{Nickname: "Ethan"}).First(&users)
	// 使用int数组
	//result := db.DB.Not([]uint{1, 2, 3, 4}).Find(&users)

	// or 条件
	//result := db.DB.Where("id = ?", 4).Or("id = ?", 5).Find(&users)
	// 支持struct
	//result := db.DB.Where("nickname = ?", "Ethan").Or(models.User{Model: gorm.Model{ID: 5}}).Find(&users)
	// 支持map
	//result := db.DB.Where("nickname = ?", "哈哈哈哈").Or(map[string]interface{}{"wx":18}).Find(&users)

	// 选定特定字段
	//result := db.DB.Select("nickname", "wx").Find(&users)

	// 排序
	//result := db.DB.Order("wx").Find(&users)
	// 分页
	//result := db.DB.Limit(1).Offset(0).Find(&users)
	/*
		limit 一页多少数据
		offset 偏移量
		offset = (pagesize - 1) * limit
	*/
	result := db.DB.Limit(1).Offset(1).Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "success",
		"data":  users,
		"count": result.RowsAffected,
		//"user":  user,
	})
}

func UpdateUser(c *gin.Context) {
	userId := c.PostForm("userId")
	log.Println(userId)
	var userModel UserModel
	var user models.User
	result := db.DB.Where("id = ?", userId).Find(&user)
	if result.Error != nil {
		c.JSON(http.StatusAccepted, gin.H{})
		return
	}
	if c.ShouldBind(&userModel) == nil {
		log.Println(userModel.NickName)
		log.Println(userModel.Wx)
		//user.Nickname = userModel.NickName
		//user.Wx = userModel.Wx
		//db.DB.Save(&user)

		// 更新,先将user查出，再更新column
		//db.DB.Model(&user).Update("nickname", "mol")

		// 更新多个参数
		// 使用结构体更新
		//db.DB.Model(&user).Updates(models.User{Nickname: "swift", Wx: "10"})
		//db.DB.Model(&user).Updates(map[string]interface{}{"nickname": "Text", "wx": 30})

		// 只更新选中的字段
		//db.DB.Model(&user).Select("nickname").Updates(map[string]interface{}{"nickname": "go", "wx": 18})
		// 忽略
		//db.DB.Model(&user).Omit("nickname").Updates(map[string]interface{}{"nickname": "gosin", "wx": 18, "email": "rxjava@126.com"})
		//
		//// 选择所有更新
		//db.DB.Model(&user).Select("*").Updates(map[string]interface{}{"nickname": "gosin", "wx": 30, "email": "rxjava@126.com"})
		//// 忽略所有
		//db.DB.Model(&user).Select("*").Updates(map[string]interface{}{"nickname": "gosin", "wx": 30, "email": "rxjava@126.com"})
		//
		// 批量更新
		//db.DB.Model(&models.User{}).Where("wx < ?", 20).Updates(map[string]interface{}{"nickname": "gosin", "wx": 30})
		db.DB.Table("users").Where("id IN ?", []int{5, 6, 7}).Updates(map[string]interface{}{"wx": 50})
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": user,
		})
	}
}

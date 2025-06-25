package service

import (
	"carbon/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type rankRequest struct {
	Month int `json:"month"`
}

type userRank struct {
	Name        string `json:"name"`
	CarbonValue int    `json:"carbon_value"`
}

func RankList(c *gin.Context) {
	var req rankRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Failed(c, err.Error())
		return
	}

	if req.Month < 0 || req.Month > 12 {
		MonthError(c)
		return
	}

	uid, _ := c.Get("curUser")
	id := uid.(uint)
	db := model.GetDB()
	var curUser model.Auth
	if err := db.Where("id = ?", id).First(&curUser).Error; err != nil {
		UserIsNotExist(c)
		return
	}

	if req.Month == 0 {
		rankSumList(c, db)
		return
	}

	var carbonRank []userRank
	db.Model(&model.User{}).Select("name, carbon as carbon_value").Where("`month` = ?", req.Month).
		Order("carbon_value desc").Find(&carbonRank)
	Success(c, carbonRank)
}

func rankSumList(c *gin.Context, db *gorm.DB) {
	var carbonSumRank []userRank
	db.Model(&model.User{}).Select("name, sum(carbon) as carbon_value").
		Group("name").Order("carbon_value desc").Find(&carbonSumRank)
	Success(c, carbonSumRank)
}

func Rank(c *gin.Context) {
	var req rankRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Failed(c, err.Error())
		return
	}

	if req.Month < 0 || req.Month > 12 {
		MonthError(c)
		return
	}

	uid, _ := c.Get("curUser")
	id := uid.(uint)
	db := model.GetDB()

	var curUser model.Auth
	if err := db.Where("id = ?", id).First(&curUser).Error; err != nil {
		UserIsNotExist(c)
		return
	}

	var user model.User
	db.Where("name = ?", curUser.Name).First(&user)

	if req.Month == 0 {
		rankSum(c, db, user.Name)
		return
	}

	var rank int
	db.Raw("select count(*) + 1 from user where carbon > ? and `month` = ?", user.Carbon, req.Month).
		Scan(&rank)
	Success(c, rank)
}

func rankSum(c *gin.Context, db *gorm.DB, uname string) {
	var rank int
	subQuery := db.Model(&model.User{}).Select("name, sum(carbon) as carbon_sum").Group("name")
	carbonSum := db.Model(&model.User{}).Select("sum(carbon)").Where("name = ?", uname)
	db.Raw("select count(*) + 1 from (?) as ranked_user where ranked_user.carbon_sum > (?)", subQuery, carbonSum).Scan(&rank)
	Success(c, rank)
}

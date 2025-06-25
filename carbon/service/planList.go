package service

import (
	"carbon/model"
	"github.com/gin-gonic/gin"
)

type list struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	Finished int    `json:"finished"`
}

func PlanList(c *gin.Context) {
	uid, _ := c.Get("curUser")
	id := uid.(uint)
	db := model.GetDB()
	var curUser model.Auth
	if err := db.Where("id = ?", id).First(&curUser).Error; err != nil {
		Response(c, 0, "用户不存在", uid)
		return
	}

	var lists []list
	db.Model(&model.Plan{}).Where("uid = ?", id).Find(&lists)
	Success(c, lists)
}

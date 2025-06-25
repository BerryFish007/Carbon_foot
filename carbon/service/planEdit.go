package service

import (
	"carbon/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type editRequest struct {
	ID      int    `json:"ID" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func PlanEdit(c *gin.Context) {
	var req editRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Failed(c, err.Error())
		return
	}

	uid, _ := c.Get("curUser")
	id := uid.(uint)
	db := model.GetDB()
	var curUser model.Auth
	if err := db.Where("id = ?", id).First(&curUser).Error; err != nil {
		Response(c, 0, "用户不存在", uid)
		return
	}

	err := db.Model(&model.Plan{}).Where("id = ?", req.ID).
		Update("content", req.Content).Error
	if err != nil {
		Failed(c, err.Error())
	}
	Success(c, nil)
}

type finishedRequest struct {
	ID int `json:"ID" binding:"required"`
}

func PlanFinished(c *gin.Context) {
	var req finishedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Failed(c, err.Error())
		return
	}

	uid, _ := c.Get("curUser")
	id := uid.(uint)
	db := model.GetDB()
	var curUser model.Auth
	if err := db.Where("id = ?", id).First(&curUser).Error; err != nil {
		Response(c, 0, "用户不存在", uid)
		return
	}

	err := db.Model(&model.Plan{}).Where("id = ?", req.ID).
		Update("finished", gorm.Expr("1 - finished")).Error
	if err != nil {
		Failed(c, err.Error())
		return
	}

	Success(c, nil)
}

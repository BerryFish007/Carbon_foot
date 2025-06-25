package service

import (
	"carbon/model"
	"github.com/gin-gonic/gin"
)

type deleteRequest struct {
	ID int `json:"id" binding:"required"`
}

func PlanDelete(c *gin.Context) {
	var req deleteRequest
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

	err := db.Delete(&model.Plan{}, "id = ?", req.ID).Error
	if err != nil {
		Failed(c, err.Error())
		return
	}
	Success(c, nil)
}

package service

import (
	"carbon/model"
	"github.com/gin-gonic/gin"
)

type planRequest struct {
	Content string `json:"content" binding:"required"`
}

func PlanAdd(c *gin.Context) {
	var req planRequest
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

	newPlan := model.Plan{
		Uid:      int(id),
		Content:  req.Content,
		Finished: 0,
	}
	db.Create(&newPlan)
	Success(c, newPlan.ID)
}

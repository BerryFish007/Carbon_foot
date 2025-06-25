package service

import (
	"carbon/middleware"
	"carbon/model"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Failed(c, err.Error())
		return
	}

	db := model.GetDB()
	var user model.Auth
	db.Where("name = ?", req.Username).First(&user)
	if user.ID == 0 {
		UserIsNotExist(c)
		return
	}

	token, err := middleware.ReleaseToken(&user)
	if err != nil {
		Failed(c, "token生成失败")
		return
	}
	Success(c, token)
}

package service

import (
	"carbon/model"
	"github.com/gin-gonic/gin"
)

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Failed(c, err.Error())
		return
	}

	db := model.GetDB()
	var user model.Auth
	db.Where("name = ?", req.Username).First(&user)
	if user.ID != 0 {
		UsernameIsUsed(c)
		return
	}

	newUser := model.Auth{
		Name:     req.Username,
		Password: req.Password,
	}
	db.Create(&newUser)
	Success(c, nil)
}

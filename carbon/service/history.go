package service

import (
	"carbon/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

type historyData struct {
	Carbon    int             `json:"carbon"`
	Month     int             `json:"month"`
	Detail    json.RawMessage `json:"detail"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func History(c *gin.Context) {
	uid, _ := c.Get("curUser")
	id := uid.(uint)
	db := model.GetDB()
	var curUser model.Auth
	if err := db.Where("id = ?", id).First(&curUser).Error; err != nil {
		Response(c, 0, "用户不存在", uid)
		return
	}

	var history []historyData
	err := db.Model(&model.User{}).Order("updated_at desc").Find(&history).Error
	if err != nil {
		Failed(c, err.Error())
		return
	}
	Success(c, history)
}

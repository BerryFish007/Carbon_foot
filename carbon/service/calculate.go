package service

import (
	"carbon/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type calculate struct {
	Month  int             `json:"month" binding:"required"`
	Carbon int             `json:"carbon" binding:"required"`
	Detail json.RawMessage `json:"detail" binding:"required"`
}

func Calculate(c *gin.Context) {
	var req calculate
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

	newRecord := model.User{
		Name:   curUser.Name,
		Carbon: req.Carbon,
		Month:  req.Month,
		Detail: req.Detail,
	}

	//err := db.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "name"}, {Name: "month"}},
	//	DoUpdates: clause.AssignmentColumns([]string{"carbon"}),
	//}).Create(&newRecord).Error
	//if err != nil {
	//	Failed(c, err.Error())
	//	return
	//}

	db.Create(&newRecord)

	Success(c, nil)
	//if err := db.Model(model.User{}).Where("name = ? and `month` = ?", req.Name, req.Month).First(&user).Error; err != nil {
	//	if err = db.Model(&user).Update("carbon", req.Carbon).Error; err != nil {
	//		Failed(c, err.Error())
	//	}
	//} else if errors.Is(err, gorm.ErrRecordNotFound) {
	//	newRecord := model.User{
	//		Name:   req.Name,
	//		Carbon: req.Carbon,
	//		Month:  req.Month,
	//	}
	//	if err = db.Model(&model.User{}).Create(&newRecord).Error; err != nil {
	//		Failed(c, err.Error())
	//	}
	//}
}

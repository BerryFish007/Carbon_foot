package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string          `gorm:"type:varchar(100);not null" json:"name"`
	Month  int             `gorm:"type:int;not null" json:"month"`
	Carbon int             `json:"carbon"`
	Detail json.RawMessage `json:"detail"`
}

func (User) TableName() string {
	return "user"
}

type Auth struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
}

func (Auth) TableName() string {
	return "auth"
}

type Plan struct {
	gorm.Model
	Uid      int    `json:"uid"`
	Content  string `gorm:"type:text" json:"content"`
	Finished int    `json:"finished"`
}

func (Plan) TableName() string {
	return "plan"
}

//type User struct {
//	gorm.Model
//	Name   string `gorm:"type:varchar(100);not null" json:"name"`
//	Month  int    `gorm:"type:int;not null" json:"month"`
//	Carbon int    `json:"carbon"`
//}

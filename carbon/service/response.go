package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type resJson struct {
	Code int
	Msg  string
	Data any
}

func Response(c *gin.Context, code int, msg string, data any) {
	json := resJson{Code: code, Msg: msg, Data: data}
	c.JSON(http.StatusOK, json)
}

func Failed(c *gin.Context, msg string) {
	Response(c, 0, msg, nil)
}

func Success(c *gin.Context, data any) {
	Response(c, 1, "成功", data)
}

func UsernameIsUsed(c *gin.Context) {
	Response(c, 101, "用户名已被占用", nil)
}

func UserIsNotExist(c *gin.Context) {
	Response(c, 102, "用户名不存在", nil)
}

func MonthError(c *gin.Context) {
	Response(c, 103, "月份出错", nil)
}

package router

import (
	"carbon/middleware"
	"carbon/service"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Use(middleware.CORSMiddleware())
	r.POST("/register", service.Register)
	r.POST("/login", service.Login)

	u := r.Group("/user")
	u.Use(middleware.AuthMiddleware())
	//u.GET("/rank", service.Rank)
	//u.GET("/rankList", service.RankList)
	u.POST("/calculate", service.Calculate)
	u.GET("/history", service.History)

	p := r.Group("/plan")
	p.Use(middleware.AuthMiddleware())
	p.GET("/list", service.PlanList)
	p.POST("/add", service.PlanAdd)
	p.POST("/delete", service.PlanDelete)
	p.POST("/edit", service.PlanEdit)
	p.POST("/finished", service.PlanFinished)
}

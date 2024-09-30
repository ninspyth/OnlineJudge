package router 

import (
	"github.com/ninspyth/OnlineJudge/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//Initialize gin Engine
	r := gin.Default()
	
	//API endpoints
	r.POST("/submit", handlers.SubmitHandle)
	r.POST("/run", handlers.RunHandle)

	return r
}

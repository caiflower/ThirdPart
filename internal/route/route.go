package route

import (
	"com.caiflower/commons/thirdpart/internal/common"
	"com.caiflower/commons/thirdpart/internal/config"
	v1 "com.caiflower/commons/thirdpart/internal/route/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var engine *gin.Engine

func Init() (e error) {
	config := config.Config
	if config.Port == "" {
		e = &common.ThirdError{Msg: "port不能为空"}
		return
	}

	engine = gin.Default()
	engine.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	initRouteGroup(engine)

	s := &http.Server{
		Addr:         ":" + config.Port,
		ReadTimeout:  time.Duration(config.RT * int(time.Second)),
		WriteTimeout: time.Duration(config.WT * int(time.Second)),
		Handler:      engine,
	}
	s.ListenAndServe()

	return
}

func initRouteGroup(engine *gin.Engine) {
	g1 := engine.Group("/email")
	{
		g1.POST("/string", v1.SendString)
		g1.POST("/byte", v1.SendByte)
	}
}

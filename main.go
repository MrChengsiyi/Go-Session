package main

import (
	"Session/session"
	"github.com/gin-gonic/gin"
)



func main()  {
	r:=gin.Default()
	r.LoadHTMLGlob("templates/*")
	session.InitMgr("redis")
	r.Use(session.SessionMid(session.MgrObj))
	r.Any("/index",indexHandle)
	r.GET("/loaded",AuthMidHandle,loadedhandle)
	r.Run()
}

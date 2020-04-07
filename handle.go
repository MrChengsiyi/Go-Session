package main

import (
	"Session/session"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func AuthMidHandle(c *gin.Context)  {
	SessionData,_:=c.Get("session_data")
	sd:=SessionData.(session.SessionData)
	value,err:=sd.Get("isLogin")
	if err!=nil{
		fmt.Println("err:",err)
		c.Redirect(http.StatusFound,"/index")
		return
	}
	isLogin,ok:=value.(bool)
	if !ok{
		fmt.Println("ok :",ok)
		c.Redirect(http.StatusFound,"/index")
		return
	}

	if !isLogin{
		fmt.Println("isLogin :",isLogin)
		c.Redirect(http.StatusFound,"/index")
		return
	}
	c.Next()
}

func indexHandle(c *gin.Context)  {
	if c.Request.Method=="POST"{
		fmt.Println("POST")
		var user User
		c.ShouldBind(&user)
		if user.Username=="csy" && user.Password=="416171575csy"{
			sessionData,ok:=c.Get("session_data")
			if !ok{
				panic("session middleware failed")
			}
			sd:= sessionData.(session.SessionData)
			sd.Set("isLogin",true)
			sd.Save()
			c.Redirect(http.StatusFound,"/loaded")
		}else{
			c.HTML(http.StatusOK,"index.html",gin.H{
				"err":"用户名或密码错误",
			})
		}
	}else{
		fmt.Println("GET")
		c.HTML(http.StatusOK,"index.html",nil)
	}

}

func loadedhandle(c *gin.Context)  {
	c.String(http.StatusOK,"登陆成功")
}


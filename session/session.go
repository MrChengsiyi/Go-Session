package session

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"sync"
)

type SessionData interface {
	GetID() (id string)
	Get(key string) (value interface{}, err error)
	Set(key string, value interface{})
	Save()
}

type Mgr interface {
	Init(string)
	GetSessionData(string) (SessionData,error)
	CreateSession()(sd SessionData)

}

var MgrObj Mgr

func InitMgr(name string,option ...string)  {
	switch name {
		case "memory":
			MgrObj=NewMemManager()
			fmt.Println("memory")
		case "redis":
			MgrObj=NewRedisManager()
			fmt.Println("redis")
	}
	MgrObj.Init("127.0.0.1:6379")
}


func SessionMid(m Mgr) gin.HandlerFunc  {
	return func(c *gin.Context){
		var sd SessionData
		session_id,err:=c.Cookie("Cookie_name")
		if err!=nil{
			sd=m.CreateSession()
			session_id=sd.GetID()
			fmt.Println("1111")
		}else{
			sd,err=m.GetSessionData(session_id)
			if err!=nil{
				sd=m.CreateSession()
				session_id=sd.GetID()
				fmt.Println("2222")
			}
		}

		c.Set("session_data",sd)
		c.SetCookie("Cookie_name", session_id,  600, "/", "127.0.0.1", false,true)
		c.Next()
	}
}





package session

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type RedisSD struct {
	ID string
	Data map[string]interface{}
	rwlock sync.RWMutex
	client *redis.Client
	expired int //过期时间
}

type RedisManager struct {
	Session  map[string] SessionData
	rwlock sync.RWMutex
	client *redis.Client

}


func NewRedisManager()(r *RedisManager)  {
	return &RedisManager{Session:make(map[string] SessionData,1024)}
}

func (r *RedisManager)GetSessionData(sessionID string) (sd SessionData,err error) {
	r.rwlock.RLock()
	defer r.rwlock.RUnlock()
	sd,ok:=r.Session[sessionID]
	if !ok{
		err=fmt.Errorf("invalid session id")
		return
	}
	return
}

func NewRedisSD(id string,client *redis.Client) (session SessionData ){
	return &RedisSD{
		ID:     id,
		Data:   make(map[string]interface{},8),
		client:client,
	}
}

func (r *RedisManager) CreateSession()(sd SessionData)  {
	uuidObj,_:=uuid.NewV4()
	sd=NewRedisSD(uuidObj.String(),r.client)
	r.Session[sd.GetID()]=sd
	return
}


func (r *RedisManager)Init(addr string){
	r.client=redis.NewClient(&redis.Options{
		Addr:               addr,
		Password:           "",
		DB:                 0,
	})
	_, err:= r.client.Ping().Result()
	if err != nil {
		panic(err)
	}
}



func (sd *RedisSD)GetID()(id string) {
	return sd.ID
}

func (sd *RedisSD)Get(key string) (value interface{},err error) {
	sd.rwlock.RLock()
	defer sd.rwlock.RUnlock()
	value,ok:=sd.Data[key]
	if !ok{
		err=fmt.Errorf("invalid key")
		return
	}
	return
}

func (sd *RedisSD)Set(key string,value interface{}) {
	sd.rwlock.Lock()
	defer sd.rwlock.Unlock()
	sd.Data[key]=value
}

func (sd *RedisSD)Save()  {
	sd.expired=20
	value,err:=json.Marshal(sd.Data)
	if err != nil {
		fmt.Println("Marshal failed ,err: ",err)
		return
	}
	sd.client.Set(sd.ID,value,time.Second*time.Duration(sd.expired))

}

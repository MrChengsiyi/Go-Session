package session

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type MemSD struct {
	ID string
	Data map[string]interface{}
	rwlock sync.RWMutex
}

type MemManager struct {
	Session  map[string] SessionData
	rwlock sync.RWMutex
}

func NewMemManager() (m *MemManager) {
	return &MemManager{Session:make(map[string] SessionData,1024)}
}

func (manager *MemManager)GetSessionData(sessionID string) (sd SessionData,err error) {
	manager.rwlock.RLock()
	defer manager.rwlock.RUnlock()
	sd,ok:=manager.Session[sessionID]
	if !ok{
		err=fmt.Errorf("invalid session id")
		return
	}
	return
}

func NewMemSD(id string) (session SessionData ){
	return &MemSD{
		ID:     id,
		Data:   make(map[string]interface{},8),
	}
}


func (manager *MemManager) CreateSession()(sd SessionData)  {
	uuidObj,_:=uuid.NewV4()
	sd=NewMemSD(uuidObj.String())
	manager.Session[sd.GetID()]=sd
	return
}

//func (manager *MemManager) Save()  {
//	return
//}

func (manager *MemManager)Init(addr string){
	return
}

func (sd *MemSD)GetID()(id string) {
	return sd.ID
}

func (sd *MemSD)Get(key string) (value interface{},err error) {
	sd.rwlock.RLock()
	defer sd.rwlock.RUnlock()
	value,ok:=sd.Data[key]
	if !ok{
		err=fmt.Errorf("invalid key")
		return
	}
	return
}

func (sd *MemSD)Set(key string,value interface{}) {
	sd.rwlock.Lock()
	defer sd.rwlock.Unlock()
	sd.Data[key]=value
}

func (sd *MemSD)Save()  {
	return
}
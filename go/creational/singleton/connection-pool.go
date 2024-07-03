package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type connectionPool struct {
}

var connPoolInstance *connectionPool

func GetInstance() *connectionPool {
	if connPoolInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if connPoolInstance == nil {
			fmt.Println("Creating Connection Pool Instance Now")
			connPoolInstance = &connectionPool{}
		} else {
			fmt.Println("Connection Pool Instance already created-1")
		}
	} else {
		fmt.Println("Connection Pool Instance already created-2")
	}
	return connPoolInstance
}

var once sync.Once

func GetInstanceUsingSync() *connectionPool {
	if connPoolInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating Connection Pool Instance Now")
				connPoolInstance = &connectionPool{}
			})
	} else {
		fmt.Println("Connection Pool Instance already created-2")
	}
	return connPoolInstance
}

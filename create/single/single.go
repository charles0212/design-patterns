package single

import (
	"fmt"
	"sync"
)

var once sync.Once
var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func getInstanceSafe() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single safe instance now.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single safe instance already created.")
	}
	return singleInstance
}

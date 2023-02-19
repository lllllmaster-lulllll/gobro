package singleton

import "sync"

// todo:Singleton 饿汉式单例
type Singleton struct {
}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

// todo:获取实例
func GetInstance() *Singleton {
	return singleton
}

// todo: 懒汉式
var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}

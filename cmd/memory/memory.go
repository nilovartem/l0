package memory

import (
	"time"

	"github.com/nilovartem/l0/cmd/model"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

var Cache *cache.Cache

// check for database and restore cache
func Restore() {

}
func Set(key string, value *model.Order) {
	Cache.Set(key, value, cache.NoExpiration)
	logrus.Infoln("[SUCCESS] Message was set in cache")
}
func Get(key string) /*model.Order*/ {
	result, found := Cache.Get(key)
	if found {
		logrus.Infof("[SUCCESS] Message was found in cache - %v", result)
	}
}
func New() {
	if Cache == nil {
		Cache = cache.New(5*time.Minute, 10*time.Minute)
		Restore()
	}

}
func delete() {

}

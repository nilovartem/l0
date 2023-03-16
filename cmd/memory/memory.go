package memory

import (
	"time"

	"github.com/nilovartem/l0/cmd/database"
	"github.com/nilovartem/l0/cmd/model"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

var Cache *cache.Cache

// проверяем, нуждается ли кеш в восстановлении и при необходимости восстанавливаем его из бд
func Restore() {
	orders := database.GetOrders()
	if len(orders) == 0 {
		logrus.Warnln("The database is empty, recovery is impossible")
		return
	}
	for key, value := range orders {
		Set(key, value)
	}
	logrus.Infoln("[SUCCESS] The cache was recovered from the database")
}
func Set(key string, value *model.Order) {
	Cache.Set(key, value, cache.NoExpiration)
	logrus.Infoln("[SUCCESS] Message was set in cache")
}
func Get(key string) *model.Order {
	if Cache != nil {
		result, found := Cache.Get(key)
		if found {
			var Order *model.Order = result.(*model.Order)
			logrus.Infoln("[SUCCESS] Message was found in cache")
			return Order
		}
	}
	return nil

}
func New() {
	if Cache == nil {
		//TODO: мб стоит сделать кеш вечным, ну или чуть длиннее
		Cache = cache.New(5*time.Minute, 10*time.Minute)
		Restore()
	}

}

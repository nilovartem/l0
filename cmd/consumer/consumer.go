package consumer

import (
	"net/url"
	"runtime"

	"github.com/nats-io/stan.go"
	"github.com/nilovartem/l0/cmd/config"
	"github.com/nilovartem/l0/cmd/memory"
	"github.com/nilovartem/l0/cmd/model"
	"github.com/nilovartem/l0/cmd/server"

	"github.com/sirupsen/logrus"
)

var (
	cfg config.Config
)

func getURL() string {
	url := url.URL{
		Scheme: cfg.STAN.Scheme,
		Host:   cfg.STAN.Host + cfg.STAN.Port,
	}
	return url.String()
}
func getConfig() { cfg.GetConfig() }

// Читаем канал

// Пишем данные в кэш и в БД - (ДОПОЛНИТЕЛЬНОЕ УСЛОВИЕ: ЕСЛИ ДАННЫЕ ИЗМЕНИЛИСЬ)
func caching() {

}
func save() {

}

// Читаем кэш и отображаем полученные данные
func restore() {

}
func runServer() {
	server.Run()
}

// Обмениваемся данными через каналы?
/*
Основная функция приложения consumer
*/
func Run() {
	logrus.Infoln("Starting consumer")
	getConfig()
	sc, err := stan.Connect(cfg.STAN.ClusterID, cfg.STAN.ConsumerID, stan.NatsURL(getURL()))
	if err != nil {
		logrus.Error("[FAIL] Failed to connect to STAN, exiting")
		return
	}
	_, err = sc.Subscribe(cfg.STAN.Channel,
		func(m *stan.Msg) {
			logrus.Infof("[SUCCESS] Message on subject '%v' was received", m.Subject)
			var order *model.Order
			order, _ = model.Unmarshal(m.Data)
			if order != nil {
				logrus.Infoln("[SUCCESS] Message was unmarshalled")
				//TODO: Caching
				memory.New()
				memory.Set(cfg.Cache.ModelKey, order)

			} else {
				logrus.Warnln("[FAIL] Failed to unmarshal received message, needs retry transmission")
			}
		})
	if err != nil {
		panic(err)
	}
	runtime.Goexit()
}

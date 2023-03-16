package server

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/julienschmidt/httprouter"
	"github.com/nilovartem/l0/cmd/config"
	"github.com/nilovartem/l0/cmd/memory"
	"github.com/nilovartem/l0/cmd/model"
	"github.com/sirupsen/logrus"
)

var (
	cfg  config.Config
	tmpl *template.Template
)

type Order struct {
	Uid  string
	Data string
}

func getConfig() { cfg.GetConfig() }

func getOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	logrus.Infof("Attempt to get order %v", uid)
	var order *model.Order = memory.Get(uid)
	orderJSON, _ := json.Marshal(order)
	//шаблончики
	Order := Order{uid, string(orderJSON)}
	tmpl.Execute(w, Order)
}

func Run() {
	getConfig()
	logrus.Infoln("Starting server")
	templateFilename, _ := filepath.Abs("../server/index.html")
	tmpl = template.Must(template.ParseFiles(templateFilename))
	router := httprouter.New()
	router.GET("/:uid", getOrder)
	err := http.ListenAndServe(":3333", router)
	//TODO: реализовать нормальную обработку ошибок
	if err != nil {
		logrus.Errorln("[FAIL] Critical error on server")
	}
	runtime.Goexit()
}

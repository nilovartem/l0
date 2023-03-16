package database

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/nilovartem/l0/cmd/config"
	"github.com/nilovartem/l0/cmd/model"
	"github.com/sirupsen/logrus"
)

var (
	db  *sql.DB
	cfg config.Config
)

func connect() error {
	cfg.GetConfig()
	var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	db, err = sql.Open("postgres", psqlconn)
	if err == nil {
		logrus.Infoln("[SUCCESS] Connection to DB was established")
		return nil
	}
	return err

}

func disconnect() {
	err := db.Close()
	if err != nil {
		logrus.Infoln("[SUCCESS] Connection to DB was aborted")
	}
}

// нужен для восстановления кеша из бд
func GetOrders() map[string]*model.Order {
	orders := make(map[string]*model.Order)
	err := connect()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(`SELECT "uid", "data" FROM "model"`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var uid string
		var data string

		err = rows.Scan(&uid, &data)
		if err != nil {
			return nil
		}
		var order *model.Order
		order, err = model.Unmarshal([]byte(data))
		if err != nil {
			return nil
		}
		orders[uid] = order

	}
	return orders
}
func InsertOrder(order *model.Order) {
	err := connect()
	if err != nil {
		panic(err)
	}
	orderJSON, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}
	insertStatement := `insert into "model"("uid", "data") values($1, $2)`
	_, err = db.Exec(insertStatement, order.OrderUID, orderJSON)
	if err != nil {
		logrus.Warn("[FAIL] Message is already presented in database")
	}
	logrus.Infoln("[SUCCESS] Message was saved in database")
	disconnect()
}

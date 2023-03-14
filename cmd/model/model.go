package model

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

// Тип, отображающий структуру необходимого JSON
type Order struct {
	OrderUID          string    `json:"order_uid" validate:"required"`
	TrackNumber       string    `json:"track_number" validate:"required"`
	Entry             string    `json:"entry" validate:"required"`
	Delivery          Delivery  `json:"delivery" validate:"required"`
	Payment           Payment   `json:"payment" validate:"required"`
	Items             []Items   `json:"items" validate:"required"`
	Locale            string    `json:"locale" validate:"required"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" validate:"required"`
	Shardkey          string    `json:"shardkey" validate:"required"`
	SmID              int       `json:"sm_id" validate:"required"`
	DateCreated       time.Time `json:"date_created" validate:"required"`
	OofShard          string    `json:"oof_shard" validate:"required"`
}
type Delivery struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Zip     string `json:"zip" validate:"required"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required"`
	Email   string `json:"email" validate:"required"`
}
type Payment struct {
	Transaction  string `json:"transaction" validate:"required"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency" validate:"required"`
	Provider     string `json:"provider" validate:"required"`
	Amount       int    `json:"amount" validate:"required"`
	PaymentDt    int    `json:"payment_dt" validate:"required"`
	Bank         string `json:"bank" validate:"required"`
	DeliveryCost int    `json:"delivery_cost" validate:"required"`
	GoodsTotal   int    `json:"goods_total" validate:"required"`
	CustomFee    int    `json:"custom_fee"`
}
type Items struct {
	ChrtID      int    `json:"chrt_id" validate:"required"`
	TrackNumber string `json:"track_number" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Rid         string `json:"rid" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Sale        int    `json:"sale" validate:"required"`
	Size        string `json:"size" validate:"required"`
	TotalPrice  int    `json:"total_price" validate:"required"`
	NmID        int    `json:"nm_id" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	Status      int    `json:"status" validate:"required"`
}

/*
validateOrder проверяет JSON на соответствие полям структуры Order
*/
func validateOrder(order *Order) error {
	var validator = validator.New()
	err := validator.Struct(order)
	if err != nil {
		logrus.Errorln("[FAIL] Cannot validate message fields")
		logrus.Errorf("\n%v", err)
		return err
	}
	return nil
}

/*
Unmarshal преобразовывает переданное сообщение в структуру JSON, проводя
проверки содержимого
*/
func Unmarshal(data []byte) (*Order, error) {
	logrus.Infoln("Trying to unmarshal message")
	var Order Order
	err := json.Unmarshal(data, &Order)
	//с ошибками одно слово - цирк
	if err != nil {
		return nil, err //не знал, что делать с ошибкой, пришлось делать 2 возвращаемых значения
	}
	logrus.Infoln("[SUCCESS] JSON was recongnized, starting validation")
	err = validateOrder(&Order)
	if err != nil {
		return nil, err //не знал, что делать с ошибкой, пришлось делать 2 возвращаемых значения
	}
	return &Order, nil
}

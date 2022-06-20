package repository

import (
	"database/sql"
	"time"

	"github.com/Questee29/taxi-app_orderService/models/order"
	_ "github.com/lib/pq"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) *repository {
	return &repository{db: db}
}

func (repository *repository) GetAll() (order.Order, error) {
	return order.Order{}, nil
}
func (repository *repository) CreateOrderTest(order order.UserRequest) error {
	query := `
	INSERT into test
	VALUES ($1,$2,$3,$4)`
	if _, err := repository.db.Exec(query, order.ID, order.TaxiType, order.From, order.To); err != nil {
		return err
	}
	return nil
}

func (repository *repository) CreateOrder(order order.Order) error {
	query := `
	INSERT into orders(user_id,driver_id,from,to,taxi_typename,date,status) 
	VALUES ($1,$2,$3,$4,$5,$6,"in progress")`
	if _, err := repository.db.Exec(query, order.UserId, order.DriverId, order.From, order.To, order.TaxiType, time.Now()); err != nil {
		return err
	}
	return nil
}

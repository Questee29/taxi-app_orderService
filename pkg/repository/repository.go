package repository

import (
	"context"
	"database/sql"
	"log"

	model "github.com/Questee29/taxi-app_orderService/models/order"
	_ "github.com/lib/pq"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) *repository {
	return &repository{db: db}
}

func (repository *repository) GetAll() (model.Order, error) {
	return model.Order{}, nil
}

func (repository *repository) FindFreeDriver(ctx context.Context, taxiType string) (model.DriverResponse, error) {
	return model.DriverResponse{}, nil
}
func (repository *repository) CreateOrder(order model.Order) error {
	var zeroRate float32 = 0
	query := `
	INSERT into orders(user_id,driver_id,start_point,end_point,taxi_type,order_date,status,user_rate,driver_rate)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	if _, err := repository.db.Exec(query, order.UserId, order.DriverId, order.From, order.To, order.TaxiType, order.Date, order.Status, zeroRate, zeroRate); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

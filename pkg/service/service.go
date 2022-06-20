package service

import (
	"context"
	"log"
	"time"

	model "github.com/Questee29/taxi-app_orderService/models/order"
	pb "github.com/Questee29/taxi-app_orderService/proto/protob"
)

type Repository interface {
	FindFreeDriver(ctx context.Context, taxiType string) (model.DriverResponse, error)
	CreateOrder(order model.Order) error
	GetAll() (model.Order, error)
}

type service struct {
	pb.UnimplementedOrderGrpcServer
	repository Repository
}

func New(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

//execute
func (service *service) NewOrder(ctx context.Context, user model.UserRequest, driver model.DriverResponse) error {
	order := model.Order{
		UserId:     uint32(user.ID),
		DriverId:   uint32(driver.Driverid),
		From:       user.From,
		To:         user.To,
		TaxiType:   user.TaxiType,
		Date:       time.Now().Format("01-02-2006 15:04:05"), //date + time
		Status:     "in progress",
		UserRate:   0,
		DriverRate: 0,
	}

	if err := service.repository.CreateOrder(order); err != nil {
		log.Println("ERROR REPo")
		return err
	}

	return nil
}
func (service *service) FindFreeDriver(ctx context.Context, user model.UserRequest) (model.DriverResponse, error) {
	return model.DriverResponse{}, nil
}
func (service *service) CreateOrder(order model.Order) error {
	return service.repository.CreateOrder(order)

}
func (service *service) GetAll() (model.Order, error) {
	return model.Order{}, nil
}

package service

import (
	"context"

	model "github.com/Questee29/taxi-app_orderService/models/order"
	pb "github.com/Questee29/taxi-app_orderService/proto/protob"
)

type Repository interface {
	CreateOrder(model.Order) error
	GetAll() (model.Order, error)
	CreateOrderTest(order model.UserRequest) error
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

func (service *service) OrderTaxi(ctx context.Context, user model.UserRequest) error {
	if err := service.repository.CreateOrderTest(user); err != nil {
		return err
	}

	return nil
}
func (service *service) FindDriver(ctx context.Context, req *pb.FindDriverRequest) (*pb.FindDriverResponse, error) {
	return nil, nil
}
func (service *service) CreateOrder(order model.Order) error {
	return service.repository.CreateOrder(order)

}
func (service *service) GetAll() (model.Order, error) {
	return model.Order{}, nil
}

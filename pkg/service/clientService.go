package service

type ClientRepository interface {

	//GetAll() (model.Order, error)
}

type clientService struct {
}

func NewClientService() *clientService {
	return &clientService{}

}

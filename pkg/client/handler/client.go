package handler

type ClientService interface {
}

type ClientHandler struct {
	service ClientService
}

func NewClientHandler(service ClientService) *ClientHandler {
	return &ClientHandler{
		service: service,
	}
}

package handler

import "github.com/Kolakanmi/grey_transaction/service"

type Handler struct {
	service service.ITransactionService
}

func New(service service.ITransactionService) *Handler {
	return &Handler{service: service}
}

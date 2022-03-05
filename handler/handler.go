package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Kolakanmi/grey_transaction/pkg/http/response"
	"github.com/Kolakanmi/grey_transaction/service"
)

type Handler struct {
	service service.ITransactionService
}

func New(service service.ITransactionService) *Handler {
	return &Handler{service: service}
}

//Credit -
// @Summary Credit wallet
// @Description Credit wallet
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Param user body service.TxnRequest true "Body"
// @Success 200 {object} response.RespBody{data=service.TxnResponse} "desc"
// @Router /credit [post]
func (h *Handler) Credit(w http.ResponseWriter, r *http.Request) error {
	var req service.TxnRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	result, err := h.service.Credit(r.Context(), req.Amount)
	if err != nil {
		return response.Fail(err).ToJSON(w) //err
	}
	return response.OK("Success", result).ToJSON(w)
}

//Debit -
// @Summary Debit wallet
// @Description Debit wallet
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Param user body service.TxnRequest true "Body"
// @Success 200 {object} response.RespBody{data=service.TxnResponse} "desc"
// @Router /debit [post]
func (h *Handler) Debit(w http.ResponseWriter, r *http.Request) error {
	var req service.TxnRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	result, err := h.service.Debit(r.Context(), req.Amount)
	if err != nil {
		return response.Fail(err).ToJSON(w) //err
	}
	return response.OK("Success", result).ToJSON(w)
}

//Balance -
// @Summary Wallet balance
// @Description Wallet balance
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Success 200 {object} response.RespBody{data=service.TxnResponse} "desc"
// @Router /balance [get]
func (h *Handler) Balance(w http.ResponseWriter, r *http.Request) error {
	result, err := h.service.Balance(r.Context())
	if err != nil {
		return response.Fail(err).ToJSON(w) //err
	}
	return response.OK("Success", result).ToJSON(w)
}

func (h *Handler) GetAllTransactions(w http.ResponseWriter, r *http.Request) error {
	result, err := h.service.GetAll(r.Context())
	if err != nil {
		return response.Fail(err).ToJSON(w) //err
	}
	return response.OK("Success", result).ToJSON(w)
}

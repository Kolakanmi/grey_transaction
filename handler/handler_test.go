package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	proto "github.com/Kolakanmi/grey_transaction/pkg/grpc/transaction"
	"github.com/Kolakanmi/grey_transaction/pkg/http/response"
	"github.com/Kolakanmi/grey_transaction/repository/mock"
	"github.com/Kolakanmi/grey_transaction/service"
	mockConn "github.com/Kolakanmi/grey_transaction/service/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getMockhandler() *Handler {
	repo := mock.NewMockRepo()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(mockConn.Dialer()))
	if err != nil {
		log.Fatal(err)
	}
	wc := proto.NewWalletClient(conn)

	service := service.NewTransactionService(repo, wc)

	return New(service)
}

func TestCredit(t *testing.T) {
	tests := []struct {
		name   string
		amount float64
		res    response.RespBody
	}{
		{
			name:   "When credit amount is greater than 0",
			amount: 10,
			res: response.RespBody{
				Success: true,
				Data:    service.TxnResponse{Balance: 110},
				Error:   "",
			},
		},
		{
			name:   "When credit amount is less than 0",
			amount: -10,
			res: response.RespBody{
				Success: false,
				Data:    nil,
				Error:   "Amount cannot be negative.",
			},
		},
		{
			name:   "When credit amount is equal to 0",
			amount: 0,
			res: response.RespBody{
				Success: false,
				Data:    nil,
				Error:   "Amount cannot be zero.",
			},
		},
	}

	t.Run("Credit", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				h := getMockhandler()
				req := httptest.NewRequest(http.MethodPost, "/credit", bytes.NewBuffer([]byte(fmt.Sprintf(`{"amount": %f}`, tt.amount))))
				req.Header.Set("Content-Type", "application/json")
				w := httptest.NewRecorder()
				h.Credit(w, req)
				resp := w.Result()
				var res response.RespBody
				json.NewDecoder(resp.Body).Decode(&res)
				if res.Success != tt.res.Success {
					t.Errorf("Expected success: %v, got: %v", tt.res.Success, res.Success)
				}
				if res.Data != nil {
					if res.Data.(map[string]interface{})["balance"] != tt.res.Data.(service.TxnResponse).Balance {
						t.Errorf("Expected balance: %v, got: %v", tt.res.Data.(service.TxnResponse).Balance, res.Data.(map[string]interface{})["balance"])
					}
				}
				if res.Error != tt.res.Error {
					t.Errorf("Expected error: %v, got: %v", tt.res.Error, res.Error)
				}
			})
		}
	})
}

func TestDebit(t *testing.T) {
	tests := []struct {
		name   string
		amount float64
		res    response.RespBody
	}{
		{
			name:   "When debit amount is greater than 0",
			amount: 10,
			res: response.RespBody{
				Success: true,
				Data:    service.TxnResponse{Balance: 90},
				Error:   "",
			},
		},
		{
			name:   "When debit amount is less than 0",
			amount: -10,
			res: response.RespBody{
				Success: true,
				Data:    service.TxnResponse{Balance: 90},
				Error:   "",
			},
		},
		{
			name:   "When credit amount is equal to 0",
			amount: 0,
			res: response.RespBody{
				Success: false,
				Data:    nil,
				Error:   "Amount cannot be zero.",
			},
		},
	}

	t.Run("Debit", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				h := getMockhandler()
				req := httptest.NewRequest(http.MethodPost, "/debit", bytes.NewBuffer([]byte(fmt.Sprintf(`{"amount": %f}`, tt.amount))))
				req.Header.Set("Content-Type", "application/json")
				w := httptest.NewRecorder()
				h.Debit(w, req)
				resp := w.Result()
				var res response.RespBody
				json.NewDecoder(resp.Body).Decode(&res)
				if res.Success != tt.res.Success {
					t.Errorf("Expected success: %v, got: %v", tt.res.Success, res.Success)
				}
				// fmt.Printf("%+v", res)
				if res.Data != nil {
					if res.Data.(map[string]interface{})["balance"] != tt.res.Data.(service.TxnResponse).Balance {
						t.Errorf("Expected balance: %v, got: %v", res.Data.(map[string]interface{})["balance"], res.Data.(map[string]interface{})["balance"])
					}
				}
				if res.Error != tt.res.Error {
					t.Errorf("Expected error: %v, got: %v", tt.res.Error, res.Error)
				}
			})
		}
	})
}

func TestBalance(t *testing.T) {
	tests := []struct {
		name string
		res  response.RespBody
	}{
		{
			name: "Get balance",
			res: response.RespBody{
				Success: true,
				Data:    service.TxnResponse{Balance: 100},
				Error:   "",
			},
		},
	}

	t.Run("GetBalance", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				h := getMockhandler()
				req := httptest.NewRequest(http.MethodGet, "/debit", nil)
				req.Header.Set("Content-Type", "application/json")
				w := httptest.NewRecorder()
				h.Balance(w, req)
				resp := w.Result()
				var res response.RespBody
				json.NewDecoder(resp.Body).Decode(&res)
				if res.Success != tt.res.Success {
					t.Errorf("Expected success: %v, got: %v", tt.res.Success, res.Success)
				}
				if res.Data != nil {
					if res.Data.(map[string]interface{})["balance"] != tt.res.Data.(service.TxnResponse).Balance {
						t.Errorf("Expected balance: %v, got: %v", res.Data.(map[string]interface{})["balance"], res.Data.(map[string]interface{})["balance"])
					}
				}
				if res.Error != tt.res.Error {
					t.Errorf("Expected error: %v, got: %v", tt.res.Error, res.Error)
				}
			})
		}
	})
}

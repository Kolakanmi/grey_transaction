package service

import (
	"context"
	"log"
	"testing"

	"github.com/Kolakanmi/grey_transaction/pkg/apperror"
	proto "github.com/Kolakanmi/grey_transaction/pkg/grpc/transaction"
	"github.com/Kolakanmi/grey_transaction/repository/mock"
	mockConn "github.com/Kolakanmi/grey_transaction/service/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getService() *TransactionService {
	repo := mock.NewMockRepo()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(mockConn.Dialer()))
	if err != nil {
		log.Fatal(err)
	}
	wc := proto.NewWalletClient(conn)

	// client := mockConn.NewMockClient()
	return NewTransactionService(repo, wc)
}

func TestCredit(t *testing.T) {
	tests := []struct {
		name   string
		amount float64
		res    *TxnResponse
		err    error
	}{
		{
			name:   "When credit amount is greater than 0",
			amount: 10,
			res:    &TxnResponse{Balance: 110},
			err:    nil,
		},
		{
			name:   "When credit amount is less than 0",
			amount: -10,
			res:    nil,
			err:    apperror.BadRequestError("Amount cannot be negative."),
		},
		{
			name:   "When credit amount is equal to 0",
			amount: 0,
			res:    nil,
			err:    apperror.BadRequestError("Amount cannot be zero."),
		},
	}

	t.Run("Credit", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := getService()
				got, err := s.Credit(context.Background(), tt.amount)
				if err != nil {
					if tt.err == nil {
						t.Errorf("Expected no error, but got %v", err)
					}
					if err.Error() != tt.err.Error() {
						t.Errorf("Expected error to be %v, but got %v", tt.err, err)
					}
				}
				if tt.res != nil && got.Balance != tt.res.Balance {
					t.Errorf("Expected balance to be %v, but got %v", tt.res.Balance, got.Balance)
				}
			})
		}
	})
}

func TestDebit(t *testing.T) {
	tests := []struct {
		name   string
		amount float64
		res    *TxnResponse
		err    error
	}{
		{
			name:   "When debit amount is greater than 0, convert amount to negative",
			amount: 10,
			res:    &TxnResponse{Balance: 90},
			err:    nil,
		},
		{
			name:   "When debit amount is less than 0",
			amount: -10,
			res:    &TxnResponse{Balance: 90},
			err:    nil,
		},
		{
			name:   "When debit amount is equal to 0",
			amount: 0,
			res:    nil,
			err:    apperror.BadRequestError("Amount cannot be zero."),
		},
	}
	t.Run("Debit", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := getService()
				got, err := s.Debit(context.Background(), tt.amount)
				if err != nil {
					if tt.err == nil {
						t.Errorf("Expected no error, but got %v", err)
					}
					if err.Error() != tt.err.Error() {
						t.Errorf("Expected error to be %v, but got %v", tt.err, err)
					}
				}
				if tt.res != nil && got.Balance != tt.res.Balance {
					t.Errorf("Expected balance to be %v, but got %v", tt.res.Balance, got.Balance)
				}
			})
		}
	})
}

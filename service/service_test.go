package service

import (
	"context"
	"testing"

	"github.com/Kolakanmi/grey_transaction/repository/mock"
	mockConn "github.com/Kolakanmi/grey_transaction/service/mock"
)

func getService() *TransactionService {
	repo := mock.NewMockRepo()
	client := mockConn.NewMockClient()
	return NewTransactionService(repo, client)
}

func TestCredit(t *testing.T) {
	t.Run("Credit", func(t *testing.T) {
		t.Run("Should return the balance after credit", func(t *testing.T) {
			t.Run("When credit amount is greater than 0", func(t *testing.T) {
				s := getService()
				got, err := s.Credit(context.Background(), 10)
				if err != nil {
					t.Errorf("Expected no error, but got %v", err)
				}
				if got.Balance != 110 {
					t.Errorf("Expected balance to be %v, but got %v", 110, got.Balance)
				}
			})
			t.Run("When credit amount is less than 0", func(t *testing.T) {
				s := getService()
				got, err := s.Credit(context.Background(), -10)
				if err == nil {
					t.Errorf("Expected error, but got %+v", got)
				}
			})
		})
	})
}

func TestDebit(t *testing.T) {
	t.Run("Debit", func(t *testing.T) {
		t.Run("Should return the balance after debit", func(t *testing.T) {
			t.Run("When debit amount is greater than 0, convert amount to negative", func(t *testing.T) {
				s := getService()
				got, err := s.Debit(context.Background(), 10)
				if err != nil {
					t.Errorf("Expected no error, but got %v", err)
				}
				if got.Balance != 90 {
					t.Errorf("Expected balance to be %v, but got %v", 90, got.Balance)
				}
			})
			t.Run("When debit amount is less than 0", func(t *testing.T) {
				s := getService()
				got, err := s.Debit(context.Background(), -10)
				if err != nil {
					t.Errorf("Expected no error, but got %v", err)
				}
				if got.Balance != 80 {
					t.Errorf("Expected balance to be %v, but got %v", 80, got.Balance)
				}
			})
		})
	})
}

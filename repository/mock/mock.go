package mock

import (
	"context"
	"fmt"
	"time"

	"github.com/Kolakanmi/grey_transaction/model"
)

type MockRepo struct{}

func NewMockRepo() *MockRepo {
	return &MockRepo{}
}

var timeNow, _ = time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
var txns = []model.Transaction{
	{
		Base: model.Base{
			ID:        "1",
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
		Amount: 100,
		Status: model.Success,
	},
	{
		Base: model.Base{
			ID:        "2",
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
		Amount: -50,
		Status: model.Success,
	},
}

func (m *MockRepo) Create(ctx context.Context, txn *model.Transaction) (string, error) {
	txn.ID = fmt.Sprintf("%d", len(txns)+1)
	txn.CreatedAt = timeNow
	txn.UpdatedAt = timeNow
	txns = append(txns, *txn)
	return txn.ID, nil
}

func (m *MockRepo) GetAll(ctx context.Context) ([]model.Transaction, error) {
	return txns, nil
}

func (m *MockRepo) UpdateStatus(ctx context.Context, txnID string, status string) error {
	for i, txn := range txns {
		if txn.ID == txnID {
			txns[i].Status = status
			return nil
		}
	}
	return nil
}

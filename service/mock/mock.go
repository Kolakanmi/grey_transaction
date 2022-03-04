package mock

import (
	"context"

	proto "github.com/Kolakanmi/grey_transaction/pkg/grpc/transaction"
	"google.golang.org/grpc"
)

type MockClient struct{}

func NewMockClient() proto.WalletClient {
	return &MockClient{}
}

var balance = 100.0

func (m *MockClient) GetBalance(ctx context.Context, in *proto.GetBalanceRequest, opts ...grpc.CallOption) (*proto.GetBalanceResponse, error) {
	return &proto.GetBalanceResponse{Balance: balance}, nil
}
func (m *MockClient) UpdateBalance(ctx context.Context, in *proto.UpdateBalanceRequest, opts ...grpc.CallOption) (*proto.UpdateBalanceResponse, error) {
	return &proto.UpdateBalanceResponse{Balance: balance + in.Amount}, nil
}

package mock

import (
	"context"
	"log"
	"net"

	proto "github.com/Kolakanmi/grey_transaction/pkg/grpc/transaction"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var balance = 100.0

type MockServer struct {
	proto.UnimplementedWalletServer
}

func (s *MockServer) GetBalance(ctx context.Context, req *proto.GetBalanceRequest) (*proto.GetBalanceResponse, error) {
	return &proto.GetBalanceResponse{Balance: balance}, nil
}
func (s *MockServer) UpdateBalance(ctx context.Context, req *proto.UpdateBalanceRequest) (*proto.UpdateBalanceResponse, error) {
	return &proto.UpdateBalanceResponse{Balance: balance + req.Amount}, nil
}

func Dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	proto.RegisterWalletServer(server, &MockServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

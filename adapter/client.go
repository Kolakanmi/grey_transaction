package adapter

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Kolakanmi/grey_transaction/pkg/envconfig"
	proto "github.com/Kolakanmi/grey_transaction/pkg/grpc/transaction"
)

type Config struct {
	Address string `envconfig:"GRPC_ADDRESS" default:"localhost"`
	Port    int    `envconfig:"GRPC_PORT" default:"9090"`
}

func LoadConfig() *Config {
	var conf Config
	envconfig.Load(&conf)
	return &conf
}

func NewClientConnection(config *Config) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.Address, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewClient(conn *grpc.ClientConn) proto.WalletClient {
	wc := proto.NewWalletClient(conn)
	return wc
}

func CloseConnection(conn *grpc.ClientConn) {
	conn.Close()
}

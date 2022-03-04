package api

import (
	"log"
	"net/http"

	"github.com/Kolakanmi/grey_transaction/adapter"
	"github.com/Kolakanmi/grey_transaction/handler"
	"github.com/Kolakanmi/grey_transaction/pkg/database"
	"github.com/Kolakanmi/grey_transaction/pkg/http/middleware"
	"github.com/Kolakanmi/grey_transaction/pkg/http/router"
	"github.com/Kolakanmi/grey_transaction/repository"
	"github.com/Kolakanmi/grey_transaction/service"
)

func NewRouter() (http.Handler, error) {
	dbConf := database.LoadConfig()
	db, err := database.ConnectDB(dbConf)
	if err != nil {
		log.Printf("db error %v", err)
		return nil, err
	}

	repo := repository.NewTransactionRepository(db)

	grpcConf := adapter.LoadConfig()
	conn, err := adapter.NewClientConnection(grpcConf)
	if err != nil {
		log.Printf("grpc client connection error %v", err)
		return nil, err
	}
	walletClient := adapter.NewClient(conn)

	service := service.NewTransactionService(repo, walletClient)

	handler := handler.New(service)

	routes := []router.Route{}
	routes = append(routes, handler.Routes()...)

	rConf := router.GetEmptyConfig()
	rConf.Routes = routes
	rConf.GlobalMiddlewares = []router.Middleware{
		middleware.Recover,
	}

	rConf.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	r, err := router.New(rConf)
	if err != nil {
		log.Printf("router err: %v", err)
		return nil, err
	}
	return middleware.CORS(r), nil
}

package api

import (
	"log"
	"net/http"

	"github.com/Kolakanmi/grey_transaction/adapter"
	"github.com/Kolakanmi/grey_transaction/handler"
	"github.com/Kolakanmi/grey_transaction/pkg/database"
	ch "github.com/Kolakanmi/grey_transaction/pkg/http/handler"
	"github.com/Kolakanmi/grey_transaction/pkg/http/middleware"
	"github.com/Kolakanmi/grey_transaction/pkg/http/response"
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

	txnService := service.NewTransactionService(repo, walletClient)

	handler := handler.New(txnService)

	routes := []router.Route{
		{
			Path:   "/readiness",
			Method: http.MethodGet,
			Handler: ch.CustomHandler(func(rw http.ResponseWriter, r *http.Request) error {
				return response.OK("Server is up!!!", nil).ToJSON(rw)
			}),
		},
	}
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
	log.Println("Router created")
	return middleware.CORS(r), nil
}

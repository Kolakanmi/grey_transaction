package main

import (
	"log"

	"github.com/Kolakanmi/grey_transaction/api"
	"github.com/Kolakanmi/grey_transaction/pkg/envconfig"
	"github.com/Kolakanmi/grey_transaction/pkg/http/server"

	_ "github.com/Kolakanmi/grey_transaction/docs"
)

// @title Transaction Service
// @version 1.0
// @description Transaction Endpoints.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {
	err := envconfig.SetEnvFromConfig(".env")
	if err != nil {
		log.Println("env config load err: ", err)
	}

	r, err := api.NewRouter()
	if err != nil {
		log.Printf("error : %v", err)
	}
	serverConf := server.LoadConfigFromEnv()
	server.ListenAndServe(*serverConf, r)
}

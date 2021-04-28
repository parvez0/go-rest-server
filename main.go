package main

import (
	"devops_assignment/pkg"
	"fmt"
	"net/http"
)

var cfg *pkg.Config

// init is used to initialize the singleton objects
// which will be used throughout the project
func init()  {
	cfg = pkg.InitializeConfig()
	_, err := pkg.InitializeDb()
	if err != nil {
		panic("failed to initialize db connection : " + err.Error())
	}
}

func main() {

	// defining request handlers for handling RESTApi calls by default,
	// it will listen for all methods available
	http.HandleFunc("/health-check", pkg.HandlerHealthCheck)
	http.HandleFunc("/invalid-deliveries", pkg.HandleGetInvalidDeliveries)

	// starting the server with default handler and using the server host and port
	// define in the config file. if any error encountered during the initialization
	// of the server, it will exit with panic and encountered error
	fmt.Printf("starting the server on port : %s\n", cfg.Server.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port), nil); err != nil {
		panic("server shutdown - unexpected error occurred : " + err.Error())
	}
}

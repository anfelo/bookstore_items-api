package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/anfelo/bookstore_items-api.git/api"
	"github.com/anfelo/bookstore_items-api.git/items"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

// StartApplication starts the web server
func StartApplication() {
	service := items.NewItemsService()
	handler := api.NewHandler(service)

	router.HandleFunc("/items", handler.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/:id", handler.GetByID).Methods(http.MethodGet)

	// Ping method for cloud provider
	router.HandleFunc("/ping", handler.Ping).Methods(http.MethodGet)

	fmt.Println("Listening on port :8000")
	if err := http.ListenAndServe(httpPort(), router); err != nil {
		panic(err)
	}
}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

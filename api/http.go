package api

import (
	"fmt"
	"net/http"

	"github.com/anfelo/bookstore_items-api.git/items"
)

// ItemsHandler interface
type ItemsHandler interface {
	GetByID(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)

	Ping(http.ResponseWriter, *http.Request)
}

type handler struct {
	itemsService items.ItemsService
}

// NewHandler method incharge of creating a new items handler
func NewHandler(itemsService items.ItemsService) ItemsHandler {
	return &handler{
		itemsService,
	}
}

func (h *handler) GetByID(w http.ResponseWriter, r *http.Request) {
	item, err := h.itemsService.GetByID("")
	if err != nil {
		// TODO: Write to the response the err
		return
	}

	// TODO: Write to the response the item JSON and Status: 200
	fmt.Println(item)
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	item := items.Item{}
	result, err := h.itemsService.Create(item)
	if err != nil {
		// TODO: Write to the response the err
		return
	}

	// TODO: Write to the response the result JSON and Status: 201
	fmt.Println(result)
}

func (h *handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

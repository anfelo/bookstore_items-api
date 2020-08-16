package api

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/anfelo/bookstore_items-api.git/items"
	js "github.com/anfelo/bookstore_items-api.git/serializer/json"
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

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) GetByID(w http.ResponseWriter, r *http.Request) {
	itemID := r.URL.Query().Get("id")
	item, err := h.itemsService.GetByID(itemID)
	if err != nil {
		// TODO: Respond with json error
		return
	}

	// TODO: Write to the response the item JSON and Status: 200
	itemSerializer := &js.ItemSerializer{}
	responseBody, err := itemSerializer.Encode(item)
	if err != nil {
		// TODO: Respond with json error
		return
	}
	setupResponse(w, "application/json", responseBody, http.StatusOK)
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO: Write to the response the err
		return
	}
	itemSerializer := &js.ItemSerializer{}
	item, restErr := itemSerializer.Decode(requestBody)
	if restErr != nil {
		// TODO: Write to the response the err
		return
	}
	result, restErr := h.itemsService.Create(item)
	if restErr != nil {
		// TODO: Write to the response the err
		return
	}

	// TODO: Write to the response the result JSON and Status: 201
	responseBody, restErr := itemSerializer.Encode(result)
	if restErr != nil {
		// TODO: Respond with json error
		return
	}
	setupResponse(w, "application/json", responseBody, http.StatusCreated)
}

func (h *handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

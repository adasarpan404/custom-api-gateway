package routes

import (
	"net/http"

	"github.com/adasarpan404/custom-api-gateway/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/service1", handlers.Service1Handler)
	mux.HandleFunc("/service2", handlers.Service2Handler)
}

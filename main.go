package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adasarpan404/custom-api-gateway/middleware"
	"github.com/adasarpan404/custom-api-gateway/routes"
)

func main() {
	port := 8000

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	handler := middleware.Logger(mux)
	handler = middleware.Auth(handler)

	fmt.Printf("API Gateway running on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))

}

package main

import (
	"fmt"
	"gamma/src/utils"
	"log"
	"net/http"

	health_controller "gamma/src/controller/health"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// Router
	router.GET("/health", health_controller.GetHealth)
	// Start
	var PORT string = utils.Getenv("PORT", "8080")
	log.Printf("🚀 Server is listening on port %s", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))

}

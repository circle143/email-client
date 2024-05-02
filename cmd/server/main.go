package main

import (
	"fmt"
	"net/http"
	"os"

	"mennr.tech/api/router"
)

func main() {
	PORT := "8080"
	if port := os.Getenv("PORT"); port != "" {
		PORT = port
	}

	router := router.Router()

	fmt.Println("server is running on port:", PORT)
	http.ListenAndServe(":"+PORT, router)

}

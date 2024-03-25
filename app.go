package main

import (
	"apihub_backend/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Starting server...")
	router := routes.Routes()
	// c := cors.New(cors.Options{
	// 	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	// })

	// handler := c.Handler(router)
	// http.ListenAndServe(":"+port, middlewares.LogRequest(handler))
	fmt.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}

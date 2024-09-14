// The WGS84 coordinates for Amrita Vishwa Vidyapeetham in Coimbatore, India, are approximately:
// lat: 11.0587° N
// lon: 76.9434° E
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	mux := http.NewServeMux()
	// [METHOD] /path, the space b/w method and / is important
	// Creating Path for Basic Package and Wind Package
	mux.HandleFunc("GET /basic", handlerUser)
	mux.HandleFunc("GET /wind", handlerUser)
	log.Println("Server starting at port 8080")

	// If u don't pass in your ServeMux, it will use the default ServeMux
	http.ListenAndServe(":8080", mux)

}

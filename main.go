package main

import (
	"fmt"
	"net/http"
	"smartpi/handlers"
)

const portNumber = ":8080"
const staticHtml = "static/html/"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/static/", handlers.StaticHandler)
	fmt.Println(fmt.Sprintf("Starting application on http://localhost%s/", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

package main

import (
	"fmt"
	"net/http"
)

func serveStaticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", serveStaticHandler)
	fmt.Println("Server listening on 0.0.0.0:8000...")
	http.ListenAndServe(":8000", nil)
}

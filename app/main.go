package main

import (
	"fmt"
	"net/http"
)

func welcomHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/text")
	res.Write([]byte("Welcome to Binary World!"))
}

func main() {
	fmt.Println("Welcome to Binary World !")
	http.HandleFunc("/", welcomHandler)
	http.ListenAndServe(":8080", nil)
}

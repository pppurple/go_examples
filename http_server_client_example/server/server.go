package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", handleGetHello)
	http.ListenAndServe(":8080", nil)
}

func handleGetHello(w http.ResponseWriter, req *http.Request) {
	// header
	method := req.Method
	fmt.Println("[method] " + method)
	for k, v := range req.Header {
		fmt.Print("[header] " + k)
		fmt.Println(": " + strings.Join(v, ","))
	}

	// GET
	if method == "GET" {
		req.ParseForm()
		for k, v := range req.Form {
			fmt.Print("[param] " + k)
			fmt.Println(": " + strings.Join(v, ","))
		}
		fmt.Fprint(w, "Recieve Get request!!")
	}

	// POST
	if method == "POST" {

	}
}

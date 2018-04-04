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
	req.ParseForm()
	params := ""
	for k, v := range req.Form {
		params += (k + "=" + strings.Join(v, ","))
		params += ", "
	}
	fmt.Fprintf(w, "get request. params : %s", params)
}

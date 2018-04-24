package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/dog", handleDog)
	http.HandleFunc("/dog_json", handleDogJson)
	http.ListenAndServe(":8080", nil)
}

func handleDog(w http.ResponseWriter, req *http.Request) {
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
		fmt.Fprint(w, "Recieved Get request!!")
	}

	// POST (form)
	if method == "POST" {
		defer req.Body.Close()
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[request body row] " + string(body))
		decoded, error := url.QueryUnescape(string(body))
		if error != nil {
			log.Fatal(error)
		}
		fmt.Println("[request body decoded] ", decoded)
		fmt.Fprint(w, "Recieved Post(form) request!!")
	}
}

type Dog struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func handleDogJson(w http.ResponseWriter, req *http.Request) {
	// header
	method := req.Method
	fmt.Println("[method] " + method)
	for k, v := range req.Header {
		fmt.Print("[header] " + k)
		fmt.Println(": " + strings.Join(v, ","))
	}

	// POST (json)
	if method == "POST" {
		defer req.Body.Close()
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("[request body row] " + string(body))

		// Unmarshal
		var dog Dog
		error := json.Unmarshal(body, &dog)
		if error != nil {
			log.Fatal(error)
		}
		fmt.Printf("[request body decoded] %+v\n", dog)
		fmt.Fprint(w, "Recieved Post(json) request!!")
	}
}

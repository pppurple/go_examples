package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	get()
	postAsForm()
	postAsJson()
}

func get() {
	// Get request
	res, err := http.Get("http://localhost:8080/dog?id=123&order=desc")
	if err != nil {
		log.Fatal(err)
	}

	// header
	fmt.Printf("[status] %d\n", res.StatusCode)
	for k, v := range res.Header {
		fmt.Print("[header] " + k)
		fmt.Println(": " + strings.Join(v, ","))
	}

	// body
	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("[body] " + string(body))
}

func postAsForm() {
	// form values
	values := url.Values{}
	values.Add("id", "123")
	values.Add("name", "ポメラニアン")
	values.Encode()

	res, err := http.PostForm("http://localhost:8080/dog", values)
	if err != nil {
		log.Fatal(err)
	}

	// header
	fmt.Printf("[status] %d\n", res.StatusCode)
	for k, v := range res.Header {
		fmt.Print("[header] " + k)
		fmt.Println(": " + strings.Join(v, ","))
	}

	// body
	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("[body] " + string(body))
}

type Dog struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func postAsJson() {
	// json values
	values, err := json.Marshal(Dog{Id: 2, Name: "柴犬"})

	res, err := http.Post("http://localhost:8080/dog_json", "application/json", bytes.NewBuffer(values))
	if err != nil {
		log.Fatal(err)
	}

	// header
	fmt.Printf("[status] %d\n", res.StatusCode)
	for k, v := range res.Header {
		fmt.Print("[header] " + k)
		fmt.Println(": " + strings.Join(v, ","))
	}

	// body
	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("[body] " + string(body))
}

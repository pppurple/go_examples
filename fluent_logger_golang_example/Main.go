package main

import (
	"fmt"

	"github.com/fluent/fluent-logger-golang/fluent"
)

func main() {
	logger, err := fluent.New(fluent.Config{FluentHost: "localhost", FluentPort: 24224})
	if err != nil {
		fmt.Println(err)
	}
	defer logger.Close()

	tag := "app.user"
	var data = map[string]string{
		"name":    "Bobby",
		"age":     "34",
		"country": "Japan",
	}

	error := logger.Post(tag, data)

	if error != nil {
		panic(error)
	}
}

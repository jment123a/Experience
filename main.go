package main

import (
	"Experience/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handler.Ping)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("err123")
	}

}

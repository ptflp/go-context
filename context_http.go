package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		fmt.Println("Start request")

		select {
		case <-time.After(2 * time.Second):
			writer.Write([]byte("request processed"))
			fmt.Println("request end")
		case <-ctx.Done():
			fmt.Println("request canceled")
		}
	}))
}

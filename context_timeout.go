package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	client := &http.Client{}

	ctx := context.Background()

	min := 50
	max := 150

	i := 0
	for i < 21 {
		i++

		ctr, _ := context.WithTimeout(ctx, time.Duration(rand.Intn(max-min)+min)*time.Millisecond)

		req, err := http.NewRequest(http.MethodGet, "http://google.com", nil)

		if err != nil {
			panic(err)
		}

		req = req.WithContext(ctr)
		res, err := client.Do(req)

		if err != nil {
			fmt.Println("Request failed:", err)
			continue
		}

		// Print the statuscode if the request succeeds
		fmt.Println("Response received, status code:", res.StatusCode)
	}
}

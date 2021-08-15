package main

import (
	"fmt"
	"net/http"

	"golang.org/x/time/rate"
)

func main() {
	app := http.NewServeMux()
	app.HandlerFunc("/", requestOK)
	http.ListenAndServe(":8080", limits(app))
}

var limits = rate.NewLimiter(10, 10)

func limits(handle Http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if limits.Allow() != true {
			fmt.Println("over limits")
			return
		}
		handle.ServeHTTP(w, req)
	})
}

func requestOK(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("success"))
}

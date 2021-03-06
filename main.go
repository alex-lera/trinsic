package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println(port)
	router := mux.NewRouter()
	router.HandleFunc("/", test)
	router.NotFoundHandler = http.HandlerFunc(test2)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}
}

func test(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello world"))
	fmt.Println("remote address:")
	fmt.Println(req.RemoteAddr)
	fmt.Println("forwarded for:")
	fmt.Println(req.Header.Get("X-Forwarded-For"))
}

func test2(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello world not found"))
	fmt.Println("remote address:")
	fmt.Println(req.RemoteAddr)
	fmt.Println("forwarded for:")
	fmt.Println(req.Header.Get("X-Forwarded-For"))
}

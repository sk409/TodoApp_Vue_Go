package main

import (
	"fmt"
	"net/http"
)

func f(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	fmt.Println("AAAAA!!!!")
	fmt.Fprintln(w, "BBBB")
}

func main() {
	http.HandleFunc("/todos", handleRequestTodo)
	//http.HandleFunc("/", f)
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	server.ListenAndServe()
}

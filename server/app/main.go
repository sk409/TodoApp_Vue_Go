package main

import (
	"net/http"
)

// func f(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Access-Control-Allow-Origin", "*")
// 	fmt.Println("AAAAA!!!!")
// 	fmt.Fprintln(w, "BBBB")
// }

func main() {
	http.HandleFunc("/todos", handleRequestTodo)
	//http.HandleFunc("/", f)
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	server.ListenAndServe()
	// fmt.Println("ABC")
}

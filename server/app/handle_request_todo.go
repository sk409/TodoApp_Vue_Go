package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func handleRequestTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var err error
	switch r.URL.Path {
	case "/todos":
		switch r.Method {
		case "POST":
			err = postTodo(w, r)
		}
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func postTodo(w http.ResponseWriter, r *http.Request) error {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var todo Todo
	err := json.Unmarshal(body, &todo)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return errors.New("Bad Body")
	}
	w.Write([]byte(todo.Content + todo.Title))
	Db.Create(&todo)
	return nil
}

// func handleFetch(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "RECEIVED RESPONSE")
// }

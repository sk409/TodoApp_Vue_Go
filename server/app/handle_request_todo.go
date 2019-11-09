package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func handleRequestTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OKOKOKOKOKOOKO")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	var err error
	if r.URL.Path == "/todos" {
		switch r.Method {
		case "GET":
			err = fetchTodo(w, r)
		case "POST":
			err = storeTodo(w, r)
		}
	}
	// else if regexp.MustCompile("/todos/[0-9]").Match([]byte(r.URL.Path)) {
	// 	switch r.Method {
	// 	case "PUT":
	// 		err = updateTodo(w, r)
	// 	}
	// }
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fetchTodo(w http.ResponseWriter, r *http.Request) error {
	for key, value := range r.URL.Query() {
		Db.Where(key+"=$1", value[0])
	}
	var todos []Todo
	Db.Find(&todos)
	j, err := json.Marshal(&todos)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
	return nil
}

func storeTodo(w http.ResponseWriter, r *http.Request) error {
	todo, err := makeTodoFromRequestedBody(r)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return errors.New("Bad Body")
	}
	w.Write([]byte(todo.Content + todo.Title))
	Db.Create(&todo)
	return nil
}

func updateTodo(w http.ResponseWriter, r *http.Request) error {
	todo, err := makeTodoFromRequestedBody(r)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return errors.New("Bad Body")
	}
	Db.Save(&todo)
	return nil
}

func makeTodoFromRequestedBody(r *http.Request) (Todo, error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var todo Todo
	err := json.Unmarshal(body, &todo)
	return todo, err
}

// func handleFetch(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "RECEIVED RESPONSE")
// }

package main

type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Limit   string `json:"limit"`
	// User    User
}

package main

import (
	"log"
	"net/http"

	"github.com/yousefzinsazk78/test_web_app_0.2/web/handler"
)

func main() {

	fs := http.FileServer(http.Dir("static/assets/"))
	http.Handle("/static/assets/", http.StripPrefix("/static/assets/", fs))
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/save", handler.SaveHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

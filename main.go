package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/bmizerany/pat"
	"github.com/t-fukui/alpaca/config"
)

var db gorm.DB

func main() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(IndexHandler))
	mux.Get("/new", http.HandlerFunc(NewHandler))
	mux.Post("/create", http.HandlerFunc(CreateHandler))
	mux.Get("/edit/:id", http.HandlerFunc(EditHandler))
	mux.Post("/update/:id", http.HandlerFunc(UpdateHandler))

	http.Handle("/", mux)
	// Webサーバーを起動
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe:", Log(http.DefaultServeMux))
	}
}

func init() {
	db = config.Database()
}

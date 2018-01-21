package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"html"
	//"log"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func SetDatabase(db *gorm.DB) {
    DB = db
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func ArticlesListHandler(w http.ResponseWriter, r *http.Request) {
	articles := Articles {
		Article {
			Title: "石展丞又遲到了",
			Description: "時間控制者-石展丞",
		},
		Article {
			Title: "尾牙公告",
			Description: "尾牙在 2/14 舉行，請大家務必記得", // must have
		},
	}

	json.NewEncoder(w).Encode(articles)
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
	// read
	article := Article{}
	DB.First(&article, vars["id"])

	json.NewEncoder(w).Encode(article)
}

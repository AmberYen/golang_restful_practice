package main

import (
	"fmt" // printf
  "html"
	"log"
	"net/http" // Package http provides HTTP client and server implementations
  "github.com/gorilla/mux"
)

func main() {
  router := mux.NewRouter()
	router.HandleFunc("/", fooHandler)
  router.HandleFunc("/articles/{id}", ArticlesCategoryHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "Article id: %v\n", vars["id"])
}

// s := &http.Server{
// 	Addr:           ":8080",
// 	Handler:        myHandler,
// 	ReadTimeout:    10 * time.Second,
// 	WriteTimeout:   10 * time.Second,
// 	MaxHeaderBytes: 1 << 20,
// }
// log.Fatal(s.ListenAndServe())

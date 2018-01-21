package main

import (
	"fmt" // printf
  "html"
	"log"
	"net/http" // Package http provides HTTP client and server implementations
)

func main() {
  http.HandleFunc("/foo", fooHandler)

  http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
  	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, foo")
}

// s := &http.Server{
// 	Addr:           ":8080",
// 	Handler:        myHandler,
// 	ReadTimeout:    10 * time.Second,
// 	WriteTimeout:   10 * time.Second,
// 	MaxHeaderBytes: 1 << 20,
// }
// log.Fatal(s.ListenAndServe())

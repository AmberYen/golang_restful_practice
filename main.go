package main

import (
	"log"
	"net/http" // Package http provides HTTP client and server implementations
)

func main() {
	// db
	db := ConnectDB()
	SetDatabase(db)

	// init mocks
  // db.Create(&Article{Title: "Hello", Description: "HI", Status: true})

	// router
  router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}


// s := &http.Server{
// 	Addr:           ":8080",
// 	Handler:        myHandler,
// 	ReadTimeout:    10 * time.Second,
// 	WriteTimeout:   10 * time.Second,
// 	MaxHeaderBytes: 1 << 20,
// }
// log.Fatal(s.ListenAndServe())

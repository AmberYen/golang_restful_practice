package main

import "time"

type Article struct {
	Title       string
	Description string
  Status      bool
	CreatedAt         time.Time 
}

type Articles []Article

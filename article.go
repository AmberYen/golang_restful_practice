package main

import "time"

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
  Status      bool   `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Articles []Article

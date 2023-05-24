package model

import ()

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	Content   string `json:"text"`
	TimeStamp string `json:"timestamp"`
}

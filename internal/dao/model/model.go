package model

import "gorm.io/gorm"

type User struct {
	Name string `json:"name"`
	gorm.Model
}

type Message struct {
	Content string `json:"content"`
	gorm.Model
}

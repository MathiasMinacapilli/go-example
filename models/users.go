package models

import "go-example/db"

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

/**
  To initialize
*/
func CreateModel() {
    db.Database.AutoMigrate(User{})
}
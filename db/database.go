package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:root@tcp(localhost:3306)/go-finances?chatset=utf8mb4&parseTimetrue&loc=Local"
var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la conn", err)
		// panic(err)
		return nil
	} else {
		fmt.Println("Conn ok")
		return db
	}
}()
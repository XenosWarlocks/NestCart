package main

import (
	"log"

	"github.com/XenosWarlocks/NestCart/cmd/api"
	"github.com/XenosWarlocks/NestCart/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 "roots",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "XXXXXXXX",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	db, err := db.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal()
	}
}

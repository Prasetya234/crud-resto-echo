package db

import (
	"crud_product/config"
	"database/sql"
	fmt "fmt"
	_ "github.com/lib/pq"
)

func NewInstanceDB(conf config.Configuration) *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.DBHost, conf.DBPort, conf.DBUsername, conf.DBPassword, conf.DBName))
	if err != nil {
		panic(err)
	}
	err = db.Ping() // untuk mengecek apakah database tersedia atau tidak
	if err != nil {
		panic(err)
	}
	return db
}

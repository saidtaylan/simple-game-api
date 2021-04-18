package main

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/lib/pq"
)

var (
	host     = "localhost" //os.Getenv("HOST")
	user     = "postgres"  //os.Getenv("USER")
	password = "said"      //os.Getenv("PASSWORD")
	dbname   = "gamesapi"  //os.Getenv("DBNAME")
	port     = "5432"      //os.Getenv("PORT")
	sslmode  = "disable"   //os.Getenv("SSLMODE")
	//timezone = "Europe/Istanbul" // os.Getenv("TIMEZONE")
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
)

var db *sql.DB
var err error

func connectToDB() {
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln("Could not connect to database")
	}
	log.Println("Connected to database")
	_, err := db.Query("SELECT * FROM games")
	if err != nil {
		db.Exec("Create Table games (id serial NOT NULL,name varchar(100),latest_version varchar(10),category varchar(20),descr varchar(500),producer varchar(50), PRIMARY KEY(id));")
	}
	//db.Exec("CREATE TABLE IF NOT EXISTS games(id int PRIMARY KEY, name )")
}

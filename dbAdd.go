package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type City struct {
	ID          int    `json:"id,omitempty"  db:"ID"`
	Name        string `json:"name,omitempty"  db:"Name"`
	CountryCode string `json:"countryCode,omitempty"  db:"CountryCode"`
	District    string `json:"district,omitempty"  db:"District"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}

func main() {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}

	v := os.Args
	fmt.Println("Connected!")
	//var city City
	//if err := db.Get(&city, "INPUT * FROM city WHERE Name='"+v[1]+"'"); errors.Is(err, sql.ErrNoRows) {
	fmt.Println("insert into city (Name, CountryCode, District, Population) VALUES(?,?,?,?)", v[1], v[2], v[3], v[4])
	db.Exec("insert into city (Name, CountryCode, District, Population) VALUES(?,?,?,?)", v[1], v[2], v[3], v[4])
	//		log.Printf("no such city Name = %s", v[1])
	//	} else if err != nil {
	//		log.Fatalf("DB Error: %s", err)
	//	}
	//
	//fmt.Printf(v[1]+"の人口は%d人です\n", city.Population)
}

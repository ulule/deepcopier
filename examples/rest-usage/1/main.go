package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

type Account struct {
	ID         uint `gorm:"primary_key"`
	FirstName  string
	LastName   string
	Username   string
	Password   string
	Email      string
	DateJoined time.Time
}

type Accounts struct {
	Db gorm.DB
}

func (a *Accounts) Detail(w rest.ResponseWriter, r *rest.Request) {
	account := &Account{}
	result := a.Db.First(&account, "username = ?", r.PathParam("username"))

	if result.RecordNotFound() {
		rest.NotFound(w, r)
		return
	}

	w.WriteJson(&account)
}

func main() {
	dsn := fmt.Sprintf("user=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_NAME"))

	db, err := gorm.Open("postgres", dsn)

	fmt.Println(dsn)

	if err != nil {
		panic(err)
	}

	db.DB()
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true)
	db.LogMode(true)

	api := rest.NewApi()

	api.Use(rest.DefaultDevStack...)

	accounts := &Accounts{Db: db}

	router, err := rest.MakeRouter(
		rest.Get("/users/:username", accounts.Detail),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

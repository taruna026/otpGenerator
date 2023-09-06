package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"sinarmas/controller"
	repo2 "sinarmas/repo"
	service2 "sinarmas/service"
	"strconv"
)

var db *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mysecretpassword"
	dbname   = "mydb"
)

func main() {
	connectionString :=
		"host=" + host +
			" port=" + strconv.Itoa(port) +
			" user=" + user +
			" password=" + password +
			" dbname=" + dbname +
			" sslmode=disable"

	var err error
	db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	repo := repo2.NewUserRepo(db)
	service := service2.NewUserService(repo)

	controller.NewUserController(service, router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

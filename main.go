package main

import (
	"customerCrud/customer/controller"
	"customerCrud/customer/service"
	"customerCrud/model"
	"customerCrud/repository"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/swabhav?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Connection Failed to Open", err)
		return
	}
	log.Println("Connection Established")

	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&model.Customer{})
	db.AutoMigrate(&model.Order{})
	db.Model(&model.Order{}).AddForeignKey("customer_id", "customers(id)", "CASCADE", "CASCADE")
	myRouter := mux.NewRouter().StrictSlash(true)
	repo := repository.NewRepository()
	services := service.NewService(&repo, db)
	control := controller.NewController(services)
	control.HandleRequests(myRouter)
	log.Fatal(http.ListenAndServe(":10000", myRouter))

}

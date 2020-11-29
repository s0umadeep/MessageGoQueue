package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type user struct{
	gorm.Model
	Name string
	Email string
}

func InitialMigration(){
	db,err = gorm.Open("sqlite3","test.db")
	if err!=nil{
		fmt.Print(err.Error())
		panic("Failed to connect to Database")
	}
	defer db.Close()
	db.AutoMigrate(&user{})
}
func AllMessage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"All message function hit")
	db,err=gorm.Open("sqlite3","test.db")
	if err!=nil{
		fmt.Print(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()
	var users [] user
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func CreateMessage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Create message function")
	db,err=gorm.Open("sqlite3","test.db")
	if err!=nil{
		fmt.Print(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	vars:=mux.Vars(r)
	name:=vars["Name"]
	email:=vars["Email"]

	db.Create(&user{Name:name,Email:email})
	fmt.Print("New user succesfully created : ")
}
func UpdateMessage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Update for the given message")
}
func DeleteMessage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Delete message function")
}
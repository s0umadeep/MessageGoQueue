package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Message struct{
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

type Messages[] Message

func MessageFunc(w http.ResponseWriter, r *http.Request){
	log.Print("Message Function Hit")
	MessageContent := Messages {
		Message{Title: "Message 1",Desc:"All about the new description",Content:"Full of content"},
	}
	json.NewEncoder(w).Encode(MessageContent)
}
func handlerFunc(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","text/html")
	if r.URL.Path=="/GoTime"{
		fmt.Fprint(w,"<h1> Welcome to Go Time  </h1>")
		fmt.Fprint(w,"<body> The Current Time is : </body>",time.Now().In(time.Local))
	}else if r.URL.Path=="/Contact"{
		fmt.Fprint(w,"<h1> Contact Us  </h1>")
		fmt.Fprint(w,"<body> You can contact us on : <a href =\"support@gotime.com\"> support@gotime.com </a> </body>")
	}
	log.Print("The current time is : ",time.Now().In(time.Local))

}

func PostMessageFunc(w http.ResponseWriter,r *http.Request){
	log.Print("Post Message Function")
	fmt.Fprintf(w,"Post Functions has been executed")
}

func ObjectRelationFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"ORM Functions executed")
}
func handlerFun(){
	router:=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",handlerFunc)
	router.HandleFunc("/GetMessage",MessageFunc).Methods("GET")
	router.HandleFunc("/ORM",ObjectRelationFunc).Methods("GET")
	router.HandleFunc("/Message",AllMessage).Methods("GET")
	router.HandleFunc("/Message/{name}/{email}",CreateMessage).Methods("POST")
	router.HandleFunc("/Message/{name}",DeleteMessage).Methods("DELETE")
	router.HandleFunc("/Message/{name}/{email}",UpdateMessage).Methods("PUT")
	http.ListenAndServe(":3000",router)
}

func main(){
	handlerFun()
	InitialMigration()
}
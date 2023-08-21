package web_server

import (
	"net/http"
	"fmt"
	"log"
	"github.com/dimalien/db-web-server-golang.git/db/mysql"
	//"github.com/dimalien/db-web-server-golang.git/db/model"
	)


func mainPageHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request)
	writer.Write([]byte("my main page")) 
}

func user1PageHendler(writer http.ResponseWriter, request *http.Request) {
	database := db.ConnnectToDataBase()
	user, _ := db.GetUserByID(1, database)
	writer.Write([]byte("user 1 info\n"))
	writer.Write([]byte(user.String()))
}

func user2PageHendler(writer http.ResponseWriter, request *http.Request) {
	database := db.ConnnectToDataBase()
	user, _ := db.GetUserByID(2, database)
	writer.Write([]byte("user 2 info\n"))
	writer.Write([]byte(user.String()))
}

func RunServer(port string) error {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(mainPageHandler))
	mux.Handle("/1/", http.HandlerFunc(user1PageHendler))
	mux.Handle("/2/", http.HandlerFunc(user2PageHendler))

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
	return nil
}


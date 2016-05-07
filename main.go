package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var wsMap map[int]*websocket.Conn
var counter int

type tester struct {
	ID      int    `json:"id"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func main() {
	counter = 0
	wsMap = make(map[int]*websocket.Conn)

	router := mux.NewRouter()

	router.HandleFunc("/ws", WsHandler)
	router.Methods("POST").Path("/api/notification").HandlerFunc(notificationHandler)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Panic(http.ListenAndServe("localhost:1337", router))
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	t := &tester{}
	err := decoder.Decode(t)

	if err != nil {
		log.Println(err)
	}

	wsMap[t.ID].WriteJSON(t)
}

package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/goim/chat"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "home.html")
}

func listUsers(hub *chat.Hub, w http.ResponseWriter, r *http.Request) {
	users := hub.GetUsers()

	userList, _ := json.Marshal(users)

	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Write(userList)
}

func main() {
	flag.Parse()
	hub := chat.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})
	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		listUsers(hub, w, r)
	})
	log.Println("Server is running on port:", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

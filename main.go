package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var addr = flag.String("addr", ":3000", "http service address")

func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	conn.WriteMessage(websocket.TextMessage, []byte("今天的天气真好"))
	// if err := conn.Close(); err != nil {
	// 	log.Println(err)
	// }
	for {
		conn.WriteMessage(websocket.TextMessage, []byte("今天的天气真好"))
		time.Sleep(time.Second)
	}
}

func main() {
	flag.Parse()
	serveStatic := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", serveStatic))
	http.HandleFunc("/client", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("static/index.html"))
	})
	http.HandleFunc("/ws", serveWs)
	log.Printf("listening on %s", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

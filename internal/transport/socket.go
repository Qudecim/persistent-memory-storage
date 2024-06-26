package transport

import (
	"flag"
	"log"
	"net/http"
	"qudecim/db/internal/app"
)

var addr = flag.String("addr", ":8080", "http service address")

func Run(app *app.App) {

	flag.Parse()
	hub := newHub(app)
	go hub.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

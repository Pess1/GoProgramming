package main

import (

)

type message struct {
	Text string 'json:"text"'
}

func main () {
	var port int

	flag.IntVar(&port, "p", 6969, "--p *port number*, default 6969")

	mux := http.NewServeMux()
	mux.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
		handler(ws, h)
	}))

	s

}

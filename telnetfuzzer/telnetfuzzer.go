package main

import (
	"fmt"
	"flag"
)

func main() {
	var address string
	var port string
	var uname string
	var pass string

	flag.StringVar(&address, "address", "0.0.0.0", "--address *Ip/host name*, default value 0.0.0.0")
	flag.StringVar(&port, "port", "8080", "--port *port number*, default value 8080")
	flag.StringVar(&uname, "username", "user", "--username *username*, default value user")
	flag.StringVar(&pass, "password", "password", "--password *password*, default value password")

	flag.Parse()

	fmt.Println("telnet", address, port, uname, pass)

}

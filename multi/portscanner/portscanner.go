package main

//Sources:
//https://medium.com/@KentGruber/building-a-high-performance-port-scanner-with-golang-9976181ec39d
//https://golang.org/pkg/net/

import (
	"fmt"
	"net"
	"strings"
	"time"
	"flag"
)

func scanPorts(ip string, port int, tout time.Duration, openOnly bool) {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, tout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			fmt.Println(err)
			time.Sleep(tout)
			scanPorts(ip, port, tout, openOnly)
		} else {
			if openOnly {
				fmt.Println("Port:", address, "===> closed")
			}
		}
		return
	}

	fmt.Println(conn, "Test")

	fmt.Println("Port:", address, "===> open")

	conn.Close()

}

func main() {
	var ip string
	var startport int
	var endport int
	var openOnly bool

	flag.StringVar(&ip, "a", "localhost", "--a *address*, default localhost")
	flag.IntVar(&startport, "sP", 80, "--sP *startport*, default 80")
	flag.IntVar(&endport, "eP", 0, "--eP *endport* if left empty scans just one port")
	flag.BoolVar(&openOnly, "c", false, "--c, type this to scan also closed ports")
	flag.Parse()

	tout := time.Second

	for startport <= endport {
		scanPorts(ip, startport, tout, openOnly)
		startport += 1
	}

	fmt.Println("=== SCAN COMPLETE ===")
}

// p47

package main

import (
	"log"
	"net"
	"time"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatalln("resolve tcp addr:", err)
	}

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		log.Fatalln("listen tcp:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listenner accept:", err)
		}
		daytime:= time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()	 // finished with this client
	}
}

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatalln("resolve tcp add:", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalln("listen tcp:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listenner accept:", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	var buf [512]byte
	defer conn.Close()
	for {
		_, err := conn.Read(buf[0:])
		if err != nil {
			log.Println("conn read:", err)
			return
		}

		fmt.Println("read from conn:", string(buf[0:]))

		_, err = conn.Write([]byte("i received your message"))
		if err != nil {
			log.Println("conn write:", err)
			return
		}
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("请传入tcp地址和端口, 如: localhost:1201")
	}

	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		log.Fatalln("tcp dial:", err)
	}

	defer conn.Close()

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		log.Fatalln("write to tcp:", err)
	}
	all, err := ioutil.ReadAll(conn)

	if err != nil {
		log.Fatalln("tcp read:", err)
	}

	fmt.Println(string(all))

	os.Exit(0)
}

package main

import (
	"net"
	"log"
	"time"
)

func main()  {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)

	if err != nil {
		log.Fatal("resolve tcp address failed")
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal("start server listener failed")
	}

	log.Printf("your tcp server will be run on %s", tcpAddr.String())
	for {
		//conn, err := listener.AcceptTCP()
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept tcp request failed")
			continue
		}
		go handleClient(conn.(*net.TCPConn))
	}
}

//1 使用接口
//  conn, err := listener.Accept()
//  go handleClient(conn)
//  func handleClient(conn net.Conn) {//}
//
//2 使用具体类型
//  conn, err := listener.AcceptTCP()
//  go handleClient(conn)
//  func handleClient(conn *TCPConn)
//
//3 使用类型断言
//  conn, err := listener.Accept()
//  go handleClient(conn.(*net.TCPConn))
//  func handleClient(conn *TCPConn)
//
//Go 建议使用接口，而不是具体的类型，
//但是这样做会损失具体类型的方法，如TCPConn.SetKeepAlive, UDPConn.SetReadBuffer等（除非做类型转换）
//用接口还是用类型取决于你自己的决定
func handleClient(conn *net.TCPConn)  {
	defer conn.Close()
	var buf [512]byte
	_, _ = conn.Write([]byte("\nyour client connect successful"))

	cnt := 0
	for {
		cnt++
		if cnt > 10 {
			return
		}
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		_, err = conn.Write(buf[0: n])
		if err != nil {
			return
		}

		daytime := time.Now().String()
		_, _ = conn.Write([]byte(daytime))
	}
}


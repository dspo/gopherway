package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("you have to input the remote address")
	}

	addr, err := net.ResolveIPAddr("ip", os.Args[1])
	if err != nil {
		log.Fatalln("invalid remote address")
	}

	conn, err := net.DialIP("ip4:icmp", addr, addr)
	if err != nil {
		log.Fatalln("dial ip:", err)
	}

	var msg [512]byte
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4] = 0
	msg[5] = 13
	msg[6] = 0
	msg[7] = 37

	length := 8

	sum := checkSum(msg[0:length])
	msg[2] = byte(sum>>8)
	msg[3] = byte(sum&225)

	log.Println(msg[0:length])
	_, err = conn.Write(msg[0:length])
	if err != nil {
		log.Fatalln("write to remote:", err)
	}

	_, err = conn.Read(msg[0:length])
	if err != nil {
		log.Fatalln("read from remote")
	}

	fmt.Println("got response", string(msg[0:]))
	if msg[5] == 13 {
		log.Println("identifier matches")
	}
	if msg[7] == 37 {
		log.Println("sequence matches")
	}

	os.Exit(0)
}

func checkSum(msg []byte) uint16 {
	sum := 0

	// assume even for now
	for n := 1; n < len(msg)-1; n+=2{
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum>>16) + (sum & 0xffff)
	sum+=(sum>>16)
	var answer uint16 = uint16(^sum)
	return answer
}

package main

import (
	"fmt"
	"github.com/micro/go-micro/web"
	"io/ioutil"
	"log"
)

func main() {
	s := web.NewService(web.Name("go.micro.web"))
	c := s.Client()
	response, err := c.Get("http://127.0.0.1:51353/websocket/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ioutil.ReadAll(response.Body))
	response.Body.Close()
}

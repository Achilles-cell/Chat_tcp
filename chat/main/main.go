package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("无法开启服务器：%s", err.Error())
	}

	defer listener.Close()
	log.Printf("开启服务器，端口为：8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("无法接收连接：%s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}

package main

import (
	"log"
	"net"
	"os"
)

func main() {

	/* 1. 환경설정 init */
	listen, error := net.Listen("tcp", ":10000")
	if error != nil {
		log.Println(error.Error())
		os.Exit(0)
	}

	/*
		service loop
	*/
	for {
		/*
			Client 로부터의 연결을 대기한다.
		*/
		conn, error := listen.Accept()
		if error != nil {
			log.Println(error.Error())
		}

		/*
			연결된 conn에서 메시지를 수신할 goroutine을 만든다.
		*/
		go proc_recv_msg(conn)
	}
}

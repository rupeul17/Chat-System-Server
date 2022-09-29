package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

/*
	global variable
*/
var Client_List [3]net.Conn
var Msg_List chan []byte

func main() {

	var Client_cnt int

	fmt.Println("Chat System Loading... OK.")

	/* 1. 환경설정 init */
	listen, error := net.Listen("tcp", ":10000")
	if error != nil {
		log.Println(error.Error())
		os.Exit(0)
	}

	go proc_send_msg()
	go proc_check_connection()

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
			현재 활성화된 세션이 최대일 경우 접속을 해제한다.
		*/
		if Client_cnt >= 3 {
			fmt.Println("Session is full... ")
			conn.Close()
		}

		for idx := range Client_List {
			if Client_List[idx] == nil {
				Client_List[idx] = conn
				Client_cnt++
				break
			}
		}

		/*
			연결된 conn에서 메시지를 수신할 goroutine을 만든다.
		*/
		fmt.Printf("Client Login:: Address(%s), Type(%s)\n", conn.RemoteAddr().String(), conn.RemoteAddr().Network())
		go proc_recv_msg(conn)
	}
}

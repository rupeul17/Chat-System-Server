package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func proc_recv_msg(Conn net.Conn) {

	recv := make([]byte, 4096)

	for {
		/*
			Client 로부터 msg를 수신한다.
		*/
		n, error := Conn.Read(recv)
		if error != nil {
			if error == io.EOF {
				fmt.Println("Connection is Close by Client\n", Conn.RemoteAddr().String())
			}

			log.Println(error.Error())

			for idx := range Client_List {
				if Client_List[idx] == Conn {
					fmt.Printf("(%s,%s) Disconnecting... ", Conn.RemoteAddr().String(), Conn.RemoteAddr().Network())
					Client_List[idx] = nil
					Client_cnt--
					fmt.Println("Disconnect success... Current Session : ", Client_cnt)
				}
			}
			break
		}

		Msg_List = make(chan []byte, 4096)

		if n > 0 {
			fmt.Printf("Recv From Client Msg (%s, %s)\n", Conn.RemoteAddr().String(), Conn.RemoteAddr().Network())
			Msg_List <- recv[:n]
		}
	}
}

func proc_check_accept() {

	/* 1. 환경설정 init */
	listen, error := net.Listen("tcp", ":10000")
	if error != nil {
		log.Println(error.Error())
		os.Exit(0)
	}

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
			fmt.Println("Session is full... Current Session : ", Client_cnt)
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

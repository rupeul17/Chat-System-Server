package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func proc_check_login_info(body_buf []byte) int {

	var count int
	login_info := DecodeToLoginInfo(body_buf)

	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/member_db")
	if err != nil {
		log.Println(err.Error())
		return -1
	}

	defer db.Close()

	result, err := db.Query("SELECT count(*) FROM member WHERE ID=? AND PASSWD=?", login_info.Id, login_info.Pwd)
	if err != nil {
		log.Println(err.Error())
		return -1
	}

	for result.Next() {
		if err := result.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	return count
}

func proc_recv_msg(Conn net.Conn) {

	recv := make([]byte, 4096)
	Msg_List = make(chan []byte, 4096)

	defer Conn.Close()

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

		if n > 0 {
			fmt.Printf("Recv From Client Msg (%s, %s)\n", Conn.RemoteAddr().String(), Conn.RemoteAddr().Network())

			clnt_msg := DecodeToMyMsg(recv)

			switch clnt_msg.Head.MsgType {
			case 2:
				switch proc_check_login_info(clnt_msg.Body) {
				case -1:
				case 0:
					proc_send_res_msg(Conn, 400)
				case 1:
					proc_send_res_msg(Conn, 200)
				}
			case 1:
				Msg_List <- recv[:n]
			}
		}
	}
}

func proc_listen() {

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

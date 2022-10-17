package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func CheckLoginInfo(body_buf []byte) int {

	var UserCount int
	login_info := Login{}
	json.Unmarshal([]byte(body_buf), &login_info)

	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/member_db")
	if err != nil {
		log.Println(err.Error())
		return DB_ACCESS_FAIL
	}

	defer db.Close()

	result, err := db.Query("SELECT count(*) FROM member WHERE ID=? AND PASSWD=?", login_info.Id, login_info.Pwd)
	if err != nil {
		log.Println(err.Error())
		return DB_SELECT_FAIL
	}

	for result.Next() {
		if err := result.Scan(&UserCount); err != nil {
			log.Fatal(err)
		}
	}

	if UserCount > 0 {
		return OK
	} else {
		return NOT_EXIST_USER_DATA
	}
}

func RequestHandler(Conn net.Conn) {

	recv := make([]byte, 4096)
	ClientMsgBodyList = make(chan []byte, 4096)

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

			for idx := range ClientConnList {
				if ClientConnList[idx] == Conn {
					ClientConnList[idx] = nil
					ClientConnCnt--
					fmt.Printf("Disconnect success... IP(%s, %s), Current Connected Client Count(%d)\n",
						Conn.RemoteAddr().String(), Conn.RemoteAddr().Network(), ClientConnCnt)
				}
			}
			break
		}

		if n > 0 {
			clnt_msg := DecodeToMyMsg(recv)

			switch clnt_msg.Head.MsgType {
			case TYPE_MESSAGE:
				ClientMsgBodyList <- recv[:n]
			case TYPE_LOGIN:
				SendMessageToClient(Conn, CheckLoginInfo(clnt_msg.Body))
			default:
			}
		}
	}
}

func ConnectHandler() {

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
		if ClientConnCnt >= 3 {
			fmt.Println("Session is full... Current Session : ", ClientConnCnt)
			conn.Close()
		}

		for idx := range ClientConnList {
			if ClientConnList[idx] == nil {
				ClientConnList[idx] = conn
				ClientConnCnt++
				break
			}
		}

		/*
			연결된 conn에서 메시지를 수신할 goroutine을 만든다.
		*/
		fmt.Printf("Client Login:: Address(%s), Type(%s), Current Client Count(%d)\n", conn.RemoteAddr().String(), conn.RemoteAddr().Network(), ClientConnCnt)
		go RequestHandler(conn)
	}
}

package main

import (
	"fmt"
	"io"
	"log"
	"net"
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

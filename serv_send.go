package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func SendMessageToClient(conn net.Conn, r int) {

	msg := MyMsg{
		Head: Header{
			MsgType: TYPE_LOGIN,
			Res:     r,
			BodyLen: 0,
		},
	}

	bytedata := EncodeToBytes(msg)

	_, error := conn.Write(bytedata)
	if error != nil {
		log.Println(error.Error())
	}
}

func BroadcastToClient() {

	for {
		select {
		case Msg, ok := <-ClientMsgBodyList:
			if ok {
				for idx := range ClientConnList {
					if ClientConnList[idx] != nil {
						ClientConnList[idx].Write(Msg)
					}
				}
			} else {
				fmt.Printf("Channel Close\n")
			}
		default:
		}
		time.Sleep(1 * time.Second)
	}
}

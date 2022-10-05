package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func proc_send_res_msg(conn net.Conn, r int) {

	msg := MyMsg{
		Head: Header{
			MsgType: 9,
			Ip:      conn.LocalAddr().String(),
			Res:     r,
			BodyLen: 0,
		},
	}

	bytedata := EncodeToBytes(msg)

	_, error := conn.Write(bytedata)
	if error != nil {
		log.Println(error.Error())
	}

	fmt.Println("MSG SEND >> res : ", msg.Head.Res)

}

func proc_broadcast() {

	for {
		select {
		case Msg, ok := <-Msg_List:
			if ok {

				fmt.Printf("Msg Send to Client>>\n")

				for idx := range Client_List {
					if Client_List[idx] != nil {
						Client_List[idx].Write(Msg)
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

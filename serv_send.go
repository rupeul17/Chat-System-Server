package main

import (
	"fmt"
	"time"
)

func proc_send_msg() {

	for {
		select {
		case Msg, ok := <-Msg_List:
			if ok {

				fmt.Printf("Msg Send to Client>> %s\n", Msg)

				for idx := range Client_List {
					if Client_List[idx] != nil {
						Head := []byte(Client_List[idx].RemoteAddr().String() + " >> ")
						Msg = append(Head, Msg...)
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

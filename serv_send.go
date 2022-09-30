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

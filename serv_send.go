package main

import "fmt"

func proc_send_msg() {

	for range Msg_List {
		Msg := <-Msg_List
		fmt.Printf("Msg Send to Client>> %s\n", Msg)

		for idx := range Client_List {
			if Client_List[idx] != nil {
				Client_List[idx].Write(Msg)
			}
		}
	}
}

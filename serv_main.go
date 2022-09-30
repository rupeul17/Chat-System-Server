package main

import (
	"fmt"
	"net"
	"time"
)

/*
	global variable
*/
var Client_cnt int
var Client_List [3]net.Conn
var Msg_List chan []byte

func main() {

	fmt.Println("Chat System Loading... OK.")

	/*
		고루틴을 생성한다.
		1) send_msg    : Client 로 메시지 전송
		2) command_msg : Command 처리
	*/
	go proc_send_msg()
	go proc_check_accept()

	/*
		service loop
	*/
	for {
		fmt.Println("==========Chat System Command==========")
		fmt.Println(" 1. All Connection destroy.")
		fmt.Println(" 2. Change Configuration.")
		fmt.Println(" 0. Shutdown System.")

		cmd_num := input_number()

		if cmd_num == 1 {

		} else if cmd_num == 2 {

		} else if cmd_num == 0 {

		}

		time.Sleep(2 * time.Second)
	}
}

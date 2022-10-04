package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

/*
	global variable
*/
var Client_cnt int
var Client_List [3]net.Conn
var Msg_List chan []byte
var Max_Session_Cnt int

func main() {

	fmt.Println("Chat System Loading... OK.")

	Max_Session_Cnt = 3

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
		ClearTerminal()
		fmt.Println("==========Chat System Command==========")
		fmt.Println(" 1. All Connection destroy.")
		fmt.Println(" 2. Change Configuration.")
		fmt.Println(" 0. Shutdown System.")
		fmt.Println("=======================================")

		cmd_num := input_number()

		switch cmd_num {
		case 1:
			/*
				현재 연결된 모든 세션을 끊는다.
			*/
			for idx := range Client_List {
				if Client_List[idx] != nil {
					Client_List[idx].Close()
					Client_List[idx] = nil
				}
			}
			fmt.Println("All Session is clear...")

		case 2:
			/*
				설정을 변경한다.
			*/
			proc_chg_conf()

		case 0:
			/*
				server를 종료한다.
			*/
			fmt.Println("Chat System Shutdown")
			os.Exit(0)

		}

		time.Sleep(1 * time.Second)
	}
}

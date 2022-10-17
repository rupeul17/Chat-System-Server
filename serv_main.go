package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
	global variable
*/
var ClientConnCnt int
var ClientConnList [100]net.Conn
var ClientMsgBodyList chan []byte
var MaxClientConnCnt int

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Chat System Loading... OK.")

	MaxClientConnCnt = 3

	/*
		고루틴을 생성한다.
		1) send_msg    : Client 로 메시지 전송
		2) command_msg : Command 처리
	*/

	go BroadcastToClient()
	go ConnectToClient()

	go func() {
		/*
			signal check
		*/
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	go func() {
		/*
			command check
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
				for idx := range ClientConnList {
					if ClientConnList[idx] != nil {
						ClientConnList[idx].Close()
						ClientConnList[idx] = nil
					}
				}
				fmt.Println("All Session is clear...")

			case 2:
				/*
					설정을 변경한다.
				*/
				ChangeConfiguration()

			case 0:
				/*
					server를 종료한다.
				*/
				fmt.Println("Chat System Shutdown")
				os.Exit(0)
			}

			time.Sleep(1 * time.Second)
		}
	}()

	<-done
	fmt.Println("Exiting")

}

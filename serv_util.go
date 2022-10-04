package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func DecodeToLoginInfo(s []byte) Login {

	login_info := Login{}
	dec := gob.NewDecoder(bytes.NewReader(s))

	err := dec.Decode(&login_info)
	if err != nil {
		log.Fatal(err)
	}

	return login_info
}

func DecodeToMyMsg(s []byte) MyMsg {

	myMsg := MyMsg{}
	dec := gob.NewDecoder(bytes.NewReader(s))

	err := dec.Decode(&myMsg)
	if err != nil {
		log.Fatal(err)
	}

	return myMsg
}

func EncodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
}

func input_string() string {
	rd := bufio.NewReader(os.Stdin)

	TmpStr, err := rd.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	TmpStr = strings.TrimSpace(TmpStr)

	return TmpStr
}

func input_number() int {
	rd := bufio.NewReader(os.Stdin)

	TmpInt, err := rd.ReadString('\n')
	if err != nil {
		log.Println(err.Error())
		return -1
	}

	TmpInt = strings.TrimSpace(TmpInt)
	input_num, err := strconv.Atoi(TmpInt)
	if err != nil {
		log.Println(err.Error())
		return -1
	}

	return input_num
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

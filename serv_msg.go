package main

/*
	Defined Client <-> Server Message Protocol
*/

const (
	TYPE_LOGIN = iota + 1
	TYPE_MESSAGE
)

const (
	OK = iota
	NOT_EXIST_USER_DATA
	DB_ACCESS_FAIL
	DB_SELECT_FAIL
	DB_UPDATE_FAIL
	DB_DELETE_FAIL
)

type Header struct {
	MsgType int
	BodyLen int
	Res     int
}

type MsgInfo struct {
	Id      string
	Message []byte
}

type Login struct {
	Id  string
	Pwd string
}

type MyMsg struct {
	Head Header
	Body []byte
}

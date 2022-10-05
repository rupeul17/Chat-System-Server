package main

/*
	Defined Client <-> Server Message Protocol
*/
type Header struct {
	MsgType int
	BodyLen int
	Res     int
	Ip      string
}

type Login struct {
	Id  string
	Pwd string
}

type MyMsg struct {
	Head Header
	Body []byte
}

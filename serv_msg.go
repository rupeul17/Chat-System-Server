package main

/*
	Defined Client <-> Server Message Protocol
*/
type Header struct {
	MsgType int
	Ip      string
	BodyLen int
	res     int
}

type Login struct {
	Id  string
	Pwd string
}

type MyMsg struct {
	Head Header
	Body []byte
}

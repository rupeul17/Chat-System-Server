package main

/*
	Defined Client <-> Server Message Protocol
*/
type Header struct {
	MsgType int
	Ip      string
	BodyLen int
}

type MyMsg struct {
	Head Header
	Body []byte
}

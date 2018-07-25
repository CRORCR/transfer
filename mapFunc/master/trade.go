package main

import (
	"fmt"
	"math/rand"
)

func getMessage() {
	for{
		createMessage()
	}
}

var chMess = make(chan []string)

func createMessage() {
	var send = make([]string, 0)
	for i := 0; i < 10000; i++ {
		s := random()
		send = append(send, s)
	}
	closecreateMess<-"ok"
	chMess <- send
	fmt.Printf("生成了:%d 数据",len(send))
	send = make([]string, 0)
}

var str = "abcdefghighlmnopqrstuvwxyzABCDEFGHIGHLMNOPQRSTUVWXYZ0123456789"

func random() (stri string) {
	s := []byte(str)
	for i := 0; i < 203; i++ {
		i2 := rand.Intn(62)
		stri = stri + string(s[i2])
	}
	return stri
}

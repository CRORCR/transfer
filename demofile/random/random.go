package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for{
		createMessage()
	}
}

var chMess = make(chan []string)
var count int32=0

func createMessage() {
	q := rand.Intn(5)
	b := rand.Intn(10)
	g := rand.Intn(10)
	if q==0{q=1}
	if b==0{b=1}
	num:=40000+q*1000+b*100+4*10+g
	fmt.Println(num)
}

//存入数据库的数据,需要排序,使用自增为key,达到12秒(12万)数据,再取出来 打包,存入数据库
var levelDBSlice=make([]string,0)

var str = "abcdefghighlmnopqrstuvwxyzABCDEFGHIGHLMNOPQRSTUVWXYZ0123456789"

func random() (stri string) {
	s := []byte(str)
	for i := 0; i < 5; i++ {
		i2 := rand.Intn(62)
		stri = stri + string(s[i2])
	}
	return stri
}

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func getMessage() {
	for{
		createMessage()
	}
}

var chMess = make(chan []string)
var count int32=0

func createMessage() {
	var send = make([]string, 0)
	for i := 0; i < 10000; i++ {
		s := random()
		send = append(send, s)
	}
	closecreateMess<-"ok"
	saveLevelDB(send)
	chMess <- send
	fmt.Printf("生成了:%d 数据\n", len(send))
	send = make([]string, 0)
}

//存入数据库的数据,需要排序,使用自增为key,达到12秒(12万)数据,再取出来 打包,存入数据库
var levelDBSlice=make([]string,0)

func saveLevelDB(send []string) {
	fmt.Println("数据量",atomic.LoadInt32(&count))
	atomic.AddInt32(&count, 1)
	isBlock := atomic.CompareAndSwapInt32(&count, 12, 0)

	bytes, err := json.Marshal(send)
	if err != nil {
		fmt.Println("存储序列化失败")
		return
	}
	intType := fmt.Sprintf("%v",time.Now().UnixNano())
	levelPut([]byte(intType), bytes)
	//0-12w个消息 存入slice中,打包的时候,需要情况这个切片
	levelDBSlice=append(levelDBSlice,intType)

	//fmt.Printf("levelDBSlice:%+v\n",levelDBSlice)
	if isBlock {
		block=append(block,levelDBSlice)
		levelDBSlice=make([]string,0)
		fmt.Println("12条记录,存储block")
		//fmt.Printf("%+v\n",block)
	}
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

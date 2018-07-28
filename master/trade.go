package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

var timeSaveBlock = time.NewTicker(12 * time.Second)

func getMessage() {
	var blockSave = make([]string, 0)

	for {
		select {
			//blockchan <- blockSave
		case <-timeSaveBlock.C:
			go saveLevelDB()
			fmt.Println("12秒一共多少数据", len(blockSave))
			blockchan <- blockSave
			blockSave = make([]string, 0)
		default:
			var send = make([]string, 0)
			q := rand.Intn(5)
			b := rand.Intn(10)
			g := rand.Intn(10)
			if q == 0 {
				q = 1
			}
			if b == 0 {
				b = 1
			}
			num := config.SendNum + q*1000 + b*100 + 4*10 + g
			for i := 0; i < num; i++ {
				s := random()
				send = append(send, s)
				blockSave = append(blockSave, s)
			}
			chMess <- send
		}
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}

var chMess = make(chan []string)

var blockchan = make(chan []string)

//var count int32=0

//func createMessage() {
//	q := rand.Intn(5)
//	b := rand.Intn(10)
//	g := rand.Intn(10)
//	if q == 0 {
//		q = 1
//	}
//	if b == 0 {
//		b = 1
//	}
//
//	num = config.sendNum + q*1000 + b*100 + 4*10 + g
//	go func() {
//		for {
//			s := random()
//			send = append(send, s)
//		}
//	}()
//	if len(send) >= num {
//		blockSave = append(blockSave, send[:num]...)
//		fmt.Println("chan")
//		closecreateMess <- "ok"
//		chMess <- send[:num]
//		send = send[num:]
//		fmt.Println("chan结束")
//		send = make([]string, 0)
//	}
//	//fmt.Printf("生成了:%d 数据\n", len(send))
//}

//存入数据库的数据,需要排序,使用自增为key,达到12秒(12万)数据,再取出来 打包,存入数据库
//var levelDBSlice=make([]string,0)

//对应哪个块一共多少交易
var blockNumber = make(map[string]int)

func saveLevelDB() {
	//fmt.Println("数据量",atomic.LoadInt32(&count))
	//atomic.AddInt32(&count, 1)
	//isBlock := atomic.CompareAndSwapInt32(&count, 12, 0)
	//start := time.Now().UnixNano() / 1e6
	send := <-blockchan
	fmt.Println("save")
	bytes, err := json.Marshal(send)
	//end := time.Now().UnixNano() / 1e6
	//fmt.Println(end - start)
	if err != nil {
		fmt.Println("存储序列化失败")
		return
	}
	intType := fmt.Sprintf("%v", time.Now().UnixNano())
	//start = time.Now().UnixNano() / 1e6
	levelPut([]byte(intType), bytes)
	//end = time.Now().UnixNano() / 1e6
	//fmt.Println(end - start)
	//0-12w个消息 存入slice中,打包的时候,需要情况这个切片
	//levelDBSlice=append(levelDBSlice,intType)

	//fmt.Printf("levelDBSlice:%+v\n",levelDBSlice)
	//if isBlock {
	blockNumber[intType] = len(send)
	fmt.Println("len", len(send))
	block = append(block, intType)
	fmt.Printf("block:%+v\n", block)
	//levelDBSlice=make([]string,0)
	//fmt.Println("12条记录,存储block")
	//fmt.Printf("%+v\n",block)
	//}
}

var str = []byte("abcdefghighlmnopqrstuvwxyzABCDEFGHIGHLMNOPQRSTUVWXYZ0123456789")

func random() (stri string) {
	for i := 0; i < 203; i++ {
		i2 := rand.Intn(62)
		stri = stri + string(str[i2])
	}
	return stri
}

package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"
)

const(
	//连接主节点
	//ADDR_1 ="localhost:9004"
	ADDR_1 ="192.168.1.4:9004"
	recvMessageNum=10000
	train=120000
	trainNum       = 12 //12秒轮训
)

//需要连接的地址
var addrList = []string{ADDR_1}

var block =make([][]string,0)
//var messMap =make(map[string][]byte)
var sendMess =make(chan string)
var levelDB *LevelDB
var savePackage=false
var tickerEnd *time.Ticker=time.NewTicker(1 * time.Hour)
var tickerStart *time.Ticker=time.NewTicker(1 * time.Hour)

func main() {
	//time.AfterFunc(12*time.Second, func() {
	//	saveLevel()
	//})
	levelDB = NewLevelDB()
	go Server()
	for {
		select {
		case <-sendMess:
			levelDBSlice=make([]string,0)
			//saveLevel()todo
			//go Client()
			fmt.Println("去重存储 success")
		case <-tickerStart.C:
			//se24:=time.Now().UnixNano()/1e6
			//fmt.Println("24秒:",se24-startTimes)//todo
			savePackage = true
			//开始准备打包
		case <-tickerEnd.C:
			//12秒通知其他节点处理包
			//callOther()
			//se36:=time.Now().UnixNano()/1e6
			//fmt.Println("36秒:",se36-startTimes)//todo
			savePackage = false
			nano := fmt.Sprintf("%v", time.Now().UnixNano())
			//fmt.Println("保存的数据:",MessSlcie)//todo
			bytes, _ := json.Marshal(levelDB.MessSlcie)
			levelPut([]byte(nano), bytes)
			levelDB.MessSlcie = make([]string, 0)      //保存完,置空
			fmt.Println("保存后应该清空:", levelDB.MessSlcie) //todo
		default:
		}
	}
}

func saveLevel() {
	savePackage = false
	//第一次打包控制,以后都是有定时器控制
	nano := fmt.Sprintf("%v", time.Now().UnixNano())
	levelDB.lock.RLock()
	bytes, _ := json.Marshal(levelDB.MessSlcie)
	levelDB.lock.RUnlock()
	//test
	//fmt.Printf("为什么是空的 %v", levelDB.MessSlcie)
	sum := md5.Sum([]byte(levelDB.MessSlcie[0]))
	sum2 := md5.Sum([]byte(levelDB.MessSlcie[len(levelDB.MessSlcie)-1]))
	fmt.Println("hash", sum)
	fmt.Println("hash", sum2)
	levelPut([]byte(nano), bytes)
	levelDB.MessSlcie = make([]string, 0)
	//保存完,置空
	//saveBytes=[]byte("hello world")
	tickerEnd = time.NewTicker(36 * time.Second)
	tickerStart = time.NewTicker(24 * time.Second)
	//after:=time.Now().UnixNano()/1e6
	//fmt.Println("after Func",after-startTimes)//todo
}
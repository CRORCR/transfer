package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

const (
	//连接主节点
	//ADDR_1 = "localhost:9004"
	ADDR_1         = "192.168.1.4:9004"
	recvMessageNum = 10000
	train          = 120000
	trainNum       = 12 //12秒轮训
)

//需要连接的地址
var addrList = []string{ADDR_1}
var block = make([][]string, 0)
//var messMap =make(map[string][]byte)
var sendMess = make(chan string)
var levelDB *LevelDB
var savePackage = false
var tickerEnd *time.Ticker = time.NewTicker(1 * time.Hour)
var tickerStart *time.Ticker = time.NewTicker(1 * time.Hour)
var selectLevel *time.Ticker = time.NewTicker(1 * time.Hour)
var cleanDB *time.Ticker = time.NewTicker(60 * time.Second)

func main() {
	//time.AfterFunc(12*time.Second, func() {
	//	saveLevel()
	//})
	levelDB = NewLevelDB()
	go Server()
	for {
		select {
		case <-sendMess:
			//打包
			fmt.Println("client 9004节点")
			go Client()
		case <-tickerStart.C:
			//se24:=time.Now().UnixNano()/1e6
			//fmt.Println("24秒:",se24-startTimes)//todo
			//savePackage = true
			//开始准备打包
			fmt.Println("超时~~~")
		case <-tickerEnd.C:
			fmt.Println("超时~~~")
			//12秒通知其他节点处理包
			//callOther()
			//se36:=time.Now().UnixNano()/1e6
			//fmt.Println("36秒:",se36-startTimes)//todo
			//savePackage = false
			//nano := fmt.Sprintf("%v", time.Now().UnixNano())
			////fmt.Println("保存的数据:",MessSlcie)//todo
			//bytes, _ := json.Marshal(levelDB.MessSlcie)
			//levelDB.levelPut([]byte(nano), bytes)
			//levelDB.MessSlcie = make([]string, 0)      //保存完,置空
			//fmt.Println("保存后应该清空:", levelDB.MessSlcie) //todo
		case <-selectLevel.C:
			GetBlockKey2 := GetBlockKey()
			fmt.Printf("我就想看看存的是啥:%+v\n", GetBlockKey2)
			for k, v := range GetBlockKey2 {
				fmt.Printf("key:%v value:%v\n", k, v)
			}

			fmt.Println("查询对应的数据 1 hour")
			s := GetBlockKey2[0][2]
			page := GetPage(s, 10, 20)
			fmt.Printf("我就想看看package查询的分页是啥:%v\n", page)
		case <-cleanDB.C:
			db.Close()
			time.Sleep(10*time.Nanosecond)
			err := os.RemoveAll("./db")
			if err != nil {
				panic("删除db失败")
			}
			db, _ = leveldb.OpenFile("./db", nil)

			block = make([][]string, 0)
			levelDBSlice=make([]string,0)
			levelDB.MessSlcie=make([]string,0)
			levelDB.MessMap=make(map[[16]byte]int,0)
			writeInfo = make(chan []string,1)
		default:
		}
	}
}

func saveLevel() {
	//savePackage = false
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

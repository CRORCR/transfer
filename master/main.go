package main

import (
	"fmt"
	"time"

)

//程序还可以优化,----发送数据使用string 转 byte发送    去重使用byte 16位
//主节点 1.2 9002
//从节点 1.3 9003
//从节点 1.4 9004
const (
	//从节点ip
	ADDR_1,ADDR_2 = "localhost:9003","localhost:9004"
	//ADDR_1,ADDR_2 = "192.168.1.3:9003","192.168.1.4:9004"
	recvMessageNum=10000
	train=120000
)
//需要连接的地址
var addrList = []string{ADDR_2,ADDR_1}
//var addrList = []string{ADDR_1,ADDR_2}
//已经连过的地址
//var historyAddrList []string

var closecreateMess = make(chan string)

var levelDB *LevelDB

//存储levelDB
var tickerEnd =time.NewTicker(1 * time.Hour)
var tickerStart =time.NewTicker(1 * time.Hour)
var selectLevel *time.Ticker = time.NewTicker(10 * time.Hour)
//以上是提供webServer相关参数
var block = make([]string, 0)

var sendMess =make(chan string,1)//收到数据,可以往其他节点发送

//test
//var startTimes=time.Now().UnixNano()/1e6 //todo
func main() {
	initConf()
	levelDB=NewLevelDB()
	//ticker := time.NewTicker(12 * time.Second)
	/*time.AfterFunc(12*time.Second, func() {
		//callOther()
		savePackage=false //第一次打包控制,以后都是有定时器控制
		nano := fmt.Sprintf("%v",time.Now().UnixNano())
		bytes, _ := json.Marshal(saveBytes)
		keySlice=append(keySlice,nano)
		openLevelDB()
		levelPut([]byte(nano),bytes)

		//test for webServer 后期需要抽取test用例
		/*sum := md5.Sum(saveBytes[0])
		sum2 := md5.Sum(saveBytes[len(saveBytes)-1])
		fmt.Println("hash",sum)
		fmt.Println("hash",sum2)

		blockKery:=levelDB.GetBlockKey()
		fmt.Println("一共存储多少块",len(blockKery))

		page := levelDB.GetPage(blockKery[0], 0, 20)
		fmt.Println("0-20条数据是:")
		for _,v:=range page{
			fmt.Printf("%s\n",v)
		}

		closeLevelDB()

		saveBytes=make([]string,0)//保存完,置空
		//saveBytes=[]byte("hello world")
		tickerEnd = time.NewTicker(36 * time.Second)
		tickerStart = time.NewTicker(24 * time.Second)

		//after:=time.Now().UnixNano()/1e6
		//fmt.Println("after Func",after-startTimes)//todo
	})*/
	go getMessage()
	go webServer()
	go ServerListen()

	for{
		select {
		case <-closecreateMess:
			//fmt.Println("开始client")
			//go SendPeer()
			go Client()
		case <-selectLevel.C:
			GetBlockKey2 := GetBlockKey()
			fmt.Printf("我就想看看存的是啥:%+v\n", GetBlockKey2)
			for k, v := range GetBlockKey2 {
				fmt.Printf("key:%v value:%v\n", k, v)
			}

			fmt.Println("查询对应的数据 1 hour")
			s := GetBlockKey2[0]
			page := GetPage(s, 10, 20)
			fmt.Printf("我就想看看package查询的分页是啥:%v\n", page)
		case <-tickerStart.C:
			fmt.Println("程序结束")
			//se24:=time.Now().UnixNano()/1e6
			//fmt.Println("24秒:",se24-startTimes)//todo
			//savePackage=true
			//开始准备打包
		case <-tickerEnd.C:
			fmt.Println("程序结束")
		default:
		}
	}
}

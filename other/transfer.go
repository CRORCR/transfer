package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net"
	"net/rpc"
	"sync/atomic"
	"time"
)

var timeSaveBlock = time.NewTicker(12 * time.Second)

//服务端监听自己端口
func Server() {
	listen, err := net.Listen("tcp", ":9003")
	if err != nil {
		fmt.Println("端口占用")
	}

	for {
		select {
		case <-timeSaveBlock.C:
			//block = append(block, levelDBSlice)
			//fmt.Println("6秒",atomic.LoadInt32(&count))
			levelDB.MessMap = make(map[[16]byte]int, 0)
			//sendMess <- "ok"
			//writeInfo <- levelDBSlice
			//levelDBSlice = make([]string, 0)
			return
		default:
			conn, err := listen.Accept()
			if err != nil {
				fmt.Println("connection is fail,err: ", err)
				continue
			}
			proess(conn)
		}
	}
}

var writeInfo = make(chan []string)

func proess(conn net.Conn) {
	defer conn.Close()
	for {
		gotMap := make([]string, 0)
		start := time.Now().UnixNano() / 1e6
		json.NewDecoder(conn).Decode(&gotMap)
		end := time.Now().UnixNano() / 1e6

		if len(gotMap) != 0 {
			fmt.Println("接收数据时间:", end-start)
			levelDB.manage(gotMap)
			gotMap = make([]string, 0)
		} else {
			continue
		}
	}
}


func (level LevelDB) manage(gotMap []string) {
	var count int32
	//去重
	start := time.Now().UnixNano() / 1e6
	for _, v := range gotMap {
		sum := md5.Sum([]byte(v))
		levelDB.lock.Lock()
		//fmt.Printf("收到的数据:%s\n", v)
		if _, ok := levelDB.MessMap[sum]; ok {
			//fmt.Printf("重复的是啥 line:%s hash:%v \n", v, sum)
			levelDB.lock.Unlock()
			continue
		}else if len(levelDB.MessMap)==500000{
			levelDB.MessMap=make(map[[16]byte]int,0)
			levelDB.lock.Unlock()
		} else {
			levelDB.MessMap[sum] = 1
			//levelDB.MessSlcie = append(levelDB.MessSlcie, v)
			atomic.AddInt32(&count, 1)
			levelDB.lock.Unlock()
		}
	}
	writeInfo<-gotMap
	end := time.Now().UnixNano() / 1e6
	fmt.Println("去重", end-start)
	fmt.Println("一共多少数据", atomic.LoadInt32(&count))
	//去重完成,存入数据库
	//bytes, err := json.Marshal(levelDB.MessSlcie)
	//if err != nil {
	//	fmt.Println("存储序列化失败")
	//	return
	//}
	//intType := fmt.Sprintf("%v", time.Now().UnixNano())
	//levelPut([]byte(intType), bytes)
	////0-12w个消息 存入slice中,打包的时候,需要情况这个切片
	//levelDBSlice = append(levelDBSlice, intType)
	return
}
//存入数据库的数据,需要排序,使用自增为key,达到12秒(12万)数据,再取出来 打包,存入数据库
//var levelDBSlice = make([]string, 0)

func Client() {
	fmt.Println("进入client")
	//主节点需要往从节点都发送数据
	conn := dialSer(ADDR_1)
	defer conn.Close()
	for {
		select {
		case addMessage := <-writeInfo:
			json.NewEncoder(conn).Encode(addMessage)
		default:
		}
	}
}

func dialSer(ip string) (conn net.Conn) {
	for {
		conn, err := net.Dial("tcp", ip)
		if err != nil {
			continue
		} else {
			//fmt.Println("成功建立连接")
			return conn
		}
	}
}

func DialRpc() (client *rpc.Client, err error) {
	for {
		client, err = rpc.DialHTTP("tcp", ADDR_1)
		if err != nil {
			continue
		} else {
			return
		}
	}
	return
}

package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net"
	"net/rpc"
	"time"
)

//服务端监听自己端口
func Server() {
	listen, err := net.Listen("tcp", ":9003")
	if err != nil {
		fmt.Println("端口占用")
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("connection is fail,err: ", err)
			continue
		}
		ch <- struct{}{}
		go proess(conn)
	}
}

var writeInfo = make(chan []string, 1)
var ch = make(chan interface{}, 1)

func proess(conn net.Conn) {
	for {
		gotMap := make([]string, 0)
		start := time.Now().UnixNano() / 1e6
		json.NewDecoder(conn).Decode(&gotMap)
		end := time.Now().UnixNano() / 1e6

		if (end - start) == 0 {
			<-ch
			return
		}
		fmt.Println("接收数据时间:", end-start)

		if len(gotMap) != 0 {
			conn.Close()
			levelDB.manage(gotMap)
			gotMap = make([]string, 0)
			if len(levelDBSlice) >= trainNum {
				//打包时间 生成一个包   12w数据备份清空,map去重清空 只留下levelDB供webserver使用
				block = append(block, levelDBSlice)
				fmt.Println("已经有12w了")
				fmt.Printf("block:%+v\n", block)
				levelDB.MessMap = make(map[[16]byte]int, 0)
				//fmt.Println("然后map置空,下一个12秒,再次计:",len(levelDB.MessMap)) //success
				sendMess <- "ok"
				writeInfo <- levelDBSlice
				levelDBSlice = make([]string, 0)
				//下次重新计数,打包/发送之前,用于存储数据库 作为key值,方便取出
				<-ch
				return
			}
			//return
		} else {
			<-ch
			return
		}
	}
}

//存入数据库的数据,需要排序,使用自增为key,达到12秒(12万)数据,再取出来 打包,存入数据库
var levelDBSlice = make([]string, 0)

func (level LevelDB) manage(gotMap []string) {
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
		} else {
			levelDB.MessMap[sum] = 1
			levelDB.MessSlcie = append(levelDB.MessSlcie, v)
			levelDB.lock.Unlock()
		}
	}
	end := time.Now().UnixNano() / 1e6
	fmt.Println("去重", end-start)
	//去重完成,存入数据库
	bytes, err := json.Marshal(levelDB.MessSlcie)
	if err != nil {
		fmt.Println("存储序列化失败")
		return
	}
	intType := fmt.Sprintf("%v", time.Now().UnixNano())
	levelPut([]byte(intType), bytes)
	//0-12w个消息 存入slice中,打包的时候,需要情况这个切片
	levelDBSlice = append(levelDBSlice, intType)
	return
}

func Client() {
	fmt.Println("进入client")
	fmt.Println("levelNum:", len(levelDBSlice))
	//主节点需要往从节点都发送数据

	select {
	case send := <-writeInfo:
		for i := 0; i < len(send); i++ {
			conn := dialSer(ADDR_1)
			key, err := GetKey(send[i])
			if err != nil {
				conn.Close()
				continue
			}
			sendMess := make([]string, recvMessageNum)
			json.Unmarshal(key, sendMess)
			err = json.NewEncoder(conn).Encode(&sendMess)
			if err != nil {
				//fmt.Println("序列化失败")
				return
			}
			conn.Close()
		}
	default:

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

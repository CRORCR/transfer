package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

type Salve struct{
	MessMap map[[16]byte]int
	lock sync.Mutex
}
var salve Salve
func NewSalve(){
	salve=Salve{MessMap:make(map[[16]byte]int)}
}

//:= make(map[[16]byte]int,0)
//服务端监听自己端口
func main() {
	listen, err := net.Listen("tcp", ":9004")
	NewSalve()
	if err != nil {
		fmt.Println("端口占用")
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("connection is fail,err: ", err)
			continue
		}
		go proess(conn)
	}
}

func proess(conn net.Conn ) {
	defer conn.Close()
	for {
		gotMap := make([]string, 0)
		//start := time.Now().UnixNano() / 1e6
		json.NewDecoder(conn).Decode(&gotMap)
		//end := time.Now().UnixNano() / 1e6

		if len(gotMap) != 0 {
			//fmt.Println("接收数据时间:", end-start)
			manage(gotMap)
			gotMap=nil
		} else {
			continue
		}
	}
}

func manage(gotMap []string) {
	//去重
	//	start := time.Now().UnixNano() / 1e6
	for _, v := range gotMap {
		sum := md5.Sum([]byte(v))
		salve.lock.Lock()
		if _, ok := salve.MessMap[sum]; ok {
			salve.lock.Unlock()
			continue
		}else {
			salve.MessMap[sum] = 1
			salve.lock.Unlock()
		}
	}
	if len(salve.MessMap) >= 50000 {
		salve.lock.Lock()
		salve.MessMap = make(map[[16]byte]int, 0)
		salve.lock.Unlock()
	}

	//	end := time.Now().UnixNano() / 1e6
	//if (end - start) != 0 {
	//	fmt.Println("去重", end-start)
	//	fmt.Println("一共多少数据", len(messSlcie))
	//}
	//去重完成,存入数据库
	//bytes, err := json.Marshal(levelDB.MessSlcie)
	//if err != nil {
	//	fmt.Println("存储序列化失败")
	//	return
	//}
	//intType := fmt.Sprintf("%v", time.Now().UnixNano())
	//levelPut([]byte(intType), bytes)
	//0-12w个消息 存入slice中,打包的时候,需要情况这个切片
	//levelDBSlice = append(levelDBSlice, intType)
	return
}

package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"
)

var blockSave = make([]string, 0)
var messMap = make(map[[16]byte]int, 0)

const BLOCKTIME=12
func main() {
	listen, err := net.Listen("tcp", ":9004")
	if err != nil {
		fmt.Println("端口占用")
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("connection is fail,err: ", err)
			continue
		}
		timeSaveBlock = time.NewTicker(BLOCKTIME * time.Second)
		go proess(conn)
	}
}

var timesave int64

var timeSaveBlock *time.Ticker

func proess(conn net.Conn) {
	defer conn.Close()
	for {
		gotMap := make([]string, 0)
		buf := make([]byte, 1024*1024*30)

		n, err := conn.Read(buf)
		if err == io.EOF {
			return
		}
		json.Unmarshal(buf[:n], &gotMap)

		if len(gotMap) == 0 {
			gotMap = nil
			continue
		}

		//最后12秒的时候,卡着时间点,3数据发过来了,需要处理
		select {
		case <-timeSaveBlock.C:
			conn.Read(buf)
			gotMapCopy := make([]string, len(gotMap))
			copy(gotMapCopy, gotMap)
			manage(gotMapCopy)
			gotMap = nil

			messMap = nil
			messMap = make(map[[16]byte]int, 0)
			fmt.Println("6秒一共多少数据", len(blockSave))
			s := blockSave[0] + blockSave[len(blockSave)-1]
			//fmt.Println("第一条加最后一条", s)
			sum := md5.Sum([]byte(s))
			fmt.Println("12s hash:", sum)
			blockSave = nil
			blockSave = make([]string, 0)
			messMap = make(map[[16]byte]int, 0)
		default:
			gotMapCopy := make([]string, len(gotMap))
			copy(gotMapCopy, gotMap)
			go manage(gotMapCopy)
			gotMap = nil
		}
		//fmt.Println("接收数据时间:", end - start)
		//fmt.Println("当前时间",nowtime)
		//fmt.Println("上次时间",timesave)

	}
}

//去重
func manage(gotMap []string) {
	var messSlcie = make([]string, 0, 60000)
	//去重
	//start := time.Now().UnixNano() / 1e6
	for _, v := range gotMap {
		sum := md5.Sum([]byte(v))
		if _, ok := messMap[sum]; ok {
			continue
		} else {
			messMap[sum] = 1
			messSlcie = append(messSlcie, v)
			blockSave = append(blockSave, v)
		}
	}
	//fmt.Println("一共多少数据", len(messSlcie))todo
}

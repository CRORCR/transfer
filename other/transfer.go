package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"
)

var writeInfo = make(chan []string)
var blockSave = make([]string, 0)
var messMap = make(map[[16]byte]int, 0)

func Server() {
	listen, err := net.Listen("tcp", ":9003")
	if err != nil {
		fmt.Println("端口占用")
	}

	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("connection is fail,err: ", err)
		return
	}
	timeSaveBlock = time.NewTicker(BLOCKTIME * time.Second)
	proess(conn)
}

var timesave int64

var timeSaveBlock *time.Ticker


func proess(conn net.Conn) {
	defer conn.Close()
	for {
		send := time.NewTicker(1 * time.Second)
		gotMap := make([]string, 0)
		//start := time.Now().UnixNano() / 1e6
		//decoder := json.NewDecoder(conn)
		//decoder.Decode(&gotMap)
		buf := make([]byte, 1024*1024*30)

		n, err := conn.Read(buf)
		if err == io.EOF {
			return
		}
		json.Unmarshal(buf[:n], &gotMap)

		//end := time.Now().UnixNano() / 1e6

		if len(gotMap) == 0 {
			gotMap = nil
			continue
		}

		select {
		case <-timeSaveBlock.C:
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
		case <-send.C:
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

	//s := gotMap[0] + gotMap[len(gotMap)-1]
	//sum := md5.Sum([]byte(s))
	//fmt.Println("测试每条数据hash:", sum)

	//去重完成,保存起来,等待12秒,存储
	writeInfo <- messSlcie
	//end := time.Now().UnixNano() / 1e6
	//fmt.Print("去重	", end-start)
	//fmt.Println("一共多少数据", len(messSlcie))todo
}

//var ti = time.NewTicker(time.Second * 120)

func Client() {
	conn := dialSer(ADDR_1)
	//defer func() {
	//	conn.Close()
	//	flag.Parse()
	//	if *cpuprofile != "" {
	//		f, err := os.Create(*cpuprofile)
	//		if err != nil {
	//			log.Fatal("could not create CPU profile: ", err)
	//		}
	//		if err := pprof.StartCPUProfile(f); err != nil {
	//			log.Fatal("could not start CPU profile: ", err)
	//		}
	//		defer pprof.StopCPUProfile()
	//	}
	//
	//	// ... rest of the program ...
	//
	//	if *memprofile != "" {
	//		f, err := os.Create(*memprofile)
	//		if err != nil {
	//			log.Fatal("could not create memory profile: ", err)
	//		}
	//		runtime.GC() // get up-to-date statistics
	//		if err := pprof.WriteHeapProfile(f); err != nil {
	//			log.Fatal("could not write memory profile: ", err)
	//		}
	//		f.Close()
	//	}
	//}()
	for {
		select {
		//case <-ti.C:
		//	return
		case addMessage := <-writeInfo:
			encoder := json.NewEncoder(conn)
			encoder.Encode(addMessage)
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
			return conn
		}
	}
}

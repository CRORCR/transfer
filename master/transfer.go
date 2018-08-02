package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

var blockchan = make(chan []string)

//func Client() {
//	var isFlag = true
//	var blockSave = make([]string, 0)
//	conn2 := dialSer(ADDR_3)
//	defer conn2.Close()
//	timeSaveBlock = time.NewTicker(BLOCKTIME * time.Second)
//	ti := time.NewTicker(time.Second * 7)
//	for {
//		select {
//		case <-ti.C:
//			return
//		case addMessage := <-chMess:
//			select {
//			case <-timeSaveBlock.C:
//				go saveLevelDB()
//				blockchan <- blockSave
//				fmt.Println("6秒一共多少数据", len(blockSave))
//				//计算hash
//				s := blockSave[0] + blockSave[len(blockSave)-1]
//				sum := md5.Sum([]byte(s))
//				fmt.Println("12s hash:", sum)
//				blockSave = make([]string, 0)
//			default:
//				json.NewEncoder(conn2).Encode(addMessage)
//				fmt.Println("主节点发送多少数据", len(addMessage))
//				blockSave = append(blockSave, addMessage...)
//				if isFlag {
//					timeSaveBlock = time.NewTicker(BLOCKTIME * time.Second)
//					isFlag = false
//				}
//			}
//		default:
//		}
//	}
//}

func Client() {
	var timeSaveBlock *time.Ticker
	var blockSave = make([]string, 0)
	//var isFlag = true
	conn := dialSer(ADDR_1)
	conn2 := dialSer(ADDR_2)
	defer conn.Close()
	defer conn2.Close()
	timeSaveBlock = time.NewTicker(BLOCKTIME * time.Second)
	send := time.NewTicker(1 * time.Second)
	//ti := time.NewTicker(time.Second * 24)
	for {
		select {
		//case <-ti.C:
		//	key := GetBlock()
		//	fmt.Println("多少个块",len(key))
		//	//获得交易详情
		//	page := GetPage(key[0],500,1000)
		//	fmt.Println("交易数量==500?",len(page))
		//
		//	//获得每个交易数量
		//	for _,v:=range key{
		//		num := GetKeyNum(v)
		//		fmt.Printf("区块:%v 交易数量:%v\n",v,num)
		//	}
		//	return
			//查看内存
			//f,err:=os.OpenFile("./men.out",os.O_RDWR|os.O_CREATE,0644)
			//if err!=nil{
			//	log.Fatal("err",err)
			//}
			//pprof.WriteHeapProfile(f)
			//f.Close()

		case addMessage := <-chMess:
			select {
			case <-timeSaveBlock.C:
				go saveLevelDB()
				blockchan <- blockSave
				//fmt.Println("count", len(blockSave))
				//计算hash
				if len(blockSave)!=0{
					//s := blockSave[0] + blockSave[len(blockSave)-1]
					//sum := md5.Sum([]byte(s))
					//fmt.Println("hash:", sum)
					blockSave = make([]string, 0)
				}else{
					blockSave = make([]string, 0)
					timeSaveBlock=time.NewTicker(BLOCKTIME * time.Second)
				}

			case <-send.C:
				buf:=make([]byte,1024*1024*20)
				writer := bufio.NewWriter(conn)
				writer2 := bufio.NewWriter(conn2)
				buf,_=json.Marshal(addMessage)
				writer.Write(buf)
				writer2.Write(buf)
				writer.Flush()
				writer2.Flush()
				//json.NewEncoder(conn).Encode(addMessage)
				//json.NewEncoder(conn2).Encode(addMessage)
				//fmt.Println("send message", len(addMessage))
				intType := fmt.Sprintf("%v", time.Now().UnixNano())
				levelPut([]byte(intType),buf)
				blockSave = append(blockSave, intType)
				//if isFlag {
				//	timeSaveBlock = time.NewTicker(BLOCKTIME * time.Second)
				//	isFlag = false
				//}
			}
		default:
		}
	}
}

func saveLevelDB() {
	send := <-blockchan
	bytes, _ := json.Marshal(send)
	intType := fmt.Sprintf("%v", time.Now().UnixNano())
	levelPut([]byte(intType), bytes)
	SaveBlock(intType)
}

func dialSer(id string) (conn net.Conn) {
	for {
		conn, err := net.Dial("tcp", id)
		if err != nil {
			continue
		} else {
			return conn
		}
	}
}


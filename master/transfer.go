package main

import (
	"bufio"
	"crypto/md5"
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
var blockNum = make([]int, 0)
var deleteBlockList = make([]string,0)

func Client() {
	var timeSaveBlock *time.Ticker
	var blockSave = make([]string, 0)
	//var isFlag = true
	conn := dialSer(ADDR_1)
	conn2 := dialSer(ADDR_2)
	defer conn.Close()
	defer conn2.Close()
	timeSaveBlock = time.NewTicker(BLOCKTIME * time.Second)
	i := 1 * time.Second
	send := time.NewTicker(i)
	//ti := time.NewTicker(time.Second * 13)
	for {
		select {
		//case <-ti.C:
		//	key := GetBlock()
		//	fmt.Println("块详情",key)
		//	fmt.Println("多少个块",len(key))
			//获得交易详情
		//	page2 := GetPage(key[1],500,501)
		//	fmt.Printf("第2个区块的第500个交易详细信息:%+v\n",page2)

			//获得每个交易数量
		//	for _,v:=range key{
				//获得每个块的交易全集 int
		//		num := GetKeyNum(v)
		//		fmt.Printf("区块:%v 交易数量:%v\n",v,num)
		//	}
			//第1个区块多少交易记录
		//	fmt.Println(GetBlockForNum(key[0]+"blockNum"))
			//第2个区块多少交易记录
		//	fmt.Println(GetBlockForNum(key[1]+"blockNum"))
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
				fmt.Println("count", len(blockSave))
				//计算hash
				if len(blockSave) != 0 {
					s := blockSave[0] + blockSave[len(blockSave)-1]
					sum := md5.Sum([]byte(s))
					fmt.Println("hash:", sum)
					blockSave = make([]string, 0)
				} else {
					blockSave = make([]string, 0)
					timeSaveBlock = time.NewTicker(BLOCKTIME * time.Second)
				}
			case <-send.C:
				//default:
				buf := make([]byte, 1024*1024*30)
				writer := bufio.NewWriter(conn)
				writer2 := bufio.NewWriter(conn2)
				buf, _ = json.Marshal(addMessage)
				writer.Write(buf)
				writer2.Write(buf)
				writer.Flush()
				writer2.Flush()
				//json.NewEncoder(conn).Encode(addMessage)
				//json.NewEncoder(conn2).Encode(addMessage)
				//fmt.Println("发送message", len(addMessage))
				intType := fmt.Sprintf("%v", time.Now().UnixNano())
				levelPut([]byte(intType), buf)
				blockSave = append(blockSave, intType)
				blockNum = append(blockNum, len(addMessage))
			}
		default:
		}
	}
}

func saveLevelDB() {
	send := <-blockchan
	//12个时间戳
	bytes, _ := json.Marshal(send)
	intType := fmt.Sprintf("%v", time.Now().UnixNano())
	levelPut([]byte(intType), bytes)
	SaveBlock(intType)
	SaveBlockNum(intType+"blockNum", blockNum)
	blockNum = make([]int, 0)

	deleteBlockList=append(deleteBlockList,intType)
	if len(deleteBlockList)==20{
		deleteFlag:=deleteBlockList[0]
		deleteLevelDB(deleteFlag)
		deleteLevelNum(deleteFlag+"blockNum")
		deleteBlockList=deleteBlockList[1:]
	}
	//fmt.Println("数据库容量",len(deleteBlockList))
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

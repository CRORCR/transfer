package main

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net"
	"net/rpc"
	"sync/atomic"
	"time"
)


//服务端监听自己端口
func Server() {
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
		go proess(conn)
	}
}

func proess(conn net.Conn) {
	for {
		gotMap := make([]string, 0)
		start:=time.Now().UnixNano()/1e6
		json.NewDecoder(conn).Decode(&gotMap)
		end:=time.Now().UnixNano()/1e6
		if (end-start)==0{
			return
		}
		fmt.Println("处理数据时间:",end-start)
		if len(gotMap) != 0 {
			conn.Close()
			levelDB.manage(gotMap)
			gotMap = make([]string, 0)
			if len(levelDBSlice) >= trainNum {
				//打包时间 生成一个包   12w数据备份清空,map去重清空 只留下levelDB供webserver使用
				block=append(block,levelDBSlice)
				fmt.Println("已经有12w了")
				fmt.Printf("block:%+v\n",block)
				levelDB.MessMap=make(map[[16]byte]int,0)
				//fmt.Println("然后map置空,下一个12秒,再次计:",len(levelDB.MessMap))
				sendMess <- "ok"
				//下次重新计数,打包/发送之前,用于存储数据库 作为key值,方便取出
				return
			}
		}else{return}
	}
}


//存入数据库的数据,需要排序,使用自增为key,达到12秒(12万)数据,再取出来 打包,存入数据库
var levelDBSlice=make([]string,0)

//test 打印重复
var count int32=1
func (level LevelDB) manage(gotMap []string) {
	level.lock.Lock()
	//去重
	start := time.Now().UnixNano() / 1e6
	for _, v := range gotMap {
		sum := md5.Sum([]byte(v))
		levelDB.lock.Lock()
		//fmt.Printf("收到的数据:%s\n", v)
		if _, ok := levelDB.MessMap[sum]; ok {
			atomic.AddInt32(&count,1)
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
	if count!=1{
		fmt.Println("重复记录:",atomic.LoadInt32(&count))
		atomic.CompareAndSwapInt32(&count,120000,1)
		atomic.CompareAndSwapInt32(&count,120001,1)
		levelDB.MessSlcie=make([]string,0)
		//fmt.Println("levelDB.MessSlcie应该空的:",len(levelDB.MessSlcie))
		return
	}
	//if count!=0{
	//	fmt.Printf("如果重复,序列化是什么,byte:%v err:%v \n",bytes, err)
	//}
	intType := fmt.Sprintf("%v",time.Now().UnixNano())
	//levelPut([]byte(intType), bytes)
	levelDB.MessSlcie = make([]string, 0)
	//0-12w个消息 存入slice中,打包的时候,需要情况这个切片
	levelDBSlice=append(levelDBSlice,intType)
	return
}

func Client() {
	//主节点需要往从节点都发送数据
	for _, v := range addrList {
		k := len(levelDB.MessSlcie) / recvMessageNum
		conn := dialSer(v)
		for i := 0; i < k; i++ {
			sendMess := levelDB.MessSlcie[:recvMessageNum]
			levelDB.MessSlcie = levelDB.MessSlcie[recvMessageNum:]
			err := json.NewEncoder(conn).Encode(&sendMess)
			if err != nil {
				fmt.Println("序列化失败")
				return
			}
		}
		conn.Close()
	}
	return
}

func sendPack(ip string){
	conn, err := net.Dial("tcp", ip)
	defer conn.Close()
	if err!=nil{
		fmt.Println("连接失败")
		return
	}
	writer := bufio.NewWriter(conn)
	writer.WriteString("flying")

}
func dialSer(ip string) (conn net.Conn) {
	for {
		conn, err := net.Dial("tcp", ip)
		if err != nil {
			continue
		} else {
			fmt.Println("成功建立连接")
			return conn
		}
	}
}

func callOther() {
	client, err := DialRpc()
	var reply string
	for {
		err = client.Call("GetStart", 1, &reply)
		if err != nil {
			fmt.Println("调用远程服务失败", err)
			return
		}
		fmt.Println("远程服务返回结果：", reply)
	}
}

func GetStart(arg int, result *string) error {
	//收到消息开始计时,12秒后打包
	*result = "ok"
	return nil
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
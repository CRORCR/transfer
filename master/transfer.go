package main

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net"
	"net/rpc"
)

//服务端监听自己端口
func ServerListen() {
	listen, err := net.Listen("tcp", ":9002")
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
	defer conn.Close()
	for {
		var gotMap = make([]string, recvMessageNum)
		json.NewDecoder(conn).Decode(&gotMap)
		manage(gotMap)
		if len(levelDB.MessSlcie) == train {
			//fmt.Println("容器长度:", len(levelDB.MessSlcie))
			return
		}
	}
}

func manage(gotMap []string) {
	//去重
	for _, v := range gotMap {
		sum := md5.Sum([]byte(v))
		if _, ok := levelDB.MessMap[sum]; ok {
			//fmt.Println("存在数据", vv)
			return
		} else {
			levelDB.MessMap[sum] = 1
			levelDB.MessSlcie = append(levelDB.MessSlcie, v)
		}
	}
}

func Client() {

	conn := dialSer(addrList[0])
	conn2 := dialSer(addrList[1])
	defer conn.Close()
	defer conn2.Close()
	for {
		select {
		case addMessage := <-chMess:
			json.NewEncoder(conn).Encode(addMessage)
			json.NewEncoder(conn2).Encode(addMessage)
		default:
		}
	}
}

func sendPack(ip string) {
	conn, err := net.Dial("tcp", ip)
	defer conn.Close()
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	writer := bufio.NewWriter(conn)
	writer.WriteString("flying")
}

func dialSer(id string) (conn net.Conn) {
	for {
		conn, err := net.Dial("tcp", id)
		if err != nil {
			continue
		} else {
			//fmt.Println("成功建立连接")
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

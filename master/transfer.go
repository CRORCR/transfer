package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

var blockchan = make(chan []string)

var deleteBlockList = make([]string,0)

func Server() {
	listen, err := net.Listen("tcp", ":9003")
	if err != nil {
		panic(err)
	}

	conn, err := listen.Accept()
	if err != nil {
		return
	}
	proess(conn)
}
func proess(conn net.Conn){
	defer conn.Close()
	for {
		gotMap := make([]string, 0)
		json.NewDecoder(conn).Decode(&gotMap)
	}
}

func Client() {
	var timeSaveBlock,send *time.Ticker
	send = time.NewTicker(1 * time.Second)
	timeSaveBlock = time.NewTicker(6 * time.Second)
	var blockSave = make([]string, 0)
	conn := dialSer(ADDR_1)
	defer conn.Close()

	ti := time.NewTicker(time.Second * 30)
	for {
		select {
		case addMessage := <-chMess:
			select {
			case <-ti.C:
				return
			case <-timeSaveBlock.C:
				go saveLevelDB()
				blockchan <- blockSave
				fmt.Println("count", len(blockSave))
				blockSave = make([]string, 0)
			case <-send.C:
				json.NewEncoder(conn).Encode(addMessage)
				blockSave = append(blockSave, addMessage...)
			}
		default:
		}
	}
}

func saveLevelDB() {
	<-blockchan
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

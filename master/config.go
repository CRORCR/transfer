package main

import (
	"fmt"
	"strconv"
)

type Config struct {
	BlockTime        string
	SendNum        int
}

var config = &Config{}

func initConf() {
	conf, err := NewConfig("./app.conf")
	if err != nil {
		fmt.Println("err",err)
		return
	}
	config.BlockTime = conf.GetString("block_time")
	sendnum := conf.GetString("send_num")
	config.SendNum, _= strconv.Atoi(sendnum)
	//fmt.Printf(" config:%+v\n",config)
	return
}

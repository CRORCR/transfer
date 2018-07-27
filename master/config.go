package main

import (
	"fmt"
	"strconv"
)

type Config struct {
	blockTime        string
	sendNum        int

}

var config = &Config{}

func initConf() {
	conf, err := NewConfig("./aa.txt")
	if err != nil {
		fmt.Println("err",err)
		return
	}
	blocktime := conf.GetString("block_time")
	sendnum := conf.GetString("send_num")
	fmt.Println("Sendum2",sendnum)
	value, _:= strconv.Atoi(sendnum)
	fmt.Println("value",value)

	config.blockTime=blocktime
	config.sendNum=value
	fmt.Printf(" config:%+v\n",config)
	return
}

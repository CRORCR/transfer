package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

func TestMaster(t *testing.T) {
	initConf()
	go getMessage()
	Client()
}

var count2 int32
func TestRand(t *testing.T) {
	var str = []byte("abcdefghighlmnopqrstuvwxyzABCDEFGHIGHLMNOPQRSTUVWXYZ")
	var arr [95]uint8
	var timeStamp = time.Now().Unix()
	for i := 0; i < 95; i++ {
		i2 := rand.Intn(52)
		arr[i] = str[i2]
	}
	//from
	sprintf := fmt.Sprintf("TxHash:%v,TxReceiptStatus:success,Height:%v,TimeStamp:%v,From:%v,To:%v,Value:200,GasLimit:200,GasUsedByTxn:17,GasPrice:100,ActualTxCost:7,Nonce:2", string(arr[32:]),atomic.LoadInt32(&count2),timeStamp,string(arr[:16]), string(arr[16:32]))
	fmt.Println(sprintf)
	//to

	//hash
}

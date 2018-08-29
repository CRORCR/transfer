package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var chMess = make(chan []string)

func init() {
	rand.Seed(time.Now().Unix())
}

var unix  int64
func getMessage() {
	for {

		atomic.AddInt32(&count,1)
		q := rand.Intn(1000)
		num := 50000 + q
		var send = make([]string, num)
		for i := 0; i < num; i++ {
			s := random()
			send[i] = s
		}
		chMess <- send
		send=nil
	}
}

var str = []byte("abcdefghighlmnopqrstuvwxyzABCDEFGHIGHLMNOPQRSTUVWXYZ")
var count int32
func random() string {
	var arr string
	for i := 0; i < 95; i++ {
		i2 := rand.Intn(52)
		arr = arr + string(str[i2])
	}
	unix = time.Now().Unix()
	return fmt.Sprintf("TxHash %v,TxReceiptStatus success,Height %v,TimeStamp %v,From %v,To %v,Value 200,GasLimit 200,GasUsedByTxn 17,GasPrice 100,ActualTxCost 7,Nonce 2", string(arr[32:]),atomic.LoadInt32(&count),unix,string(arr[:16]), string(arr[16:32]))
}
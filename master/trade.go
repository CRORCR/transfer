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
		unix = time.Now().Unix()
		atomic.AddInt32(&count,1)
		q := rand.Intn(1000)
		num := 40000 + q
		var send = make([]string, num)
		for i := 0; i < num; i++ {
			s := random()
			send[i] = s
		}
		//fmt.Println("send first:",send[0])
		//fmt.Println("send last:",send[len(send)-1])
		chMess <- send
		send=nil
	}
}
//func random() (stri string) {
//	for i := 0; i < 203; i++ {
//		i2 := rand.Intn(52)
//		stri = stri + string(str[i2])
//	}
//	return stri
//}



var str = []byte("abcdefghighlmnopqrstuvwxyzABCDEFGHIGHLMNOPQRSTUVWXYZ")
var count int32
func random() string {
	var arr string
	for i := 0; i < 95; i++ {
		i2 := rand.Intn(52)
		arr = arr + string(str[i2])
	}
	return fmt.Sprintf("TxHash %v,TxReceiptStatus success,Height %v,TimeStamp %v,From %v,To %v,Value 200,GasLimit 200,GasUsedByTxn 17,GasPrice 100,ActualTxCost 7,Nonce 2", string(arr[32:]),atomic.LoadInt32(&count),unix,string(arr[:16]), string(arr[16:32]))
	//other节点会丢失数据,可能是生成时间戳耗时,我先注释
	//return fmt.Sprintf("TxHash:%v,TxReceiptStatus:success,Height:%v,TimeStamp:%v,From:%v,To:%v,Value:200,GasLimit:200,GasUsedByTxn:17,GasPrice:100,ActualTxCost:7,Nonce:2", string(arr[32:]),atomic.LoadInt32(&count),time.Now().Unix(),string(arr[:16]), string(arr[16:32]))
}
//func random() string {
//	var arr string
//	for i := 0; i < 95; i++ {
//		i2 := rand.Intn(52)
//		arr =arr+string(str[i2])
//	}
//	return fmt.Sprintf("TxHash-%v,TxReceiptStatus-success,Height-%v,From:%v,To-%v,Value-200,GasLimit-200,GasUsedByTxn-17,GasPrice-100,ActualTxCost-7,Nonce-2", string(arr[32:95]),atomic.LoadInt32(&count),string(arr[:16]), string(arr[16:32]))
//	//other节点会丢失数据,可能是生成时间戳耗时,我先注释
//	//return fmt.Sprintf("TxHash:%v,TxReceiptStatus:success,Height:%v,TimeStamp:%v,From:%v,To:%v,Value:200,GasLimit:200,GasUsedByTxn:17,GasPrice:100,ActualTxCost:7,Nonce:2", string(arr[32:]),atomic.LoadInt32(&count),time.Now().Unix(),string(arr[:16]), string(arr[16:32]))
//}

//var arr [95]uint8
//var timeStamp = time.Now().Unix()
//for i := 0; i < 95; i++ {
//i2 := rand.Intn(52)
//arr[i] = str[i2]
//}
////from
//sprintf := fmt.Sprintf("TxHash:%v,TxReceiptStatus:success,Height:%v,TimeStamp:%v,From:%v,To:%v,Value:200,GasLimit:200,GasUsedByTxn:17,GasPrice:100,ActualTxCost:7,Nonce:2", string(arr[32:]),atomic.LoadInt32(&count),timeStamp,string(arr[:16]), string(arr[16:32]))
//fmt.Println(sprintf)
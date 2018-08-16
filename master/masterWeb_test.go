package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	"github.com/syndtr/goleveldb/leveldb"
)

var count2 int32



func TestRand(t *testing.T) {
	var str = []byte("abcdefghighlmnopqrstuvwxyzABCDEFGHIGHLMNOPQRSTUVWXYZ")
	var arr [95]uint8
	var timeStamp = time.Now().Unix()
	for i := 0; i < 95; i++ {
		arr[i] = str[rand.Intn(52)]
	}
	sp := fmt.Sprintf("TxHash:%v,TxReceiptStatus:success,Height:%v,TimeStamp:%v,From:%v,To:%v,Value:200," +
		"GasLimit:200,GasUsedByTxn:17,GasPrice:100,ActualTxCost:7,Nonce:2", string(arr[32:]), atomic.LoadInt32(&count2),
		timeStamp, string(arr[:16]), string(arr[16:32]))
	assert.Equal(t, len(sp), 241, "got trading error")
}

func init() {
	db, _ = leveldb.OpenFile("./dbtest", nil)
}
func TestFindBlockNum(t *testing.T) {
	key := GetBlock()
	//t.Log(key)
	assert.Equal(t, len(key), 2, "got block num error")
}

func TestBlockTr(t *testing.T) {
	key := GetBlock()
	for _, v := range key {
		num := GetKeyNum(v)
		t.Logf("block:%v num:%v\n", v, num)
	}
	t.Log(GetBlockForNum(key[0] + "blockNum"))
	assert.Equal(t, len(GetBlockForNum(key[0] + "blockNum")), 6, "got blockNum error")
	t.Log(GetBlockForNum(key[1] + "blockNum"))
	assert.Equal(t, len(GetBlockForNum(key[0] + "blockNum")), 6, "got blockNum error")
}

func TestBlockInfo(t *testing.T) {
	key := GetBlock()
	assert.Equal(t, len(key), 2, "got block num error")
	page := GetPage(key[1], 500, 501)
	assert.Equal(t, len(page[0]), 241, "got trading error")
}


func TestWriteFile(t *testing.T){
	file, err := os.OpenFile("G:/block.txt", os.O_APPEND|os.O_CREATE, 0644)
	if err!=nil{
		fmt.Println("打开文件失败")
		return
	}
	key := GetBlock()
	for _,v:=range key{
		var arr []string
		getKey := GetKey(v)
		json.Unmarshal(getKey,&arr)
		for _,vv:=range arr{
			var arr2 []string
			getKey := GetKey(vv)
			json.Unmarshal(getKey,&arr2)
			for _,v3:=range arr2{
				file.WriteString(v3+"\n")
			}
		}
	}

	//GetBlockForNum(key[1])
}
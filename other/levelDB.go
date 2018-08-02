package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/syndtr/goleveldb/leveldb"

)

var db *leveldb.DB



func init() {
	err := os.RemoveAll("./db")
	if err!=nil{
		panic("删除db失败")
	}
	db, _ = leveldb.OpenFile("./db", nil)
}

func levelPut(key []byte, val []byte) {
	db.Put(key, val, nil)
}

func GetKey(key string) []byte {
	ids, err := db.Get([]byte(key), nil)
	if err != nil {
		panic(err)
	}
	return ids
}

func SaveBlock(key string)error {
	getKey := GetKey("block")
	if getKey==nil{
		bytes, _ := json.Marshal(key)
		levelPut([]byte("block"),bytes)
	}
	block:=make([]string, 0)
	json.Unmarshal(getKey,&block)
	block=append(block,key)
	bytes, _ := json.Marshal(block)
	levelPut([]byte("block"),bytes)
	return nil
}

func GetBlock()(key []string){
	ids, err := db.Get([]byte("block"), nil)
	if err!=nil {
		return nil
	}
	block:=make([]string, 0)
	json.Unmarshal(ids,&block)
	return block
}

//根据对应的区块获得区块详细交易数据
func GetPage(blockKey string,start,end int)([]string){
	ids, err := db.Get([]byte(blockKey), nil)
	if err != nil {
		panic(err)
	}
	var blockInfo=make([]string,0)
	err = json.Unmarshal(ids,&blockInfo)
	if err != nil {
		panic(err)
	}

	if start<0 || len(blockInfo)<end{
		panic(fmt.Sprintf("获取数据小于0 或 大于最大交易数,最大交易数为:%d", len(ids)))
	}

	return blockInfo[start:end]
}
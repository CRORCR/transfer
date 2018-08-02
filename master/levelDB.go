package main

import (

	"encoding/json"
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

func closeLevelDB(){
	db.Close()
}

func levelPut(key []byte, val []byte) {
	db.Put(key, val, nil)
}

func GetKey(key string) []byte {
	ids, err := db.Get([]byte(key), nil)
	if err != nil {
		return nil
	}
	return ids
}

//func GetPage(blockKey string,start,end int)([]string){
//	ids, err := db.Get([]byte(blockKey), nil)
//	if err != nil {
//		panic(err)
//	}
//	var blockInfo=make([]string,0)
//	err = json.Unmarshal(ids,&blockInfo)
//	var block =make([]string,0)
//	var blockList =make([]string,0)
//	for _,v:=range blockInfo{
//		ids, _ := db.Get([]byte(v), nil)
//		json.Unmarshal(ids,&block)
//		blockList=append(blockList,block...)
//	}
//	if err != nil {
//		panic(err)
//	}
//	return blockList[start:end]
//}

func GetKeyNum(key string)int {
	ids, err := db.Get([]byte(key), nil)
	if err != nil {
		panic(err)
	}
	line:=make([]string,0)
	json.Unmarshal(ids, &line)
	line2:=make([]string,0)
	var length int
	for _,v:=range line{
		ids, _ := db.Get([]byte(v), nil)
		json.Unmarshal(ids, &line2)
		length=length+len(line2)
	}
	return  length
}
//保存block到数据库,每次保存都判断是否存储,如果不存在就直接存储,如果存在,取出,再存储
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
	var block =make([]string,0)
	var blockList =make([]string,0)
	for _,v:=range blockInfo{
		ids, _ := db.Get([]byte(v), nil)
		json.Unmarshal(ids,&block)
		blockList=append(blockList,block...)
	}
	if err != nil {
		panic(err)
	}
	return blockList[start:end]
}

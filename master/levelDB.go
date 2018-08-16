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

//记录数据
func SaveBlockNum(intType string,blockNum []int){
	bytes, _ := json.Marshal(blockNum)
	levelPut([]byte(intType),bytes)
}

//删除block的数据以及block里面的所有数据
func deleteLevelDB(blockKey string){
	ids, _ := db.Get([]byte(blockKey), nil)

	block:=make([]string, 0)
	json.Unmarshal(ids,&block)
	//删除所有的block里面的时间戳 对应的交易12个
	for _,v:=range block{
		db.Delete([]byte(v), nil)
	}
	db.Delete([]byte(blockKey), nil)
}

func deleteLevelNum(intType string){
	err := db.Delete([]byte(intType), nil)
	if err!=nil{
		fmt.Printf("err:%v\n",err)
	}
}

func GetBlock()(key []string){
	//所有的时间戳,这应该是个二维数组
	ids, err := db.Get([]byte("block"), nil)
	if err!=nil {
		return nil
	}
	block:=make([]string, 0)
	json.Unmarshal(ids,&block)
	//fmt.Printf("所有的时间戳,这应该是个二维数组:%+v\n",block)
	return block
}

//根据对应的区块获得区块详细交易数据
func GetPage(blockKey string,start,end int)([]string){
	ids, err := db.Get([]byte(blockKey), nil)
	if err != nil {
		return nil
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
		return nil
	}
	return blockList[start:end]
}

//对应区块一共多少数据
func GetBlockForNum(blockKey string)[]int {
	ids, _ := db.Get([]byte(blockKey), nil)
	var blockNum []int
	json.Unmarshal(ids,&blockNum)
	return blockNum
}
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

var file *os.File
var db *leveldb.DB

type LevelDB struct {
	//存储levelDB
	lock    sync.Mutex
	MessSlcie []string
	MessMap map[[16]byte]int
	//关闭
	//CloseChan chan int
	//存储所有key值(目前是所有的数据都发)
	//KeySlice []string
}

func NewLevelDB()*LevelDB{
	return &LevelDB{MessSlcie:make([]string,0),MessMap:make(map[[16]byte]int)}
}

func openLevelDB(){
	err := os.RemoveAll("./db")
	if err!=nil{
		panic("删除db失败")
	}
	db, _ = leveldb.OpenFile("./db", nil)
}
func closeLevelDB(){
	db.Close()
}

func levelPut(key []byte,val []byte) {
	db.Put(key,val, nil)
	levelDB.MessSlcie=make([]string,0)
}

//暂时不用,后期有N值再使用
//func (level LevelDB) GetKey(key string) string {
//	ids, err := db.Get([]byte(key), nil)
//	if err != nil {
//		panic(err)
//	}
//	return string(ids)
//}

//获得所有区块key的集合
func GetBlockKey() []string {
	return keySlice
}

//根据对应的区块获得区块数量,供分页显示
//func (level LevelDB) GetInfoForBlockKey(blockKey string)(allNum int){
//	ids, err := db.Get([]byte(blockKey), nil)
//	if err != nil {
//		panic(err)
//	}
//	var blockInfo=make([][]byte,0)
//	err = json.Unmarshal(ids,&blockInfo)
//	if err != nil {
//		panic(err)
//	}
//	return len(blockInfo)
//}

//根据对应的区块获得区块详细交易数据
func GetPage(blockKey []byte,start,end int)([]string){
	ids, err := db.Get(blockKey, nil)
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
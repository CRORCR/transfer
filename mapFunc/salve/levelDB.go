package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

type LevelDB struct {
	//存储levelDB
	//LogChan chan []byte
	MessSlcie []string
	MessMap   map[[16]byte]int
	lock      sync.RWMutex
	//存储所有key值(目前是所有的数据都发)
	//KeySlice []string
}

func NewLevelDB()*LevelDB{
	return &LevelDB{MessSlcie:make([]string,0),MessMap:make(map[[16]byte]int,0)}
}

func closeLevelDB() {
	db.Close()
}

func init() {
	err := os.RemoveAll("./db")
	if err!=nil{
		panic("删除db失败")
	}
	db, _ = leveldb.OpenFile("./db", nil)
}

func levelPut(key []byte, val []byte) {
	db.Put(key, val, nil)
	//if err!=nil{
		//fmt.Println(" 存入就是有问题的 err:",err)
	//}
	levelDB.MessSlcie = make([]string, 0)
	//fmt.Printf("存储了 %v %v \n", key, val)
}

func GetKey(key string)(ids []byte,err error) {
	ids, err = db.Get([]byte(key), nil)
	return ids,err
}

//获得所有区块key的集合
func GetBlockKey() [][]string {
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
package main

import (
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
//var key int

func levelPut(key []byte, val []byte) {
	db.Put(key, val, nil)
	levelDB.MessSlcie = make([]string, 0)
	//fmt.Printf("存储了 %v %v \n", key, val)
}

func GetKey(key string) []byte {
	ids, err := db.Get([]byte(key), nil)
	if err != nil {
		panic(err)
	}
	return ids
}

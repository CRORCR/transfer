package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

type LevelDB struct {
	//存储levelDB
	lock    sync.Mutex
	MessSlcie []string
	MessMap map[[16]byte]int
}

func NewLevelDB()*LevelDB{
	return &LevelDB{MessSlcie:make([]string,0),MessMap:make(map[[16]byte]int)}
}

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
	//if err!=nil{
	//fmt.Println(" 存入就是有问题的 err:",err)
	//}
	//levelDB.MessSlcie = make([]string, 0)
	//fmt.Printf("存储了 %v %v \n", key, val)
}

func GetKey(key string) []byte {
	ids, err := db.Get([]byte(key), nil)
	if err != nil {
		panic(err)
	}
	return ids
}

func GetKeyNum(key string)int {
	v,ok:=blockNumber[key]
	if !ok{
		return 0
	}
	return v
}
//func GetKeyNum(key string)int {
//	ids, err := db.Get([]byte(key), nil)
//	if err != nil {
//		panic(err)
//	}
//	var blockSlice = make([]string, 0)
//	json.Unmarshal(ids,&blockSlice)
//	return len(blockSlice)
//}

//获得所有区块key的集合
func GetBlockKey() []string {
	return block
}

//查过的数据
var searchPage=make([]string,0)
var fromPage = make([][]string,0)
//根据对应的区块获得区块详细交易数据
func GetPage(blockKey string,start,end int)([]string){
	//1.如果曾经查过的,直接返回
	info:=exitCheck(blockKey,start,end)
	if info!=nil{
		fmt.Println("之前查过,现在查询数据库")
		return info
	}
	//2.如果之前没有查过的,查出来,然后放入数组
	fmt.Println("之前没有查过,现在查询数据库")
	ids, err := db.Get([]byte(blockKey), nil)
	if err != nil {
		panic(err)
	}
	var blockInfo=make([]string,0)
	err = json.Unmarshal(ids,&blockInfo)
	if err != nil {
		panic(err)
	}
	searchPage=append(searchPage,blockKey)
	fromPage=append(fromPage,blockInfo[:300])
	return blockInfo[start:end]
}

func exitCheck(blockKey string,start,end int)([]string){
	for i,v:=range searchPage{
		if strings.Contains(v,blockKey){
			blockInfoM := fromPage[i]
			return blockInfoM[start:end]
		}
	}
	return nil
}
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers   []Server
	ServersID string
}

// // 定义返回json结构体
type BaseJsonBean struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

type Web1c struct {
	Flag         int `json:"flag"`
	Bdisplayfrom int `json: "bdisplayfrom"`
	Bdisplayto   int `json: "bdisplayto"`
}

type Web2c struct {
	Flag         int
	Nblock       int
	Tdisplayto   int
	Tdisplayfrom int
}

type Web4c struct {
	Flag         int
	Blocknum     int
	TDisplayfrom int
	TDisplayto   int
}

type StTxs struct {
	index     int32
	stamptime uint64
}
type web2s struct {
	StTxs []StTxs
}

type StTB struct {
	Blocknum int
	Txs      int
}

type web3s struct {
	BlockTxs []StTB
}

type txcs struct {
	Blocknum  int
	Txn       int
	Data      string
	Stamptime int64
}
type web4s struct {
	Total     int
	TxContent []txcs
}

var blockCount = make([]string, 0)
var blockTest []string // = make([][]string, 0)

func Cutstring(str, substri, substrii string) string {

	pos1 := strings.Index(str, substri)
	pos1 += len(substri)
	pos2 := strings.Index(str, substrii)
	des := str[pos1:pos2]
	return des
}

type Txs struct {
	Blocknum int
	Txs      int
}

//web1 server
type Web1s struct {
	Totalblock int
	BlockTxs   []Txs
}

type txps struct {
	Timestamps string
	Count      int
}

type txpi struct {
	Timestamps int64
	Count      int
}
type Web2s struct {
	//	Totalblock int
	Tpl []txpi
}

var ii int

func modify() {
	var m *sync.Mutex
	m = new(sync.Mutex)
	ii = 300

	for {
		time.Sleep(12 * time.Second)
		m.Lock()
		ii += 1
		m.Unlock()
	}
}

func webServer() {
	//	go modify()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	r.ParseForm() //解析参数，默认是不会解析的
	//	fmt.Fprintf(w, "Hello ! %s", html.EscapeString(r.URL.Path[1:]))
	if r.Method == "GET" {
		fmt.Println("method:", r.Method) //获取请求的方法

		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])

		for k, v := range r.Form {
			fmt.Print("key:", k, "; ")
			fmt.Println("val:", strings.Join(v, ""))
		}
	} else if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		var str = fmt.Sprintf("%s\n", result)
		//		index := strings.Index(str, "Flag")
		var Flag string
		if strings.Contains(str, "flag") {
			Flag = Cutstring(str, ":", ",")
		}
		// 去除空格
		Flag = strings.Replace(Flag, " ", "", -1)
		// 去除换行符
		Flag = strings.Replace(Flag, "\n", "", -1)

		intflag, err := strconv.Atoi(Flag)
		if err == nil {
			fmt.Println(intflag)
		}
		strings.Contains(str, "flag")

		if intflag == 1 {
			var f Web1c
			err := json.Unmarshal(result, &f)
			if err != nil {
				fmt.Println("***************************", err)
			}
			fmt.Println(f)

			//operation leveldb
			fp, open_err := os.OpenFile("json.txt", os.O_RDWR|os.O_CREATE, 0755)
			if open_err != nil {
				log.Fatal(err)
			}

			////////////////////////////////////////////////////
			var ress *Web1s
			ress = new(Web1s)

			//获得所有block的时间戳
			blockCount = GetBlockKey()
			//多少个区块
			ress.Totalblock = len(blockCount)
			fmt.Println("区块num",len(blockCount))
			var from, to int
			if f.Bdisplayfrom < 0 {
				from = 0
			} else {
				from = f.Bdisplayfrom
			}
			if f.Bdisplayto > len(blockCount) {
				to = len(blockCount)
			} else {
				to = f.Bdisplayto
			}
			//values := GetKey(blockCount[0])
			//lenblock := len(values)
			s := blockCount[0]
			keyNum := GetKeyNum(s)
			fmt.Println("***************",keyNum)

			for i := from - 1; i < to; i++ {
				s := blockCount[i]
				keyNum := GetKeyNum(s)
				fmt.Println("***************",keyNum)
				ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: i + 1, Txs: keyNum})
			}

			// // ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 20, Txs: 60000})
			// // ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 21, Txs: 60000})
			// // ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 23, Txs: 60000})
			// // ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 24, Txs: 60000})
			// // ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 25, Txs: 60000})
			// // ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 26, Txs: 60000})

			bytes, _ := json.Marshal(ress)

			if open_err := fp.Close(); open_err != nil {
				log.Fatal(open_err)
			}
			fmt.Fprint(fp, string(bytes))
			fp.Write(bytes)

			/////////////////////////////////////////////////////////////
			fp.Close()
			//返回response 的json数据
			//fmt.Println(bytes)
			fmt.Fprint(w, string(bytes))
		}

		if intflag == 2 {

			fp, open_err := os.OpenFile("json.txt", os.O_RDWR|os.O_CREATE, 0755)
			if open_err != nil {
				log.Fatal(err)
			}

			var f Web2c
			json.Unmarshal(result, &f)
			fmt.Println(f)

			//operation leveldb
			////////////////////////////////////////////////////////
			ccur := time.Now()
			ttimestamp := ccur.Unix()
			var ii int64
			var web2s *Web2s
			web2s = new(Web2s)
			//			web2s.Totalblock = 4
			// for {
			// 	ttimestamp -= 1
			// 	ttm := time.Unix(ttimestamp, 0)
			// 	fmt.Println(ttm.String())
			// 	//				ttms := ttm.String()
			// 	ii++
			// 	if ii == 60 {
			// 		break
			// 	}
			// 	web2s.Tpl = append(web2s.Tpl, txps{Timestamps: ttm.String(), Count: 200})

			// }

			//todo
			//blklist,err := GetPage(blk[0][0], 10, 20)
			//if err!=nil{
			//	blk = GetBlockKey()
			//	blklist,err = GetPage(blk[0][0], 10, 20)
			//}

			//var blockNum []int
			//for k,_:=range blockCount{
			//	blockNum=append(blockNum,k)
			//}
			//获得所有block
			blockCount = GetBlockKey()
			//第一个区块有多少数据
			keyNum := GetKeyNum(blockCount[0])
			for {
				ttimestamp -= 1
				ttm := time.Unix(ttimestamp, 0)
				fmt.Println(ttm.String())
				//				ttms := ttm.String()
				ii++
				if ii == 60 {
					break
				}
				web2s.Tpl = append(web2s.Tpl, txpi{Timestamps: ttimestamp, Count: keyNum})

			}
			bytes2, _ := json.Marshal(web2s)
			//			fmt.Fprint(fp, string(bytes2))
			len, errl := fp.Write(bytes2)
			if errl != nil {
				fmt.Println(errl, len)
			}

			//			time.Now().Format("2006-01-02 15:04:05")
			//			timestamp.Format("2006-01-02 15:04:05")

			bytes, _ := json.Marshal(web2s)
			//返回response 的json数据
			fmt.Println(string(bytes))
			fmt.Fprint(w, string(bytes))
		}

		var total int
		if intflag == 3 {
			var ff web3s

			//获得所有block
			blockCount = GetBlockKey()
			currentblock := len(blockCount)

			if currentblock-3 > 0 {
				keyNum := GetKeyNum(blockCount[currentblock - 4])
				total+=keyNum
				ff.BlockTxs = append(ff.BlockTxs, StTB{Blocknum: currentblock - 3, Txs: keyNum})
			}

			if currentblock-2 > 0 {
				keyNum := GetKeyNum(blockCount[currentblock - 3])
				total+=keyNum
				ff.BlockTxs = append(ff.BlockTxs, StTB{Blocknum: currentblock - 2, Txs: keyNum})
			}

			if currentblock-1 > 0 {
				keyNum := GetKeyNum(blockCount[currentblock - 2])
				total+=keyNum
				ff.BlockTxs = append(ff.BlockTxs, StTB{Blocknum: currentblock - 1, Txs: keyNum})
			}

			if currentblock > 0 {
				keyNum := GetKeyNum(blockCount[currentblock - 1])
				fmt.Println("第几个块",blockCount[currentblock - 1])
				fmt.Println("keynum是多少",keyNum)
				total+=keyNum
				ff.BlockTxs = append(ff.BlockTxs, StTB{Blocknum: currentblock, Txs: keyNum})
			}
			bytes, _ := json.Marshal(ff)
			//返回response 的json数据
			fmt.Println("total 是多少",total)
			//fmt.Println(string(bytes))
			fmt.Fprint(w, string(bytes))
		}

		if intflag == 4 {

			var f Web4c
			err := json.Unmarshal(result, &f)
			if err != nil {
				fmt.Println("***************************", err)
			}
			fmt.Println(f)

			json.Unmarshal(result, &f)
			fmt.Println(f)

			//operation leveldb
			// blockCount = GetBlockKey()
			// currentblock := len(blockCount)

			//获得所有block
			blockCount = GetBlockKey() //GetBlockKey()
			currentblock := len(blockCount)
			var blklist []string
			s := blockCount[0]
			blklist = GetPage(s, f.TDisplayfrom, f.TDisplayto+1)

			var i int
			var w4s *web4s
			w4s = new(web4s)
			w4s.Total = total
			for i = f.TDisplayfrom; i < f.TDisplayto+1; i++ {
				ccur := time.Now()
				ttimestamp := ccur.Unix()
				w4s.TxContent = append(w4s.TxContent, txcs{Blocknum: currentblock - (4 - f.Blocknum),
					Txn:       i,
					Data:      blklist[i-f.TDisplayfrom],
					Stamptime: ttimestamp})
			}

			bytes, _ := json.Marshal(w4s)
			//返回response 的json数据
			//fmt.Println(string(bytes))
			fmt.Fprint(w, string(bytes))

			// m.Lock()
			// for i = f.TDisplayfrom; i < f.TDisplayto; i++ {
			// 	stm := fmt.Sprintf("0x5aa027cdf5125197468d2cd15e%d", i)
			// 	ccur := time.Now()
			// 	ttimestamp := ccur.Unix()
			// 	w4s.TxContent = append(w4s.TxContent, txcs{Blocknum: ii + f.Blocknum - 1,
			// 		Txn:       i,
			// 		Data:      stm,
			// 		Stamptime: ttimestamp})
			// }
			// m.Unlock()
			// bytes, _ := json.Marshal(w4s)
			// //返回response 的json数据
			// fmt.Println(string(bytes))
			// fmt.Fprint(w, string(bytes))
		}

		if intflag == 5 {
			var f Web2c
			json.Unmarshal(result, &f)
			fmt.Println(f)

			//operation leveldb

			bytes, _ := json.Marshal(result)
			//返回response 的json数据
			fmt.Println(string(bytes))
			fmt.Fprint(w, string(bytes))
		}

	}
}

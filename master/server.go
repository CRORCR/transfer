package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"reflect"
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

type web2c struct {
	Flag   int
	Number int
}

type web5c struct {
	Flag   int
	Number int
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
//	Data      string
	Data     Transaction
	Stamptime int64
}
type web4s struct {
	Total     int
	TxContent []txcs
}

type blockstt struct {
	Stamptime int64
	Txs       int
}

type web5s struct {
	Bstt []blockstt
}

type statistc struct {
	Stamptime int64
	count     int
	total     uint64
}

type Transaction struct {
	TxHash          string
	TxReceiptStatus string
	Height          uint64
	TimeStamp       int64
	From            string
	To              string
	Value           uint64
	GasLimit        uint64
	GasUsedByTxn    uint64
	GasPrice        uint64
	ActualTxCost    uint64
	Nonce           uint64
	//	InputData       string
}


var blockCount = make([]string, 0)

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
	nums := list.New()
	var stc statistc
	stc.count = 20000
	stc.Stamptime = 200
	stc.total = 48000
	nums.PushBack(stc)
	nums.PushBack(stc)
	nums.PushBack(stc)

	fmt.Println(nums.Len())

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

			////////////////////////////////////////////////////
			var ress *Web1s
			ress = new(Web1s)

			//获得所有block
			blockCount = GetBlock() //GetBlock()
			ress.Totalblock = len(blockCount)

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


			for i := from - 1; i < to; i++ {
				s := blockCount[i]
				keyNum := GetKeyNum(s)
				fmt.Println("***************",keyNum)
				ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: i + 1, Txs: keyNum})
			}

			// ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 20, Txs: 60000})
			// ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 21, Txs: 60000})
			// ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 23, Txs: 60000})
			// ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 24, Txs: 60000})
			// ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 25, Txs: 60000})
			// ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 26, Txs: 60000})
			bytes, _ := json.Marshal(ress)
			fmt.Fprint(w, string(bytes))
		}

		if intflag == 2 {

			var f web2c
			json.Unmarshal(result, &f)
			fmt.Println(f)

			//operation leveldb
			////////////////////////////////////////////////////////
			ccur := time.Now()
			ttimestamp := ccur.Unix()
			var web2s *Web2s
			web2s = new(Web2s)

			//////////////////////////////////////////rand
			var intarr []int
			intarr = make([]int, 12)
			var intlen int

			var arrlen int
			var waverange int
//			var total int

			arrlen = 12
			waverange = 20000
//			total = 480000

			fuu := func(def int) {

				for intlen = 0; intlen < arrlen; intlen++ {
					intarr[intlen] = 0
				}

				var flag bool
				flag = true
				for intlen = 0; intlen < arrlen; intlen++ {
					//			rand.Seed(time.Now().Unix())
					//			rand.Seed(time.Now().UnixNano())

					tmp := rand.Intn(waverange)
					fmt.Println("*************************", tmp)
					if flag == true {
						intarr[intlen] = def/arrlen - tmp
						flag = false
					} else {
						intarr[intlen] = def/arrlen + tmp
						flag = true
					}
				}

				var sum int

				sum = 0
				for intlen = 0; intlen < arrlen; intlen++ {
					sum += intarr[intlen]
					if sum > def {
						sum -= intarr[intlen]
						intarr[intlen] = def - sum
						break
					}
					if sum < def && intlen == (arrlen-1) {
						tmp := rand.Intn(intlen)
						intarr[tmp] += def - sum
						//				intarr[intlen] = def - sum
						break
					}
				}
			}

			blocknum := f.Number / 12
			var curblocknum int64
			var timestamps int64
			var def int
			tmp := ttimestamp - int64(f.Number)
			blockCount := GetBlock()
			i := int64(len(blockCount))
			if i==0{
				web2s.Tpl = append(web2s.Tpl, txpi{Timestamps: tmp + timestamps + 12*curblocknum, Count: 0})
				bytes, _ := json.Marshal(web2s)
				fmt.Fprint(w, string(bytes))
				return
			}

			for curblocknum = 0; curblocknum < int64(blocknum); curblocknum++ {
				fmt.Println("blockCount:",i)
				if curblocknum>i{
					return
				}
				//第几个区块的12秒
				def = GetKeyNum(blockCount[curblocknum])
				fmt.Println("def:",def)
				fuu(def)
				for timestamps = 0; timestamps < 12; timestamps++ {

					web2s.Tpl = append(web2s.Tpl, txpi{Timestamps: tmp + timestamps + 12*curblocknum, Count: intarr[timestamps]})
				}
			}
//			web2s.Tpl = append(web2s.Tpl, txpi{Timestamps: tmp + timestamps + 12*curblocknum, Count: intarr[timestamps]})
			/*
				var ii int
					ii = 0
					for {
						ttimestamp -= 1
						ttm := time.Unix(ttimestamp, 0)
						fmt.Println(ttm.String())

						if ii == f.Number {
							break
						}
						web2s.Tpl = append(web2s.Tpl, txpi{Timestamps: ttimestamp, Count: 20000})
						ii++
					}
			*/
			bytes, _ := json.Marshal(web2s)
			//返回response 的json数据
			//fmt.Println(string(bytes))
			fmt.Fprint(w, string(bytes))
		}

		if intflag == 3 {
//			var ff web3s
			var ff Web1c
			json.Unmarshal(result, &ff)
			fmt.Println(ff)

			var ress *Web1s
			ress = new(Web1s)

			//获得所有block
			blockCount = GetBlock() //GetBlock()
			ress.Totalblock = len(blockCount)

			var from, to int
			if ff.Bdisplayfrom < 0 {
				from = 0
			} else {
				from = ff.Bdisplayfrom
			}
			if ff.Bdisplayto > len(blockCount) {
				to = len(blockCount)
			} else {
				to = ff.Bdisplayto
			}


			for i := from - 1; i < to; i++ {
				s := blockCount[i]
				keyNum := GetKeyNum(s)
				fmt.Println("***************",keyNum)
				ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: i + 1, Txs: keyNum})
			}

			// ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 20, Txs: 60000})
			// ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 21, Txs: 60000})
			// ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 23, Txs: 60000})
			// ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 24, Txs: 60000})
			// ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 25, Txs: 60000})
			// ress.BlockTxs = append(ress.BlockTxs, Txs{Blocknum: 26, Txs: 60000})
			bytes, _ := json.Marshal(ress)
			fmt.Fprint(w, string(bytes))

			//
			////获得所有block
			//blockCount = GetBlock()
			//currentblock := len(blockCount)
			//if currentblock-3 > 0 {
			//	keyNum := GetKeyNum(blockCount[currentblock - 4])
			//	ff.BlockTxs = append(ff.BlockTxs, StTB{Blocknum: currentblock - 3, Txs: keyNum})
			//}
			//
			//if currentblock-2 > 0 {
			//	keyNum := GetKeyNum(blockCount[currentblock - 3])
			//	ff.BlockTxs = append(ff.BlockTxs, StTB{Blocknum: currentblock - 2, Txs: keyNum})
			//}
			//
			//if currentblock-1 > 0 {
			//	keyNum := GetKeyNum(blockCount[currentblock - 2])
			//	ff.BlockTxs = append(ff.BlockTxs, StTB{Blocknum: currentblock - 1, Txs: keyNum})
			//}
			//
			//if currentblock > 0 {
			//	keyNum := GetKeyNum(blockCount[currentblock - 1])
			//	ff.BlockTxs = append(ff.BlockTxs, StTB{Blocknum: currentblock, Txs: keyNum})
			//}
			//
			//bytes, _ := json.Marshal(ff)
			////返回response 的json数据
			////fmt.Println(string(bytes))
			//fmt.Fprint(w, string(bytes))
		}

		if intflag == 4 {

			////////////////////////////////////////////////////////////////////////////
			var f Web4c
			err := json.Unmarshal(result, &f)
			if err != nil {
				fmt.Println("***************************", err)
			}
			fmt.Println(f)

			json.Unmarshal(result, &f)
			fmt.Println(f)

			//获得所有block
			blockCount = GetBlock()
			currentblock := len(blockCount)
			var blklist []string
			fmt.Println(currentblock)
//			s := blockCount[currentblock-(4 - f.Blocknum)-1]
			s := blockCount[f.Blocknum-1]
			blklist= GetPage(s, f.TDisplayfrom, f.TDisplayto+1)

			var i int
			var w4s *web4s
			w4s = new(web4s)
			w4s.Total = 120000


			for i = f.TDisplayfrom; i < f.TDisplayto+1; i++ {
				ccur := time.Now()
				ttimestamp := ccur.Unix()
//				w4s.TxContent = append(w4s.TxContent, txcs{Blocknum: currentblock-(4 - f.Blocknum)-1,
				w4s.TxContent = append(w4s.TxContent, txcs{Blocknum: f.Blocknum,
					Txn:       i,
					Data:      StrParsing(&Transaction{}, blklist[i-f.TDisplayfrom]),//blklist[i-f.TDisplayfrom],
					Stamptime: ttimestamp})
//				fmt.Println("**************************************",blklist[i-f.TDisplayfrom])
//				fmt.Println("**************************************Data",w4s.TxContent[i].Data)
			}

			fmt.Println("**************************************w4s",w4s)

			bytes, _ := json.Marshal(w4s)
			//返回response 的json数据
			//fmt.Println(string(bytes))
			fmt.Fprint(w, string(bytes))

		}

		if intflag == 5 {
			var f web5c
			json.Unmarshal(result, &f)
			fmt.Println(f)

			//operation leveldb
			var ws Web2s
			ii = 0
			ccur := time.Now()
			ttimestamp := ccur.Unix()

			blockCount = GetBlock()
			currentblock := len(blockCount)

			for {

				ttimestamp -= 12
				//
				//tmp := rand.Intn(40000)
				//var count int
				//if ii%2 > 0 {
				//	count = 480000 - tmp
				//} else {
				//	count = 480000 + tmp
				//}
				//
				//				ws..Tpl = append(web2s.Tpl, txpi{Timestamps: ttimestamp, Count: 20000})
				if ii == f.Number {
					break
				}
				s := blockCount[currentblock-1-ii]
				keyNum := GetKeyNum(s)
				ws.Tpl = append(ws.Tpl, txpi{Timestamps: ttimestamp, Count: keyNum})

				ii++
			}

			bytes, _ := json.Marshal(ws)
			//返回response 的json数据
			//fmt.Println(string(bytes))
			fmt.Fprint(w, string(bytes))
		}
	}
}

func GetTagName(structName interface{}) []string {
	t := reflect.TypeOf(structName)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		tagName := t.Field(i).Name
		tags := strings.Split(string(t.Field(i).Tag), "\"")
		if len(tags) > 1 {
			tagName = tags[1]
		}
		result = append(result, tagName)
	}
	return result
}

func StrParsing(structName interface{}, src string) Transaction{//string{//[]byte {
	var tag []string
	tag = GetTagName(structName)

	var tmp []string
	var offset int
	for offset = 0; offset < len(src); {
		var item string
		//		fmt.Println(src)
		index := strings.Index(src[offset:], ",")
		if index == -1 {
			item = src[offset:len(src)]
			tmp = append(tmp, item)
			break
		}
		//		CharCpy(item, src[:index], index)
		item = src[offset : offset+index]
		tmp = append(tmp, item)
		//		fmt.Println(item)
		offset += index + 1
	}
	//	fmt.Println(tmp)

	var i int

	var tsct Transaction

	for i = 0; i < len(tag); i++ {
		index := strings.Index(tmp[i], " ")
		//		fmt.Println(tmp[i][index+1:])
		if i == 0 {
			tsct.TxHash = tmp[i][index+1:]
		}
		if i == 1 {
			tsct.TxReceiptStatus = tmp[i][index+1:]
		}
		if i == 4 {
			tsct.From = tmp[i][index+1:]
		}
		if i == 5 {
			tsct.To = tmp[i][index+1:]
		}
		if i == 3 {
			value, err := strconv.ParseInt(tmp[i][index+1:], 10, 64)
			if err == nil {
				tsct.TimeStamp = value
			}
		}
		if i == 2 {
			value, err := strconv.ParseInt(tmp[i][index+1:], 10, 64)
			if err == nil {
				tsct.Height = uint64(value)
			}
		}
		if i == 6 {
			value, err := strconv.ParseInt(tmp[i][index+1:], 10, 64)
			if err == nil {
				tsct.Value = uint64(value)
			}
		}
		if i == 7 {
			value, err := strconv.ParseInt(tmp[i][index+1:], 10, 64)
			if err == nil {
				tsct.GasLimit = uint64(value)
			}
		}
		if i == 8 {
			value, err := strconv.ParseInt(tmp[i][index+1:], 10, 64)
			if err == nil {
				tsct.GasUsedByTxn = uint64(value)
			}
		}
		if i == 9 {
			value, err := strconv.ParseInt(tmp[i][index+1:], 10, 64)
			if err == nil {
				tsct.GasPrice = uint64(value)
			}
		}
		if i == 10 {
			value, err := strconv.ParseInt(tmp[i][index+1:], 10, 64)
			if err == nil {
				tsct.ActualTxCost = uint64(value)
			}
		}
		if i == 11 {
			value, err := strconv.ParseInt(tmp[i][index+1:], 10, 64)
			if err == nil {
				tsct.Nonce = uint64(value)
			}
		}
	}

//	bytes, _ := json.Marshal(tsct)

	//fp, open_err := os.OpenFile("json.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	//if open_err != nil {
	//	log.Fatal(open_err)
	//}
	//fwlen, fwerr := fp.Write(bytes)
	//if fwerr != nil {
	//	fmt.Println(fwerr, fwlen)
	//}
	//fp.Close()
	//return bytes
//	fmt.Println("************************************StringParsing Data", string(bytes[:]))
//	return string(bytes[:])
	return  tsct
}

type TransactionJson struct {
	TxHash          string `json:"TxHash"`
	TxReceiptStatus string `json:"TxReceiptStatus"`
	Height          uint64 `json:"Block Height"`
	TimeStamp       int64  `json:"TimeStamp"`
	From            string `json:"From"`
	To              string `json:"To"`
	Value           uint64 `json:"Value"`
	GasLimit        uint64 `json:"Gas Limit"`
	GasUsedByTxn    uint64 `json:"Gas Used By Txn"`
	GasPrice        uint64 `json:"Gas Price"`
	ActualTxCost    uint64 `json:"Actual Tx Cost/Fee"`
	Nonce           uint64 `json:"Nonce & {Position}"`
	InputData       string `json:"Input Data"`
}


//测试json数据
/*
if intflag == 4 {

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

	for {
		ttimestamp -= 1
		ttm := time.Unix(ttimestamp, 0)
		fmt.Println(ttm.String())
		//				ttms := ttm.String()
		ii++
		if ii == 60 {
			break
		}
		web2s.Tpl = append(web2s.Tpl, txpi{Timestamps: ttimestamp, Count: 20000})

	}
	bytes2, _ := json.Marshal(web2s)
	fmt.Println(string(bytes2))
	// len, errl := fp.Write(bytes2)
	// if errl != nil {
	// 	fmt.Println(errl, len)
	// }

	var ws web5s
	ii = 0
	for {
		ttimestamp -= 12
		ttm := time.Unix(ttimestamp, 0)
		fmt.Println(ttm.String())
		//				ttms := ttm.String()
		ii++
		if ii == 60 {
			break
		}
		//				ws..Tpl = append(web2s.Tpl, txpi{Timestamps: ttimestamp, Count: 20000})
		ws.Bstt = append(ws.Bstt, blockstt{Stamptime: ttimestamp, Txs: 120000})
	}

	bytes5, _ := json.Marshal(ws)
	//			fmt.Fprint(fp, string(bytes2))
	len5, err5 := fp.Write(bytes5)
	if err5 != nil {
		fmt.Println(err5, len5)
	}

	fp.Close()

}*/

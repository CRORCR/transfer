package main

import (
	"fmt"
	"strings"
	"time"
	"unsafe"
)

func main() {
	/*TestJoin()
	TestPrintf()
	TestBuilderOne()
	TestBuilderTwo()*/
	TestByte()
}

//+
func TestJoin()  {
	start:=time.Now().UnixNano()/1e6
	var str string
	for i:=0;i<100000;i++{
		str=str+"hello world"
	}
	end:=time.Now().UnixNano()/1e6
	fmt.Println("+ time:",end-start)//11000毫秒
}

func TestPrintf()  {
	start:=time.Now().UnixNano()/1e6
	var str string
	for i:=0;i<100000;i++{
		fmt.Sprintf("%s%s",str,"hello world")
	}
	end:=time.Now().UnixNano()/1e6
	fmt.Println("printf time:",end-start)//12毫秒
}

func TestBuilderTwo(){
	start:=time.Now().UnixNano()/1e6
	var str strings.Builder
	for i:=0;i<100000;i++{
		fmt.Fprint(&str,"hello world")
	}
	//fmt.Println(str.String())
	end:=time.Now().UnixNano()/1e6
	fmt.Println("builder fprint time:",end-start)//9毫秒
}

func TestBuilderOne(){
	start:=time.Now().UnixNano()/1e6
	var s strings.Builder
	for i:=0;i<100000;i++{
		s.WriteString("hello world")
	}
	end:=time.Now().UnixNano()/1e6
	fmt.Println("builder writer time:",end-start)//1毫秒
}

func TestByte(){
	var buf  []byte
	str:="hello world"
	buf=append(buf,str...)
	fmt.Printf("%+v\n",buf)
	fmt.Printf("%s",buf)//hello world
}
//比较string和byte数组是否相等,这种方式避免了内存分配
func unsafeEqual(a string, b []byte) bool {
	bbp := *(*string)(unsafe.Pointer(&b))
	return a == bbp
}

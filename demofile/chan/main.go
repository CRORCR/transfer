package main

import (
	"fmt"
	"time"
)

var block = make([][]string,0)
var t =time.NewTicker(5*time.Second)
func main() {
	for{
		select {
		case <-t.C:
			fmt.Printf("block是啥:%v\n",block)
		panic("hhh")
		default:
			go demo1()
			go demo2()

		}
	}
}
//存储
func demo1(){
	bb:=make([]string,0)
	for{
		bb=append(bb,"1")
		if len(bb)==12{
			block=append(block,bb)
			return
		}
	}
}

func demo2(){
	bb:=make([]string,0)
	var k int
	var v []string
	for k,v=range block{
		if len(v)==12{
			bb=v
		}
	}
	for _,v:=range bb{
		fmt.Print(v)
	}
	fmt.Println(k)
	block=append(block[:k],block[k+1:]...)
	fmt.Println()
}
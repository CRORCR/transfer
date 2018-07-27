package main

import "fmt"

func main() {
	go getMessage()
	for{
		select {
		case v := <-chMess:
			//go SendPeer()
			go Client(v)
		default:
		}
	}
}

func Client(addMessage []string) {
	//fmt.Printf("%v\n",addMessage)
	fmt.Println("是啥啊",len(addMessage))
}
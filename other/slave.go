package main

const (
	//连接主节点
	ADDR_1 = "localhost:9004"
	//ADDR_1         = "192.168.1.4:9004"
	aa = iota
)

//var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
//var memprofile = flag.String("memprofile", "", "write memory profile to `file`")


func main() {
	go Server()
	Client()
}
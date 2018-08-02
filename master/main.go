package main

import (
	_ "runtime/pprof"
)

const (
	//从节点ip
	ADDR_1,ADDR_2 = "localhost:9003","localhost:9004"
	//ADDR_1 = "localhost:9004"
	//ADDR_3 = "localhost:9003"
	//ADDR_1,ADDR_2 = "192.168.1.3:9003","192.168.1.4:9004"
	 BLOCKTIME=6
)
//var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
//var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {

	//f,err:=os.OpenFile("./tmp/cpu/prof",os.O_RDWR|os.O_CREATE,0644)
	//if err!=nil{
	//	log.Fatal("err",err)
	//}
	//defer f.Close()
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()
	//pprof.StopCPUProfile()
	//f.Close()
	//flag.Parse()
	//if *cpuprofile != "" {
	//	f, err := os.Create(*cpuprofile)
	//	if err != nil {
	//		log.Fatal("could not create CPU profile: ", err)
	//	}
	//	if err := pprof.StartCPUProfile(f); err != nil {
	//		log.Fatal("could not start CPU profile: ", err)
	//	}
	//	defer pprof.StopCPUProfile()
	//}
	//
	//// ... rest of the program ...
	//
	//if *memprofile != "" {
	//	f, err := os.Create(*memprofile)
	//	if err != nil {
	//		log.Fatal("could not create memory profile: ", err)
	//	}
	//	runtime.GC() // get up-to-date statistics
	//	if err := pprof.WriteHeapProfile(f); err != nil {
	//		log.Fatal("could not write memory profile: ", err)
	//	}
	//	f.Close()
	//}

	initConf()
	go getMessage()
	go webServer()
	Client()
}

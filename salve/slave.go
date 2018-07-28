package main

const (
	//连接主节点
	recvMessageNum = 10000
	trainNum       = 5 //12秒轮训
)

//var block = make([][]string, 0)
var levelDB *LevelDB
//var cleanDB  = time.NewTicker(100 * time.Second)

func main() {
	levelDB = NewLevelDB()
	Server()
	//for {
	//	select {
	//	case <-cleanDB.C:
	//		//db.Close()
	//		//time.Sleep(10*time.Nanosecond)
	//		//err := os.RemoveAll("./db")
	//		//if err != nil {
	//		//	panic("删除db失败")
	//		//}
	//		//db, _ = leveldb.OpenFile("./db", nil)
	//		block = make([][]string, 0)
	//		levelDB.MessSlcie=make([]string,0)
	//		levelDB.MessMap=make(map[[16]byte]int,0)
	//	default:
	//	}
	//}
}
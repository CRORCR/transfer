package proto

import (

"fmt"
"io"
"net"

)

func ReadPacket(conn net.Conn) (body []byte, cmd int32, err error) {
	var length int32
	var buf = make([]byte, length)
	_, err = io.ReadFull(conn, buf)
	var curReadBytes int32 = 0
	for {
		n, errRet := conn.Read(buf)
		if errRet != nil {
			err = errRet
			fmt.Printf("read body from conn %v failed, err:%v\n", conn, err)
			return
		}

		body = append(body, buf[0:n]...)
		curReadBytes += int32(n)
		if curReadBytes == length{
			break
		}
		buf = make([]byte, length - curReadBytes)
	}
	return
}

func WritePacket(conn net.Conn,body []byte) (err error) {
	//写入body
	var n int
	var sendBytes int
	msgLen := len(body) //消息长度
	for {
		n, err = conn.Write(body)
		if err != nil {
			fmt.Printf("send to client:%v failed, err:%v\n", conn, err)
			return
		}
		sendBytes += n
		//判断是否全部发送,不是 续传
		if sendBytes >= msgLen {
			break
		}
		body = body[sendBytes:]
	}
	fmt.Printf("write body succ:%v\n", string(body))
	return
}
package main

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

type ReadConfig struct {
	filename       string
	data           map[string]string
	rwLock         sync.RWMutex
}

func NewConfig(filename string) (conf *ReadConfig, err error) {
	conf = &ReadConfig{
		filename: filename,
		data:     make(map[string]string, 1024),
	}
	m, _ := conf.parse()
	conf.rwLock.Lock()
	conf.data = m
	conf.rwLock.Unlock()
	return
}

func (c *ReadConfig) parse() (m map[string]string, err error) {
	m = make(map[string]string, 1024)
	file, err := os.Open(c.filename)
	if err != nil {
		return
	}

	var lineNo int
	reader := bufio.NewReader(file)
	for {
		line, _,errRet := reader.ReadLine()
		lineStr:=string(line)

		if errRet != nil {
			err = errRet
			return
		}

		lineNo++
		//去除空格
		lineStr = strings.TrimSpace(lineStr)
		if len(line) == 0 || line[0] == '\n' || line[0] == '#' || line[0] == ';' {
			continue
		}
		arr := strings.Split(lineStr, "=")
		key := strings.TrimSpace(arr[0])
		value := strings.TrimSpace(arr[1])
		m[key] = value
	}
	return
}

func (c *ReadConfig) GetString(key string) (value string) {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	value = c.data[key]
	return
}

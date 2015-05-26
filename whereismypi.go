package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func main() {

	var w sync.WaitGroup

	for i := 1; i < 255; i++ {
		w.Add(1)
		//Scan my subnet 192.168.1.[1-255]
		go scan("192.168.1."+strconv.Itoa(i), 22, &w)
	}
	w.Wait()
}

func scan(host string, port int, w *sync.WaitGroup) {
	portString := strconv.Itoa(port)
	timeD, _ := time.ParseDuration("500ms")
	conn, err := net.DialTimeout("tcp", host+":"+portString, timeD)
	if err == nil {
		conn.Close()
		fmt.Printf("%s is UP\n", host)
	}
	(*w).Done()
}

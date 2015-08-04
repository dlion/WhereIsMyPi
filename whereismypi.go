package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	var w sync.WaitGroup
	found := false

	for i := 1; i < 255 && found != true; i++ {
		w.Add(1)
		//Scan my subnet 192.168.1.[1-255]
		go scan("192.168.1."+strconv.Itoa(i), 22, &w, &found)
	}
	w.Wait()
	// none of our goroutines found the rpi in the routing table
	if !found {
		log.Fatal("Couldn't locate your Raspberry Pi.")
	}
}

func scan(host string, port int, w *sync.WaitGroup, found *bool) {
	portString := strconv.Itoa(port)
	timeD, _ := time.ParseDuration("500ms")
	conn, err := net.DialTimeout("tcp", host+":"+portString, timeD)
	if err == nil {
		conn.Close()
		table, err := os.Open("/proc/net/arp")
		defer table.Close()
		if err != nil {
			log.Panic("Impossible to open arp table on your machine")
		}
		//Read the file line by line
		scanner := bufio.NewScanner(table)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() && *found != true {
			fields := strings.Fields(scanner.Text())
			if fields[2] == "0x2" && strings.Contains(fields[3], "b8:27:eb:") {
				fmt.Printf("You Raspberry Pi is up on: %s\n", host)
				*found = true
			}
		}
	}
	(*w).Done()
}

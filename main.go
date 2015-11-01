package main

import (
	"fmt"
	"github.com/golang/groupcache/lru"
	"os"
	"os/signal"
	"time"
)

func main() {

	server := &Server{
		host:     "127.0.0.2",
		port:     53,
		rTimeout: 5 * time.Second,
		wTimeout: 5 * time.Second,
	}

	server.Run()
	fmt.Println("DNS server start")

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)

forever:
	for {
		select {
		case <-sig:
			fmt.Println("Signal recieved, now stop and exit")
			break forever
		}
	}

}

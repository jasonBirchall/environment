package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/tatsushid/go-fastping"
)

func main() {
	ping := fastping.NewPinger()
	resolve, err := net.ResolveIPAddr("ip4:icmp", os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(resolve)
	// fmt.Println(ping)

	ping.AddIPAddr(resolve)
	ping.OnRecv = func(address *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", address.String(), rtt)
	}
	ping.OnIdle = func() {
		fmt.Println("finish")
	}

	err = ping.Run()
	if err != nil {
		fmt.Println(err)
		fmt.Println("didn't work")
	}
}

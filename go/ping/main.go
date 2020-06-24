package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/tatsushid/go-fastping"
)

func main() {
	doEvery(50*time.Millisecond, ping)
}

// Idea taken from https://gist.github.com/ryanfitz/4191392
func doEvery(duration time.Duration, function func(time.Time)) {
	for time := range time.Tick(duration) {
		function(time)
	}
}

func ping(t time.Time) {
	ping := fastping.NewPinger()
	resolve, err := net.ResolveIPAddr("ip4:icmp", os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ping.AddIPAddr(resolve)
	ping.OnRecv = func(address *net.IPAddr, rtt time.Duration) {
		fmt.Println("ping:", rtt)
	}

	err = ping.Run()
	if err != nil {
		fmt.Println(err)
		fmt.Println("didn't work")
	}
}

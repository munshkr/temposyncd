package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"time"
)

const (
	Version         = "0.1.0"
	multicastAddr   = "225.0.0.1:5000"
	maxDatagramSize = 8192
)

var isLeader bool = false

func main() {
	leaderFlagPtr := flag.Bool("leader", false, "join as leader")
	versionFlagPtr := flag.Bool("version", false, "print version")

	flag.Parse()

	if *versionFlagPtr {
		fmt.Printf("temposyncd version %s %s/%s\n", Version, runtime.GOOS, runtime.GOARCH)
		os.Exit(1)
	}

	isLeader = *leaderFlagPtr

	// Resolve multicast addr
	addr, err := net.ResolveUDPAddr("udp4", multicastAddr)
	if err != nil {
		log.Fatal(err)
	}

	if isLeader {
		log.Println("Joined as leader")

		// Start ticker
		go tickTime(addr)
	} else {
		log.Println("Joined as follower")
	}

	fmt.Println("temposyncd started")
	listenMulticast(addr)
}

func tickTime(addr *net.UDPAddr) {
	c, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	t := time.Tick(1000 * time.Millisecond)
	for now := range t {
		log.Printf("TICK %v\n", now)
		c.Write([]byte("hello!\n"))
	}
}

func listenMulticast(addr *net.UDPAddr) {
	l, err := net.ListenMulticastUDP("udp4", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	l.SetReadBuffer(maxDatagramSize)

	for {
		b := make([]byte, maxDatagramSize)
		n, src, err := l.ReadFromUDP(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		msgHandler(src, n, b)
	}
}

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))
}

package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"runtime"
	"time"
)

import (
	"github.com/hypebeast/go-osc/osc"
)

const (
	Version         = "0.1.0"
	multicastAddr   = "225.0.0.1:5000"
	maxDatagramSize = 8192
)

func main() {
	leaderFlagPtr := flag.Bool("leader", false, "join as leader")
	versionFlagPtr := flag.Bool("version", false, "print version")
	verboseFlagPtr := flag.Bool("verbose", false, "print debugging information")

	flag.Parse()

	if *versionFlagPtr {
		fmt.Printf("temposyncd version %s %s/%s\n", Version, runtime.GOOS, runtime.GOARCH)
		os.Exit(1)
	}

	if !*verboseFlagPtr {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	isLeader := *leaderFlagPtr

	start(isLeader)
}

func start(isLeader bool) {
	// Resolve multicast addr
	addr, err := net.ResolveUDPAddr("udp4", multicastAddr)
	if err != nil {
		log.Fatal("ResolveUDPAddr failed:", err)
	}

	client := osc.NewClient("localhost", 57120)

	if isLeader {
		log.Println("Joined as leader")

		// Start ticker
		go tickTime(addr)
	} else {
		log.Println("Joined as follower")
	}

	fmt.Println("temposyncd started")

	// Listen to broadcasts
	listenMulticast(addr, client)
}

func tickTime(addr *net.UDPAddr) {
	c, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("DialUDP failed:", err)
	}

	t := time.Tick(250 * time.Millisecond)
	for now := range t {
		log.Printf("TICK %v\n", now)
		c.Write([]byte("hello!\n"))
	}
}

func listenMulticast(addr *net.UDPAddr, client *osc.Client) {
	l, err := net.ListenMulticastUDP("udp4", nil, addr)
	if err != nil {
		log.Fatal("ListenMulticastUDP failed:", err)
	}
	l.SetReadBuffer(maxDatagramSize)

	for {
		b := make([]byte, maxDatagramSize)
		n, src, err := l.ReadFromUDP(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		msgHandler(src, n, b, client)
	}
}

func msgHandler(src *net.UDPAddr, n int, b []byte, client *osc.Client) {
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))

	// Send OSC /tick message
	msg := osc.NewMessage("/temposync/tick")
	client.Send(msg)
	log.Printf("Sent 'ts %v' to /temposync/tick", time.Now())
}

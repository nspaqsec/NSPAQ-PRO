package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func ShowDashboard() {
	fmt.Println("\033[H\033[2J")
	PrintBanner()
	fmt.Println("\033[34m┏━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━┓\033[0m")
	for {
		time.Sleep(1 * time.Second)
		rps := atomic.SwapUint64(&reqSent, 0)
		totalS := atomic.LoadUint64(&success)
		totalF := atomic.LoadUint64(&fails)

		fmt.Printf("\033[10;0H")
		fmt.Printf("\033[34m┃\033[0m \033[1mRPS\033[0m        \033[34m┃\033[0m \033[32m%-21d\033[0m \033[34m┃\033[0m\n", rps)
		fmt.Printf("\033[34m┃\033[0m \033[1mSUCCESS\033[0m    \033[34m┃\033[0m \033[36m%-21d\033[0m \033[34m┃\033[0m\n", totalS)
		fmt.Printf("\033[34m┃\033[0m \033[1mFAILS\033[0m      \033[34m┃\033[0m \033[31m%-21d\033[0m \033[34m┃\033[0m\n", totalF)
		fmt.Println("\033[34m┗━━━━━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━┛\033[0m")
	}
}

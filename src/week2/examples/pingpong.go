// Ping pong
package main

import (
	"fmt"
	"time"
)

func pinger(ping chan int, pong chan int) {
	for {
		<-ping
		fmt.Println("ping")
		time.Sleep(1 * time.Second)
		pong <- 1
	}
}

func ponger(ping chan int, pong chan int) {
	for {
		<-pong
		fmt.Println("pong")
		time.Sleep(1 * time.Second)
		ping <- 1
	}
}

func main() {
	ping, pong := make(chan int), make(chan int)
	go pinger(ping, pong)
	go ponger(ping, pong)
	ping <- 1
	fmt.Scanln()
}

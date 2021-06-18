package tuto

import (
	"fmt"
	"math/rand"
	"time"
)

func MakeChannel() {

	messages := make(chan string)
	fmt.Println("1")

	go worker(messages, func() string { return "ping1" })
	go worker(messages, func() string { return "ping2" })
	go worker(messages, func() string { return "ping3" })
	go worker(messages, func() string { return "ping4" })

	fmt.Println("4")
	for msg := range messages {
		println(string(msg))
	}
	fmt.Println("5")
}

func worker(messages chan string, callback func() string) {
	fmt.Println("2")
	randomTime := rand.Intn(3) + 1
	for {
		messages <- callback() // emetteur
		time.Sleep(time.Duration(randomTime) * time.Second)
	}
	// fmt.Println("3")
}

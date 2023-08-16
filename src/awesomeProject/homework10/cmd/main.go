package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var pnt1, pnt2 int
	const N = 21
	wg.Add(N)
	var ch1 = make(chan string)

	go playerOne(ch1, ch1, &pnt2, &wg)
	go playerTwo(ch1, ch1, &pnt1, &wg)
	ch1 <- "begin"
	wg.Wait()
	fmt.Println("-------------------------------------------------")

	fmt.Println("Player 1 -", pnt1, "and Player 2 - ", pnt2)

}
func playerOne(ch1 <-chan string, ch2 chan<- string, pnt *int, wg *sync.WaitGroup) {
	for {
		switch val, _ := <-ch1; val {
		case "ping":
			fmt.Println("Player 2: ping")
			i := rand.Intn(100)
			if i < 21 {
				fmt.Println("Player 2 ---> SCORE!!!")
				*pnt++
				fmt.Println("Player 2 has: ", *pnt)
				ch2 <- "stop"
				wg.Done()
			} else {
				ch2 <- "pong"
			}
		case "pong":
			fmt.Println("Player 2: pong")
			i := rand.Intn(100)
			if i < 21 {
				fmt.Println("Player 2 ---> SCORE!!!")
				defer wg.Done()
				*pnt++
				fmt.Println("Player 2 has: ", *pnt)
				ch2 <- "stop"
				wg.Done()
			} else {
				ch2 <- "ping"
			}
		case "stop":
			//fmt.Println("case stop")
			ch2 <- "begin"
		case "begin":
			//fmt.Println("case begin")
			ch2 <- "ping"
		}
	}
}

func playerTwo(ch1 <-chan string, ch2 chan<- string, pnt *int, wg *sync.WaitGroup) {
	for {
		switch val, _ := <-ch1; val {
		case "ping":
			fmt.Println("Player 1: ping")
			i := rand.Intn(100)
			if i < 21 {
				fmt.Println("Player 1 ---> SCORE!!!")
				*pnt++
				fmt.Println("Player 1 has: ", *pnt)
				ch2 <- "stop"
				wg.Done()
			} else {
				ch2 <- "pong"
			}
		case "pong":
			fmt.Println("Player 1: pong")
			i := rand.Intn(100)
			if i < 21 {
				fmt.Println("Player 1 ---> SCORE!!!")
				defer wg.Done()
				*pnt++
				fmt.Println("Player 1 has: ", *pnt)
				ch2 <- "stop"
				wg.Done()
			} else {
				ch2 <- "ping"
			}
		case "stop":
			//fmt.Println("case stop")
			ch2 <- "ping"
		case "begin":
			//fmt.Println("case begin")
			ch2 <- "ping"
		}
	}
}

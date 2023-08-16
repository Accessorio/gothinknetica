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
	var ch = make(chan string)

	go playerOne(ch, ch, &pnt2, &wg)
	go playerTwo(ch, ch, &pnt1, &wg)
	ch <- "begin"
	wg.Wait()
	fmt.Println("-------------------------------------------------")

	fmt.Println("Player 1 -", pnt1, "and Player 2 - ", pnt2)

}
func playerOne(chr <-chan string, chw chan<- string, pnt *int, wg *sync.WaitGroup) {
	for {
		switch val, _ := <-chr; val {
		case "ping":
			fmt.Println("Player 2: ping")
			i := rand.Intn(100)
			if i < 21 {
				fmt.Println("Player 2 ---> SCORE!!!")
				*pnt++
				fmt.Println("Player 2 has: ", *pnt)
				chw <- "stop"
				wg.Done()
			} else {
				chw <- "pong"
			}
		case "pong":
			fmt.Println("Player 2: pong")
			i := rand.Intn(100)
			if i < 21 {
				fmt.Println("Player 2 ---> SCORE!!!")
				*pnt++
				fmt.Println("Player 2 has: ", *pnt)
				chw <- "stop"
				wg.Done()
			} else {
				chw <- "ping"
			}
		case "stop":
			chw <- "begin"
		case "begin":
			chw <- "ping"
		}
	}
}

func playerTwo(chr <-chan string, chw chan<- string, pnt *int, wg *sync.WaitGroup) {
	for {
		switch val, _ := <-chr; val {
		case "ping":
			fmt.Println("Player 1: ping")
			i := rand.Intn(100)
			if i < 21 {
				fmt.Println("Player 1 ---> SCORE!!!")
				*pnt++
				fmt.Println("Player 1 has: ", *pnt)
				chw <- "stop"
				wg.Done()
			} else {
				chw <- "pong"
			}
		case "pong":
			fmt.Println("Player 1: pong")
			i := rand.Intn(100)
			if i < 21 {
				fmt.Println("Player 1 ---> SCORE!!!")
				*pnt++
				fmt.Println("Player 1 has: ", *pnt)
				chw <- "stop"
				wg.Done()
			} else {
				chw <- "ping"
			}
		case "stop":
			chw <- "begin"
		case "begin":
			chw <- "ping"
		}
	}
}

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

	go player(ch, &pnt1, &wg, "Player 1")
	go player(ch, &pnt2, &wg, "Player 2")
	ch <- "begin"
	wg.Wait()
	fmt.Println("-------------------------------------------------")

	fmt.Println("Player 1 -", pnt1, "and Player 2 - ", pnt2)

}
func player(ch chan string, pnt *int, wg *sync.WaitGroup, pl string) {
	for {
		switch val, _ := <-ch; val {
		case "ping":
			fmt.Println(pl + ": ping")
			i := rand.Intn(100)
			if i < 21 {
				fmt.Println(pl + " ---> SCORE!!!")
				*pnt++
				fmt.Println(pl+" has: ", *pnt)
				ch <- "stop"
				wg.Done()
			} else {
				ch <- "pong"
			}
		case "pong":
			fmt.Println(pl + ": pong")
			i := rand.Intn(100)
			if i < 21 {
				fmt.Println(pl + " ---> SCORE!!!")
				*pnt++
				fmt.Println(pl+" has: ", *pnt)
				ch <- "stop"
				wg.Done()
			} else {
				ch <- "ping"
			}
		case "stop":
			ch <- "begin"
		case "begin":
			ch <- "ping"
		}
	}
}

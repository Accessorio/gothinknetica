package main

import (
	"bufio"
	"go-core-4/homework11/pkg/netsrv"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	r := bufio.NewReader(os.Stdin)
	t := bufio.NewReader(conn)
	for {
		err = netsrv.ClientRequest(r, conn)
		if err != nil {
			log.Fatal(err)
		}
		err = netsrv.ServerResponce(t)
		if err != nil {
			log.Fatal(err)
		}
	}
}

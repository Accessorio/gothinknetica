package netsrv

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func ResponceFromC(r *bufio.Reader, conn net.Conn) {

	fmt.Println("Search: ")
	rsp, _, err := r.ReadLine()
	if err != nil {
		if err == io.EOF {
			fmt.Println("User closed the connection")
		} else {
			log.Fatal(err)
		}
	}

	rsp = append(rsp, '\n')

	_, err = conn.Write(rsp)
	if err != nil {
		log.Fatal(err)
	}

}

func RequestFromS(r *bufio.Reader) {
	rqst, err := r.ReadBytes('*')
	if err != nil {
		log.Fatalf("Bad Request", err)
	}
	if len(rqst) > 0 {
		fmt.Println("Results: \n" + string(rqst))
	}
}

func RequestFormC(r *bufio.Reader) (string, error) {
	var rqst string
	rqst, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("Error in client request")
		return "", err
	}
	rqst = strings.TrimSpace(rqst)
	return rqst, err
}

func ResponeToC(conn net.Conn, m []byte) error {
	_, err := conn.Write(m)
	if err != nil {
		fmt.Println("Error in Responce")
		return err
	}
	return err
}

package netsrv

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func ClientRequest(r *bufio.Reader, conn net.Conn) error {

	fmt.Println("Search: ")
	rsp, _, err := r.ReadLine()
	if err != nil {
		if err == io.EOF {
			fmt.Println("User closed the connection")
		} else {
			fmt.Println("Error in ClientResponce ReadLine()")
			return err
		}
	}

	rsp = append(rsp, '\n')

	_, err = conn.Write(rsp)
	if err != nil {
		fmt.Println("Error in ClientResponce Write()")
		return err
	}
	return err
}

func ServerResponce(r *bufio.Reader) error {
	rqst, err := r.ReadBytes('*')
	if err != nil {
		fmt.Println("Error in ServerRequest")
		return err
	}
	if len(rqst) > 0 {
		fmt.Println("Results: \n" + string(rqst))
	}
	return err
}

func ClientResponce(r *bufio.Reader) (string, error) {
	var rqst string
	rqst, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("Error in client request")
		return "", err
	}
	rqst = strings.TrimSpace(rqst)
	return rqst, err
}

func ServerRequest(conn net.Conn, m []byte) error {
	_, err := conn.Write(m)
	if err != nil {
		fmt.Println("Error in Responce")
		return err
	}
	return err
}

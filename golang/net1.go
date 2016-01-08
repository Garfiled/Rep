package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "wwww.baidu.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	reader := bufio.NewReader(conn)
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				continue
			}
			fmt.Println(err)
			break
		}
		fmt.Println(string(data))
	}
}

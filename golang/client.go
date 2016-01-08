package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const RECV_BUF_LEN = 1024

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	buf := make([]byte, RECV_BUF_LEN)

	for {
		fmt.Printf(">")
		cmdReader := bufio.NewReader(os.Stdin)
		if cmdstr, err := cmdReader.ReadString('\n'); err == nil {
			cmdstr = strings.Trim(cmdstr, "\r\n")
			if cmdstr == "" {
				continue
			} else if cmdstr == "q" {
				break
			} else if _, err := conn.Write([]byte(cmdstr)); err != nil {
				println("Write Buffer Error:", err.Error())
				break
			}
		}
		if n, err := conn.Read(buf); err != nil {
			println("Read Buffer Error:", err.Error())
			break
		} else {
			fmt.Println(string(buf[0:n]))
		}
		time.Sleep(time.Second)
	}
}

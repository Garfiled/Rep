package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
)

const RECV_BUF_LEN = 1024

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9001")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		fmt.Printf(">")
		cmdReader := bufio.NewReader(os.Stdin)
		if cmdstr, err := cmdReader.ReadString('\n'); err == nil {
			cmdstr = strings.Trim(cmdstr, "\r\n")
			if cmdstr == "" {
				continue
			} else if cmdstr == "q" {
				break
			} else {
				bytes := GenProtocal([]byte(cmdstr))
				_, err := conn.Write(bytes)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	conn.Close()
}

func GenProtocal(cmd []byte) []byte {
	var buf bytes.Buffer
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(len(cmd)+4))
	buf.Write(b)
	binary.BigEndian.PutUint16(b, uint16(0))
	buf.Write(b)
	buf.Write(cmd)
	return buf.Bytes()
}

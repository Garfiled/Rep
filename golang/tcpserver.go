package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

const (
	ip   = ""
	port = 9001
)

func main() {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), port, ""})
	if err != nil {
		fmt.Println(err)
		return
	}
	//	Server(listen)
	fmt.Printf("listen on %d\n", port)
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go client(conn)
	}

}

func client(conn *net.TCPConn) {
	stashBytes := make([]byte, 0)
	buf := make([]byte, 1024)
	var temp []byte
	for {
		count, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		if len(stashBytes) > 0 {
			temp = append(stashBytes, buf[:count]...)
		} else {
			temp = buf[:count]
		}
		start := 0
		for start < count {
			l := int(binary.BigEndian.Uint16(temp[start : start+2]))
			if start+l > count {
				stashBytes = temp[start:]
				break
			}
			eid, recvMsg := ParseProtocal(temp[start+2 : start+l])
			fmt.Println(eid, string(recvMsg))
			start += l
		}
	}
	conn.Close()
}

func ParseProtocal(b []byte) (uint16, []byte) {
	e := binary.BigEndian.Uint16(b[:2])
	return e, b[2:]
}

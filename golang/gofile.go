package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readall() {
	f, err := os.Open("txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}

func readfile() {
	bytes, err := ioutil.ReadFile("txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}

func readblock() {
	f, err := os.Open("txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	buf := make([]byte, 1024)
	bfrd := bufio.NewReader(f)
	for {
		n, err := bfrd.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("found EOF")
			} else {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(buf[:n])
	}
}

func readline() {
	f, err := os.Open("txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	bfrd := bufio.NewReader(f)
	for {
		line, err := bfrd.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("found EOF")
			} else {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(line)
	}
}
func main() {
	//	readall()
	//	readfile()
	//	readblock()
	readline()
}

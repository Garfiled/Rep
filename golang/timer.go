package main

import (
	"fmt"
	"time"
)

var sum int64

func add(result chan int64) {
	for i := 0; i < 10000; i++ {
		for j := 0; j < 10000; j++ {
			sum++
		}
	}
	result <- sum
}
func main() {
	timer1 := time.NewTicker(200 * time.Millisecond)
	result := make(chan int64, 1)
	go add(result)
	t1 := time.Now()
	rflag := false
	tflag := false
	for {
		select {
		case <-timer1.C:
			tflag = true
			if rflag {
				break
			}
		case <-result:
			rflag = true
			if tflag {
				break
			}
		}
		if rflag && tflag { // 等够时间再走
			break
		}
	}

	t2 := time.Now().Sub(t1).Nanoseconds() / 1000000
	fmt.Println(t2)
}

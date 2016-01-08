package main

// 导入sql包
import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
)

func main() {
	db, e := sql.Open("mysql", "root:root@tcp(localhost:3306)/gotest?charset=utf8")
	if e != nil { //如果连接出错,e将不是nil的
		fmt.Println("ERROR?")
		return
	}
	for {
		fmt.Printf(">")
		cmdReader := bufio.NewReader(os.Stdin)
		if cmdstr, err := cmdReader.ReadString('\n'); err == nil {
			cmdstr = strings.Trim(cmdstr, "\r\n")
			if cmdstr == "q" {
				break
			} else if cmdstr == "" {
				continue
			} else {
				res, e := db.Query(cmdstr)
				if e != nil {
					fmt.Println(e)
				} else {
					var a string
					for res.Next() {
						e1 := res.Scan(&a)
						if e1 != nil {
							fmt.Println("debug1")
						} else {
							fmt.Println(a)
						}
					}
				}
				// fmt.Println(res, e)
			}
		} else {
			fmt.Println(err)
		}
	}
	db.Close()
}

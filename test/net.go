/**
User:       wliangde
CreateTime: 18/2/28 下午5:30
**/
package test

import (
	//"io"
	"log"
	"net"
	//"time"
	"bufio"
	"fmt"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(con net.Conn) {
	log.Println("收到客户端连接，地址", con.RemoteAddr())
	//io.Copy(con, con)
	input := bufio.NewScanner(con)
	for input.Scan() {
		echo(con, input.Text(), time.Second)
	}
	con.Close()

	//for {
	//	_, err := io.WriteString(con, time.Now().Format("15:04:05\n"))
	//	if err != nil {
	//		log.Println(err)
	//		return // e.g., client disconnected
	//	}
	//	time.Sleep(1 * time.Second)
	//}
}

func StartServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error", err)
	}
	log.Println("服务器开始监听端口", 8080)
	log.Println("等待客户端连接")
	defer ln.Close()
	for {
		con, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(con)
	}
}

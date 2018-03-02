/**
User:       wliangde
CreateTime: 18/3/1 下午4:36
**/
package test

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//++++++++++++++++++++++++广播聊天

type NetChatTest struct {
	addr string
}

var DftNetChatTest = NetChatTest{":8080"}

func (this *NetChatTest) Main() {
	ln, err := net.Listen("tcp", this.addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("服务器启动成功，开始监听", this.addr)

	defer ln.Close()
	//启动广播routine
	go this.broadcast()
	for {
		con, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go this.handleConn(con)
	}
}

//一个客户端通道
type client chan<- string

//客户端管理
type clientMgr map[client]bool

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) //所有客户端的消息
)

func (this *NetChatTest) handleConn(con net.Conn) {
	cl := make(chan string)

	go this.sendMsg2Client(con, cl)

	who := con.RemoteAddr().String()
	cl <- "You are " + who
	messages <- who + "has arrived"
	//加入
	entering <- cl

	//接收客户消息
	input := bufio.NewScanner(con)
	for input.Scan() {
		messages <- who + ":" + input.Text()
	}
	//离开
	leaving <- cl
	messages <- who + "has left"
	con.Close()
}

//消息广播
func (this *NetChatTest) broadcast() {
	clientMgr := make(clientMgr)

	for {
		select {
		case c := <-entering:
			clientMgr[c] = true
		case c := <-leaving:
			delete(clientMgr, c)
		case msg := <-messages:
			for cli := range clientMgr {
				cli <- msg
			}
		}
	}
}

//发送消息给客户端
func (this *NetChatTest) sendMsg2Client(con net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprint(con, msg)
	}
}

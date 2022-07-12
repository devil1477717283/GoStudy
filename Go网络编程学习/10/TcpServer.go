package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from cilent failed,err", err)
			break
		}
		recvStr := "[Echo Message] " + string(buf[:n])
		fmt.Println("收到client端发来的数据:", string(buf[:n]))
		conn.Write([]byte(recvStr))
	}
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("listen failed,err", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err", err)
			continue
		}
		fmt.Println("成功接收一个连接请求")
		go process(conn)
	}

}

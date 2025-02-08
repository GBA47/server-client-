package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1. 连接RPC服务器
	conn, err := rpc.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	// 2. 远程调用
	var reply string
	err = conn.Call("hello.HelloWorld", "李白", &reply)
	if err != nil {
		fmt.Println("Call err:", err)
		return
	}
	fmt.Println(reply) // 期待输出：李白 你好！
}

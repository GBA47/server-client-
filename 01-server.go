package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

// 定义类方法
type World struct {
}

// 绑定类方法
func (this *World) HelloWorld(name string, resp *string) error {
	*resp = name + " 你好！" // 修改为正确的返回值
	return errors.New("未知的错误！")
}

func main() {
	// 1. 注册RPC服务，绑定对象方法
	/*err := rpc.RegisterName("hello", new(World))
	if err != nil {
		fmt.Println("注册RPC服务失败！", err)
		return
	}*/
	RegisterService(new(World))
	// 2. 监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	fmt.Println("开始监听...")

	// 3. 接受连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("连接成功...")

	// 4. 绑定服务
	rpc.ServeConn(conn)
}

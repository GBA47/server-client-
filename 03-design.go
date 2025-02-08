package main

import (
	"fmt"
	"net/rpc"
)

// MyInterface 定义了一个接口
type MyInterface interface {
	HelloWorld(string, *string) error
}

// 具体的服务对象类型，实现 MyInterface 接口
type MyService struct{}

// 实现 MyInterface 接口中的 HelloWorld 方法
func (s *MyService) HelloWorld(request string, reply *string) error {
	*reply = "Hello, " + request
	return nil
}

// RegisterService 注册服务
func RegisterService(i MyInterface) {
	rpc.RegisterName("hello", i)
}

func main() {
	// 创建一个具体的服务对象
	service := new(MyService)

	// 注册服务对象
	RegisterService(service)

	// 这里你可以启动服务器或客户端进行 RPC 调用
	fmt.Println("Service registered.")
}

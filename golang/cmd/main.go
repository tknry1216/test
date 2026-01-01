package main

import (
	"fmt"
	"net"
	"syscall"
)

func main() {
	// 低レベルでソケットを作成する例

	// 1. socket() システムコール
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd)
	fmt.Printf("Socket created: fd=%d\n", fd)

	// 2. bind() - ポートをバインド
	addr := syscall.SockaddrInet4{
		Port: 8080,
		Addr: [4]byte{0, 0, 0, 0}, // 0.0.0.0
	}
	err = syscall.Bind(fd, &addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Socket bound to :8080")

	// 3. listen() - リスニング開始
	err = syscall.Listen(fd, 128) // backlog = 128
	if err != nil {
		panic(err)
	}
	fmt.Println("Socket listening")

	// これが net.Listen() が内部でやっていること！

	// 高レベルAPI（推奨）
	listener, _ := net.Listen("tcp", ":8081")
	fmt.Printf("net.Listen created: %T\n", listener)
	listener.Close()
}

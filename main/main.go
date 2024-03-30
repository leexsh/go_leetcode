package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type ListNode struct {
	val  int
	next *ListNode
}
type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   int
}

func endpointToUint64(endpoint string) uint64 {
	// 分割IPv4地址和端口号
	parts := strings.Split(endpoint, ":")
	ip := strings.Split(parts[0], ".")
	// ip := net.ParseIP(parts[0]).To4()
	ip1, _ := strconv.Atoi(ip[0])
	ip2, _ := strconv.Atoi(ip[1])
	ip3, _ := strconv.Atoi(ip[2])
	ip4, _ := strconv.Atoi(ip[3])
	port, _ := strconv.Atoi(parts[1])
	var res uint64 = 0
	portValue := uint64(port)
	res |= uint64(ip1) << 24
	res |= uint64(ip2) << 16
	res |= uint64(ip3) << 8
	res |= uint64(ip4)
	fmt.Println("res is :", res)
	res = res << 32

	res |= uint64(portValue)
	return res
}

func uint64ToEndpoint(value uint64) string {
	// 从64位整数中提取IPv4地址和端口号
	port := value & 0xffff
	value = value >> 32
	ip := net.IPv4(byte(value>>24), byte(value>>16), byte(value>>8), byte(value))

	// 格式化IPv4地址和端口号为字符串
	return fmt.Sprintf("%s:%d", ip, port)
}

func main() {
	endpoint := "127.0.0.1:8080"

	// 将端点转换为uint64
	uintValue := endpointToUint64(endpoint)
	fmt.Println("Endpoint to uint64:", uintValue)

	// 将uint64转换为端点
	endpointValue := uint64ToEndpoint(uintValue)
	fmt.Println("Uint64 to endpoint:", endpointValue)
}

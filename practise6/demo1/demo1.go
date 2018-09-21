package main

import (
	"fmt"
)

func main() {
	fmt.Println(JRCodeSucc)
	fmt.Println(JRCodeFailed)

}

type JsonResultCode int

const (
	JRCodeSucc JsonResultCode = iota
	JRCodeFailed
	JRCode302 = 302 //跳转至地址
	JRCode401 = 401 //未授权访问
)

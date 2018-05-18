package main

import (
	"fmt"
	"go_learning/mime/mul"
	"io/ioutil"
)

type LoadFuTan struct {
	mul.Request
}

func main() {
	m := LoadFuTan{}
	m.GetF()
}
func (this *LoadFuTan) GetF() {
	url := "/Users/apple/Downloads/test.txt"
	f, h, err := this.GetFile(url)
	fmt.Println(err)
	fmt.Println(h.Filename)
	data, err := ioutil.ReadAll(f)
	str := string(data)
	fmt.Println(data, str, err)

}

package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func Test_echo(t *testing.T) {
	s, sep := "333", "444"
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("111", s)

	fmt.Println("aggg", strings.Join(os.Args[1:], "333"))
}

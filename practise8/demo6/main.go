package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}
func Ip2Long(ip string) (ips string) {
	var ip_pieces = strings.Split(ip, ".")
	ip_1, _ := strconv.ParseInt(ip_pieces[0], 10, 32)
	ip_2, _ := strconv.ParseInt(ip_pieces[1], 10, 32)
	ip_3, _ := strconv.ParseInt(ip_pieces[2], 10, 32)
	ip_4, _ := strconv.ParseInt(ip_pieces[3], 10, 32)
	var ip_bin string = fmt.Sprintf("%08b%08b%08b%08b", ip_1, ip_2, ip_3, ip_4)
	ip_int, _ := strconv.ParseInt(ip_bin, 2, 64)
	return
}

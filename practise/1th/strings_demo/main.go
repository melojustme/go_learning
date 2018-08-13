package main

import (
	"fmt"
	"strings"
)

//标准库分析之strings
func main() {

	fmt.Println("1.将切片s分割成由sep分隔的所有子字符串，并返回这些分隔符之间的子字符串切片")
	accout := "15111112222,13122223333,16122223333"
	accouts := strings.Split(accout, ",")
	fmt.Println(accouts)
	fmt.Println("2.TrimPrefix在没有提供的前导前缀字符串的情况下返回s。如果s不以前缀开始，则s将保持不变。")
	s := "#xxxxx.png"
	s = strings.TrimPrefix(s, "#")
	fmt.Println(s)
	fmt.Println("3.EqualFold报告在Unicode大小写折叠下a和b，解释为UTF-8字符串//是否相等")
	a := "#xdsfsdfdsss33"
	b := "#xdsfsdfdsss33"
	bl := strings.EqualFold(a, b)
	fmt.Println(bl)
	fmt.Println("4.HasPrefix测试字符串s是否以前缀开头")
	query := "Page="
	bl = strings.HasPrefix(query, "Page=")
	fmt.Println(bl)
	fmt.Println("5.Index返回s中第一个substr实例的索引，如果substr不存在于s中，则返回-1")
	query1 := "&Page=2"
	i := strings.Index(query1, "&")
	query1 = query1[i+1:]
	fmt.Println(query1)

	fmt.Println("6.Join连接a的元素以创建单个字符串。分隔符字符串// sep位于结果字符串中的元素之间。")
	ss := []string{"hahahha", "lllll"}
	k := strings.Join(ss, ",")
	fmt.Println(k)

	fmt.Println("7.ContainsRune报告Unicode代码点r是否在s内")
	order := "99999999,888888"
	bl = strings.ContainsRune(order, ',')
	fmt.Println(bl)

	fmt.Println("8.TrimSpace返回字符串s的一部分，删除所有前导//和尾部空白，如Unicode定义的。")
	a1 := "      strings.TrimSpace     "
	fmt.Println(a1)
	a1 = strings.TrimSpace(a1)
	fmt.Println(a1)

	fmt.Println("9.LastIndex返回s中最后一个substr实例的索引，如果substr不存在于s中，则返回-1。")
	query2 := "Page&=2"
	i = strings.LastIndex(query2, "&")
	query2 = query2[i+1:]
	fmt.Println(query2)

	fmt.Println("10.LastIndex返回s中最后一个substr实例的索引，如果substr不存在于s中，则返回-1。")
	order = strings.Replace(order, ",", "", -1)
	fmt.Println(order)

	fmt.Println("11.NewReader从s中返回一个新的Reader读数。 //它与bytes.NewBufferString类似，但效率更高，只读。")
	params := "dfajsdfjskjfksjdfjslkdjfklsj33333232  dfdsfs "
	par := strings.NewReader(params)
	fmt.Println(par)
}

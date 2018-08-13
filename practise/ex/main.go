package main

import (
	"fmt"

	"github.com/apcera/termtables"
)

func main() {
	table := termtables.CreateTable()

	table.AddHeaders("Name", "Age", "hoby")
	table.AddRow("John", "30", "哈哈哈哈哈哈哈哈哈")
	table.AddRow("Sam", 18)
	table.AddRow("Julie", 20.14)

	fmt.Println(table.Render())
}

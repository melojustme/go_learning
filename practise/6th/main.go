package main

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

func main() {
	var c []Config
	s := `[{"model":"hahah,333","value":"1111"}]`
	err := json.Unmarshal([]byte(s), &c)
	beego.Debug(c[0].Model, err)
	beego.Debug(c[0].Value)

}

type Config struct {
	Model string
	Value string
}

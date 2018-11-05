package main

import (
	// "fmt"
	"database/sql"
	"io"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/ziutek/mymysql/godrv"
)

// 留言结构
type Liuyan struct {
	Id      int
	Name    string
	Content string
	Time    int
}

// 显示留言时间
func (l Liuyan) ShowTime() string {
	t := time.Unix(int64(l.Time), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	godrv.Register("SET NAMES utf8")

	// 连接数据库
	db, err := sql.Open("mymysql", "tcp:127.0.0.1:3306*simon/melo/590216")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 准备模板
	tpl, err := template.New("liuyanbook").Parse(html)
	if err != nil {
		panic(err)
	}

	// 显示留言页面 /
	requestList := func(w http.ResponseWriter, req *http.Request) {
		// 查询数据
		rows, err := db.Query("select * from liuyan")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// 获取数据
		lys := []Liuyan{}
		for rows.Next() {
			ly := Liuyan{}
			err := rows.Scan(&ly.Id, &ly.Name, &ly.Content, &ly.Time)
			if nil != err {
				log.Fatal(err)
			}
			lys = append(lys, ly)
		}

		// 显示数据
		err = tpl.ExecuteTemplate(w, "list", lys)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 留言页面 /liuyan
	requestLiuyan := func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		if "POST" == req.Method {
			if len(req.Form["name"]) < 1 {
				io.WriteString(w, "参数错误!\n")
				return
			}
			if len(req.Form["content"]) < 1 {
				io.WriteString(w, "参数错误!\n")
				return
			}

			name := template.HTMLEscapeString(req.Form.Get("name"))
			content := template.HTMLEscapeString(req.Form.Get("content"))

			// sql语句
			sql, err := db.Prepare("insert into liuyan(name, content, time) values(?, ?, ?)")
			if err != nil {
				log.Fatal(err)
			}
			defer sql.Close()

			// sql参数,并执行
			_, err = sql.Exec(name, content, time.Now().Unix())
			if err != nil {
				log.Fatal(err)
			}

			// 跳转
			w.Header().Add("Location", "/")
			w.WriteHeader(302)

			// 提示信息
			io.WriteString(w, "提交成功!\n")

			return
		}

		err = tpl.ExecuteTemplate(w, "liuyan", nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/", requestList)
	http.HandleFunc("/liuyan", requestLiuyan)
	err = http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 网页模板
var html string = `{{define "list"}}{{/* 留言列表页面 */}}<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
    <p><a href="/liuyan">给我留言</a></p>
    <table>
{{range .}}
    <tr>
        <td>{{.Id}}</td><td>{{.Name}}</td><td>{{.Content}}</td><td>{{.ShowTime}}</td>
    </tr>
{{end}}
    </table>
</body>
</html>{{end}}
{{define "liuyan"}}{{/* 发布留言页面 */}}<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
    <form method="post">
        姓名:<input type="text" name="name" /><br>
        内容:<input type="text" name="content" /><br>
        <input type="submit" value="提交" />
    </form>
</body>
</html>{{end}}`

package main

import (
	"fmt"
	"mime"
	"path"
)

/**
MIME is an Internet standard that extends the format of email to support:
Text in character sets other than ASCII
Non-text attachments: audio, video, images, application programs etc.
Message bodies with multiple parts
Header information in non-ASCII character sets

MIME是一种Internet标准，它扩展了电子邮件的格式以支持：
除ASCII以外的字符集中的文本 非文字附件：音频，视频，图像，应用程序等
包含多个部分的消息主体 非ASCII字符集中的标头信息

*/

func main() {
	filepath := "./1.png"
	mimetype := mime.TypeByExtension(path.Ext(filepath))
	fmt.Println(mimetype)

	filepath = "./2.txt"
	mimetype = mime.TypeByExtension(path.Ext(filepath))
	fmt.Println(mimetype)

	filepath = "./3.html"
	mimetype = mime.TypeByExtension(path.Ext(filepath))
	fmt.Println(mimetype)

}

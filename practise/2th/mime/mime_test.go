package mime_test

import (
	"fmt"
	"mime"
	"path"
	"testing"
)

func TestAddExtensionType(t *testing.T) {
	a := mime.AddExtensionType(".svg", "image/svg+xml")
	b := mime.AddExtensionType(".m3u8", "application/x-mpegURL")
	c := mime.AddExtensionType(".ts", "video/MP2T")
	println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func TestTypeByExtension(t *testing.T) {
	fmt.Println("================")
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

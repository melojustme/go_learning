package mul

import (
	"go_learning/wulei/http"
	"mime/multipart"
)

type Request struct {
	http.Request
}

func (r Request) GetFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return r.FormFile(key)
}

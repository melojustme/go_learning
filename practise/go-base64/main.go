package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

var coder = base64.NewEncoding(base64Table)

func main() {
	file, err := os.Open("123.png")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as string: %s\n", Base64Encode(data))
}

func Base64Encode(encode_byte []byte) []byte {
	return []byte(coder.EncodeToString(encode_byte))
}

func Base64Decode(decode_byte []byte) ([]byte, error) {
	return coder.DecodeString(string(decode_byte))
}

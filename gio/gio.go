package gio

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func LoadAsBytes(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return b
}

func LoadAsString(path string) string {
	return string(LoadAsBytes(path))
}

func parseInt(in string) int32 {
	out, err := strconv.ParseInt(in, 0, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int32(out)
}

package gio

import "io/ioutil"

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

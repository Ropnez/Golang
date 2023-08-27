package main

import (
	"errors"
)

func mydelete(m map[string]string, key string) error {
	if m == nil {
		return errors.New("bos map")
	}
	if key == "" {
		return errors.New("key yok")
	}
	if _, ok := m[key]; !ok {
		delete(m, key)
	}
	return nil
}
func main() {
	m := map[string]string{
		"key": "value",
	}
	mydelete(m, "key")
	println(m["key"])
}

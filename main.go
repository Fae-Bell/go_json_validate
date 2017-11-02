package main

import (
	"fmt"

	"github.com/Keith-Ball/go_json_validate/customjson"
	"github.com/Keith-Ball/go_json_validate/test"
)

func main() {
	b := []byte(`{
		"one": 1,
		"two": "ABC",
		"three":true
	}`)
	t := &test.Test{}
	e := customjson.Unmarshal(b, t)
	fmt.Println(t)
	fmt.Println(e)
}

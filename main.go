package main

import (
	"fmt"
	_ "golua/conf"
)

var vmap = make(map[string]int)

func main() {
	vmap["a"] = 1
	v, ok := vmap["a"]
	if ok {
		fmt.Println(v)
	}
}

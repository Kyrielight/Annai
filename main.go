package main

import (
	"moe.best.annai/request"
	"moe.best.annai/resolver"
)

func main() {

	request := request.Request{
		Command:   "g",
		Arguments: []string{"hello", "world"},
	}

	println(resolver.Lookup(request).String())

}

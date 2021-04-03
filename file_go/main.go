package main

import _ "embed"

//go:embed test.txt
var s string

func main(){
	print(s)
}

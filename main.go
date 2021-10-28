package main

import (
	"fmt"
	"pkg/genid"
)

func main() {
	genid.Init("2021-10-27", 22)
	id := genid.GenID()
	fmt.Println(id)
}

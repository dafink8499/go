package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	"strconv"
)

func main() {
	badgers := 1
	for badgers <= 12 {
		fmt.Println(c.Rc() + strconv.Itoa(badgers) + " Badger")
		badgers += 1
	}
}

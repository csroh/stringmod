package main

import (
	"fmt"

	qt "github.com/royroh/stringmod/quotes"
	str "github.com/royroh/stringmod/strings"
)

func main() {
	o, e := str.CountOddEven("12345586")
	fmt.Println(o, e)

	fmt.Println(qt.GetEmoji("turtle"))
}

package main

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
)

func main() {
	s := "Go言語でCLIアプリケーション作成"
	fmt.Println(s)
	width := runewidth.StringWidth(s)
	fmt.Println(strings.Repeat("~", width))
	fmt.Println(runewidth.Wrap(s, 11))
}

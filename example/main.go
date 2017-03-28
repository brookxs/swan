package main

import (
	"fmt"
	"os"

	"github.com/lauyoume/swan"
)

func main() {
	cnstop, _ := swan.NewStopWordChinese("dict.txt")
	swan.RegisterLangStopWord("zh", cnstop)
	//ar, _ := swan.FromURL("http://bgl.shenchuang.com/ds/20170207/423821.shtml")
	ar, _ := swan.FromURL(os.Args[1])
	fmt.Println(ar.CleanedText)
}

package main

import (
	"fmt"
	"os"

	"oot.me/link-parser/utils"
)

func main() {
	htmlFile, err := utils.ReadHTMLFile("ex1.html")
	if err != nil {
		panic(err)
	}
	links, err := utils.ParseLink(htmlFile)
	if err != nil {
		panic(err)
	}
	for _, link := range links {
		fmt.Println(link.Href)
		fmt.Println(link.Text)
	}
	htmlFile.(*os.File).Close()
}

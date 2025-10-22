package main

import (
	"flag"
	"fmt"
	"os"

	"oot.me/link-parser/utils"
)

func main() {
	filename := flag.String("filename", "", "Enter specific HTML file name in full path")
	flag.Parse()
	if *filename == "" {
		fmt.Println("Please input HTML file name")
		return
	}
	htmlFile, err := utils.ReadHTMLFile(*filename)
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

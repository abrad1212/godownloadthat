package main

import (
	//"fmt"
	"github.com/abrad1212/godownloadthat"
)

func main() {
	urls := []string{
		"https://img.icons8.com/cotton/512/000000/chrome.png",
		"https://img.icons8.com/cotton/512/000000/safari.png",
	}

	fileNames := []string{
		"chrome.png",
		"safari.png",
	}

	conf := &godownloadthat.Config{
		Debug: true,
	}

	godownloadthat.RunMultiFast(urls, fileNames, conf)
}

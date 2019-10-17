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
		"chrome_large.png",
		"safari_large.png",
	}

	downloader := godownloadthat.Downloader{}
	downloader.DownloadFiles(urls, fileNames)
}

package godownloadthat_test

import (
	"testing"

	"github.com/abrad1212/godownloadthat"
)

func BenchmarkRunMultiFast(b *testing.B) {
	urls := []string{
		"https://img.icons8.com/cotton/512/000000/chrome.png",
		"https://img.icons8.com/cotton/512/000000/safari.png",
	}

	fileNames := []string{
		"chrome.png",
		"safari.png",
	}

	downloader := godownloadthat.Downloader{
		Debug: false,
	}

	for n := 0; n < b.N; n++ {
		downloader.DownloadFiles(urls, fileNames)
	}
}

func TestRunMultiFast(t *testing.T) {
	// Placeholder
}

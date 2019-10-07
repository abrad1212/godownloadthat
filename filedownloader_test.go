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

	conf := &godownloadthat.Config{
		Debug: false,
	}

	for n := 0; n < b.N; n++ {
		godownloadthat.RunMultiFast(urls, fileNames, conf)
	}
}

func TestRunMultiFast(t *testing.T) {
	// Placeholder
}

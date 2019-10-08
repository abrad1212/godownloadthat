package godownloadthat_test

import (
	"testing"

	"github.com/abrad1212/godownloadthat"
)

func BenchmarkDownloadFiles(b *testing.B) {
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

func TestDownloadFiles(t *testing.T) {
	// Placeholder
	t.Run("Main", func(t *testing.T) {
		downloader := godownloadthat.Downloader{}

		t.Run("1 URL", func(t *testing.T) {
			url := []string{
				"http://pluspng.com/img-png/google-logo-png-open-2000.png",
			}
			fileName := []string{
				"google.png",
			}
			err := downloader.DownloadFiles(url, fileName)

			if err != nil {
				t.Error(err)
			}
		})

		t.Run("Mismatch URL Array Length", func(t *testing.T) {
			urls := []string{
				"http://pluspng.com/img-png/google-logo-png-open-2000.png",
			}
			fileNames := []string{
				"google.png",
				"chrome.png",
			}
			err := downloader.DownloadFiles(urls, fileNames)

			// Expecting the error not to be nil
			if err == nil {
				t.Error("Expected an error, the URL's length doesn't match the filename's length")
			}
		})

		t.Run("Mismatch Filename Array Length", func(t *testing.T) {
			urls := []string{
				"http://pluspng.com/img-png/google-logo-png-open-2000.png",
				"http://pluspng.com/img-png/google-logo-png-open-2000.png",
			}
			fileNames := []string{
				"google.png",
			}
			err := downloader.DownloadFiles(urls, fileNames)

			// Expecting error here
			if err == nil {
				t.Error("Expected an error, the length of the URL array doesn't match the length of the filename array")
			}
		})

		t.Run("URL Doesn't Exist", func(t *testing.T) {
			url := []string{
				"http://randomImageWebsiteOrSomething.com",
			}
			fileName := []string{
				"google.png",
			}
			err := downloader.DownloadFiles(url, fileName)

			if err == nil {
				t.Error("Expected a non 200 status code error")
			}
		})
	})
}

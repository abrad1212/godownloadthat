package godownloadthat

import (
	"bytes"
	"errors"
	//"fmt"
	"io"
	"os"

	//"github.com/h2non/filetype"
	"github.com/valyala/fasthttp"
)

// Downloader implements a downloading client
type Downloader struct {
	Client fasthttp.Client
	Debug bool
}

// DownloadFiles provides a way to download files from
// urls and save them with the specified fileNames
//
// Downloads concurrently using GoRoutine
func (d *Downloader) DownloadFiles(urls []string, fileNames []string) error {
	if len(urls) != len(fileNames) {
		return errors.New("The length of URLs doesn't match the length of filenames")
	}

	done := make(chan []byte, len(urls))
	errch := make(chan error, len(urls))

	for c, url := range urls {
		go func(url string, fileName string) {
			result, err := d.downloadFile(url, fileName)
			if err != nil {
				errch <- err
				done <- nil
				return
			}
			done <- result
			errch <- err

		}(url, fileNames[c])
	}

	var errStr string

	for i := 0; i < len(urls); i++ {
		if err := <- errch; err != nil {
			errStr = errStr + " " + err.Error()
		}
	}

	var err error
	if errStr != "" {
		err = errors.New(errStr)
	}

	return err
}

func (d *Downloader) downloadFile(url string, fileName string) ([]byte, error) {
	statusCode, body, err := d.Client.Get(nil, url)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, errors.New("URL did not return 200")
	}

	out, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	var data bytes.Buffer
	r := bytes.NewReader(body)
	_, err = io.Copy(out, r)

	return data.Bytes(), nil
}
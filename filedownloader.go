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

type Downloader struct {
	Client fasthttp.Client
	Debug bool
}

func (d *Downloader) DownloadFiles(urls []string, fileNames []string) {
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
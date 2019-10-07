package godownloadthat

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"net/http"

	"github.com/h2non/filetype"
	"github.com/valyala/fasthttp"
)

// Config holds options that are used
// when downloading files
type Config struct {
	Debug bool
}

// RunMultiNative - Downloads the contents of the specified urls,
// using GoLang's builtin HTTP library.
func RunMultiNative(urls []string) ([][]byte, error) {
	defer elapsed("RunMultiNative")()
	done := make(chan []byte, len(urls))
	errch := make(chan error, len(urls))

	for count, URL := range urls {
		go func(count int, URL string) {
			countStr := strconv.Itoa(count)

			b, err := downloadFileHelperNative(countStr+".jpg", URL)

			if err != nil {
				errch <- err
				done <- nil
				return
			}

			done <- b
			errch <- nil
		}(count, URL)
	}

	bytesArray := make([][]byte, 0)
	var errStr string

	for i := 0; i < len(urls); i++ {
		bytesArray = append(bytesArray, <-done)
		if err := <-errch; err != nil {
			errStr = errStr + " " + err.Error()
		}
	}
	var err error
	if errStr != "" {
		err = errors.New(errStr)
	}
	return bytesArray, err
}

// RunMultiFast - Downloads the contents of the specified urls,
// [urls], using https://github.com/valyala/fasthttp.
func RunMultiFast(urls []string, fileNames []string, conf *Config) ([][]byte, error) {
	if conf.Debug == true {
		defer elapsed("RunMultiFast")()
	}

	if len(urls) != len(fileNames) {
		return nil, errors.New("Length of URLS doesn't match fileNames")
	}

	done := make(chan []byte, len(urls))
	errch := make(chan error, len(urls))

	for count, URL := range urls {
		go func(fileName string, URL string) {
			b, err := downloadFileHelperFast(fileName, URL, conf)

			if err != nil {
				errch <- err
				done <- nil
				return
			}

			done <- b
			errch <- nil
		}(fileNames[count], URL)
	}

	bytesArray := make([][]byte, 0)
	var errStr string

	for i := 0; i < len(urls); i++ {
		bytesArray = append(bytesArray, <-done)
		if err := <-errch; err != nil {
			errStr = errStr + " " + err.Error()
		}
	}
	var err error
	if errStr != "" {
		err = errors.New(errStr)
	}
	return bytesArray, err
}

func downloadFileHelperNative(filepath string, url string) ([]byte, error) {
	defer elapsed("DownloadFileHelperNative: " + filepath)()

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}
	var data bytes.Buffer
	_, err = io.Copy(&data, response.Body)
	if err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}

func downloadFileHelperFast(filepath string, url string, conf *Config) ([]byte, error) {
	if conf.Debug == true {
		defer elapsed("DownloadFileHelperFast: " + filepath)()
	}

	statusCode, body, err := fasthttp.Get(nil, url)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, errors.New("The URL didn't return 200")
	}

	contentType, err := filetype.Get(body)
	if err != nil {
		// Return if there was an error getting the MIME Type
		return nil, err
	}
	if conf.Debug == true {
		fmt.Printf("MimeType: %s\n", contentType)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}

	var data bytes.Buffer
	r := bytes.NewReader(body)
	_, err = io.Copy(out, r)

	return data.Bytes(), nil
}

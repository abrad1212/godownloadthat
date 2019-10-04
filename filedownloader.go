package godownloadthat

import (
	"fmt"
	"bytes"
	"errors"
	"io"
	"os"
	"strconv"
	"time"

	"net/http"

    "github.com/valyala/fasthttp"
    "github.com/h2non/filetype"
)

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
func RunMultiFast(urls []string) ([][]byte, error) {
	defer elapsed("RunMultiFast")()
	done := make(chan []byte, len(urls))
	errch := make(chan error, len(urls))

	for count, URL := range urls {
		go func(count int, URL string) {
			countStr := strconv.Itoa(count)

			b, err := downloadFileHelperFast(countStr+".jpg", URL)

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

func downloadFileHelperFast(filepath string, url string) ([]byte, error) {
	defer elapsed("DownloadFileHelperFast: " + filepath)()

	statusCode, body, err := fasthttp.Get(nil, url)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, errors.New("The URL didn't return 200")
	}

	//size := int64(binary.Size(body))
	//fmt.Printf("File Size: %s \n", byteCountSI(size))

    contentType, err := filetype.Get(body)
    if err != nil {
        // Return if the MIME Type couldn't be determined
        return nil, err
    }
    fmt.Printf("MimeType: %s", contentType)

	out, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}

	var data bytes.Buffer
	r := bytes.NewReader(body)

	_, err = io.Copy(out, r)
	return data.Bytes(), nil
}

func byteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

// Utility function to measure execution time
// of a function
//  defer elapsed("Function Name")()
func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}


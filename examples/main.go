package main

import (
    //"fmt"
    "github.com/abrad1212/godownloadthat"
)

func main() {
    urls := [2]string{
        "https://img.icons8.com/cotton/512/000000/chrome.png",
        "https://img.icons8.com/cotton/512/000000/safari.png",
    }

    godownloadthat.RunMultiFast(urls[:])
}
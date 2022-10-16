package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	fmt.Println(resp)

	/*
		byteSlice := make([]byte, 99999)
		resp.Body.Read(byteSlice)
		fmt.Println(string(byteSlice))
	*/

	io.Copy(os.Stdout, resp.Body)
}

func (logWriter) Write(byteSlice []byte) (int, error) {

	fmt.Println(string(byteSlice))
	return len(byteSlice), nil
}

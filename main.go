package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func webServe() {
	const port = 3000
	const host = "localhost"

	var url = host + ":" + strconv.Itoa(port)

	fmt.Println("Starting the server on: ", url)
	http.ListenAndServe(url, nil)
}

func main() {

	webServe()
}

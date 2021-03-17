package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
)

func main() {

	var proxy = flag.String("proxy", "localhost:8000", "Proxy dir")
	var bucketName = flag.String("bucket", "", "S3 bucket name.")
	var dir = flag.String("directory", "", "Directory name.")
	flag.Parse()

	if *bucketName == "" {
		fmt.Println("Missing parameters.")
		return
	}

	request := fmt.Sprintf("http://%v/example?bucket=%v&dir=%v", *proxy, *bucketName, *dir)
	resp, err := http.Get(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

}

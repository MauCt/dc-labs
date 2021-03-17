package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//Strut made for test 1 following the example
type onlyBucketCase struct {
	BucketName       string
	ObjectsCount     int
	DirectoriesCount int
	Extensions       map[string]int //Map used because it can have more than 1 extension
}

type bucketAndDirCase struct {
	BucketName       string
	DirectoryName    string
	ObjectsCount     int
	DirectoriesCount int
	Extensions       map[string]int //Map used because it can have more than 1 extension
}

//Identifie the object in a bucket.
//https://docs.aws.amazon.com/AmazonS3/latest/dev-retired/UsingMetadata.html
type Content struct {
	Key string `xml:"Key"`
}

//struct for the XML File from AWS
//https://www.golangprograms.com/golang-write-struct-to-xml-file.html
type XMLResult struct {
	XMLName  xml.Name  `xml:"ListBucketResult"`
	Name     string    `xml:"Name"`
	Contents []Content `xml:"Contents"`
}

func infoSearching(res http.ResponseWriter, req *http.Request) {

	extensions := make(map[string]int)
	directories := make(map[string]bool)
	objects := make(map[string]bool)

	bucketName := req.FormValue("bucket")
	reqDir := req.FormValue("dir")
	url := fmt.Sprintf("https://%v.s3.amazonaws.com", bucketName)

	resp, gerr := http.Get(url)
	if gerr != nil {
		io.WriteString(res, "Get S3 connection error.")
		return
	}
	defer resp.Body.Close()

	data, rerr := ioutil.ReadAll(resp.Body)
	if rerr != nil {
		io.WriteString(res, "Http response reading error.")
		return
	}

	var xmlResult XMLResult
	xerr := xml.Unmarshal(data, &xmlResult)
	if xerr != nil {
		io.WriteString(res, "Permission denied\n")
		return
	}
	limitDtree := 3
	for _, c := range xmlResult.Contents {
		objKey := c.Key
		dir := fmt.Sprintf("%v/", reqDir)

		if reqDir != "" && !strings.HasPrefix(objKey, dir) {
			continue
		}

		if reqDir != "" && strings.HasPrefix(objKey, dir) {
			objKey = strings.Replace(objKey, dir, "", 1)
			if objKey == "" {
				continue
			}
		}

		if strings.Count(objKey, "/") > limitDtree {
			continue
		}
		if strings.HasSuffix(objKey, "/") {
			if !directories[objKey] {
				directories[objKey] = true
			}
		}

		if strings.Contains(objKey, ".") {
			_, exists := objects[objKey]
			if !exists {
				objects[objKey] = true
			}

			ext := strings.Split(objKey, ".")
			_, exists = extensions[ext[len(ext)-1]]
			if !exists {
				extensions[ext[len(ext)-1]] = 1
			}
			if exists {
				extensions[ext[len(ext)-1]] += 1
			}

		}
	}
	//Using both structs to put the xml info in their place and then using MarshalIndent
	//MarshalIndent returns the encoding but with a specific format
	if reqDir == "" {
		bd := onlyBucketCase{
			BucketName:       xmlResult.Name,
			ObjectsCount:     len(objects),
			DirectoriesCount: len(directories),
			Extensions:       extensions,
		}

		out, err := json.MarshalIndent(bd, "", "\t")
		if err != nil {
			panic(err)
		}
		res.Header().Set("Content-Type", "application/json")
		res.Write(out)
	}
	if reqDir != "" {
		dd := bucketAndDirCase{
			BucketName:       xmlResult.Name,
			DirectoryName:    reqDir,
			ObjectsCount:     len(objects),
			DirectoriesCount: len(directories),
			Extensions:       extensions,
		}

		out, err := json.MarshalIndent(dd, "", "\t")
		if err != nil {
			panic(err)
		}
		res.Header().Set("Content-Type", "application/json")
		res.Write(out)
	}
}
func main() {

	var port = flag.Int("port", 9000, "Port numb.")
	flag.Parse()

	socket := fmt.Sprintf("localhost:%v", *port)
	http.HandleFunc("/example", func(res http.ResponseWriter, req *http.Request) {
		infoSearching(res, req)
	})

	err := http.ListenAndServe(socket, nil)
	if err != nil {
		log.Fatal(err)
	}
}

//https://www.example-code.com/golang/s3.asp

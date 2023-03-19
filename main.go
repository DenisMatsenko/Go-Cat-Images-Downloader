package main

import (
	"encoding/json" 
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	var cat = getCat()
	// ? Create path
	var path string = "catImages/cat-" + cat.ID + ".jpg"

	// ? Create file by path
	file, err := os.Create(path)
	errorCheck(err)
    defer file.Close()

	// ? Copy image to created file
    _, err = io.Copy(file, getCatImage(cat))
	errorCheck(err)

	// ? Program end
	fmt.Println("Image saved")
}

// ? Error handling
func errorCheck(err error) {
	if err != nil {log.Fatal("Fatal error -> " + err.Error())}
}

// * Cat staff
type Cat struct {
	ID string `json:"id"`
	URL string `json:"url"`
	Width int `json:"width"`
	Height int `json:"height"`
}

func getCat() Cat {
	var catArray []Cat

	// ? https request
	resp, err := http.Get("https://api.thecatapi.com/v1/images/search")
	errorCheck(err)
	defer resp.Body.Close()

	// ? Decode response to catIage object, error handling
	err = json.NewDecoder(resp.Body).Decode(&catArray)
	errorCheck(err)
	return catArray[0]
}

func getCatImage(cat Cat) io.ReadCloser {
	resp, err := http.Get(cat.URL)
	errorCheck(err)
	return resp.Body
}
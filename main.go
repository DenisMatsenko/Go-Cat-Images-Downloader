package main

import (
	"encoding/json"
	"fmt"
	"io"
	// "io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/google/uuid"
)

func getCatImageURL() string {
	// ? Return catImage URL
	fmt.Println("Getting cat image URL..." + getCatImage().URL)
	return getCatImage().URL
}

type CatImage struct {
	ID string `json:"id"`
	URL string `json:"url"`
	Width int `json:"width"`
	Height int `json:"height"`
}

func getCatImage() CatImage {
	var catImages []CatImage

	// ? https request, error handling
	resp, err := http.Get("https://api.thecatapi.com/v1/images/search")
	if err != nil {log.Fatal("https request error")}
	defer resp.Body.Close()

	//fmt.Println(resp)

	// ? Decode response to catIage object, error handling
	err = json.NewDecoder(resp.Body).Decode(&catImages)
	if err != nil {log.Fatal(err)}
	return catImages[0]
}

func generateId() string {
	// ? Read directory
	// var dir string = "./catImages"
	// files, err := ioutil.ReadDir(dir)
	// if err != nil {log.Fatal(err)}

	// // ? Count files in directory
	// var count int = 0
	// for _, file := range files {
	// 	if !file.Mode().IsRegular() {count++}
	// }

	
    var id uuid.UUID = uuid.New()
	return id.String()
}

func ErrorCheck(err error) {
	if err != nil {log.Fatal(err)}
}

func main() {
	//var countOfCatImages int = getCatImagesCount()
	//fmt.Println(countOfCatImages)

	// ? Create path
	var path string = "catImages/cat" + generateId() + ".jpg"

	resp, err := http.Get(getCatImageURL())
	ErrorCheck(err)
	//err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	file, err := os.Create(path)
	ErrorCheck(err)
    defer file.Close()

    _, err = io.Copy(file, resp.Body)
	ErrorCheck(err)

    // var catFact CatFact
	// // ? Decode data + error handling
	// err = json.NewDecoder(resp.Body).Decode(&catFact)
	// if err != nil {log.Fatal(err)} 

    // fmt.Println(catFact.Fact)
	fmt.Println("Image saved")
}
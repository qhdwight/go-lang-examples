package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type item struct {
	FileFormat string `json:"file_format"`
	S3Uri      string `json:"s3_uri"`
}

type search struct {
	Graph []item `json:"@graph"`
}

func makeRequest(baseUrl string, fields []string, name string) (*search, []byte, error) {
	// Build URL with string builder to avoid excessive copying with concatenation
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s?type=%s&format=json&limit=25", baseUrl, name))
	for _, field := range fields {
		sb.WriteString(fmt.Sprintf("&field=%s", field))
	}
	url := sb.String()
	fmt.Printf("Hitting URL: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	// Deferred functions will fire at end of function
	defer func() { err = resp.Body.Close(); if err != nil { log.Fatal(err) }}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	var search search
	err = json.Unmarshal(body, &search) // Go takes care of putting string into the format of our struct
	if err != nil {
		return nil, nil, err
	}
	return &search, body, err
}

func getJsonFields(object interface{}) []string {
	elem := reflect.ValueOf(object).Elem()
	var fileFields []string
	for i := 0; i < elem.NumField(); i++ {
		fileFields = append(fileFields, elem.Type().Field(i).Tag.Get("json"))
	}
	return fileFields
}

func main() {
	encodeUrl := "https://www.encodeproject.org/search/"
	searchResp, rawResp, err := makeRequest(encodeUrl, getJsonFields(&item{}), "File")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v+\n", searchResp)
	fmt.Println(searchResp.Graph[:2]) // Array slicing!
	err = ioutil.WriteFile("results.json", rawResp, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

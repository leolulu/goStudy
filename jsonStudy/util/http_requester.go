package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func GetDataFromJsonAPISimple() {
	url := "https://animechan.io/api/v1/quotes/random"
	r, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer r.Body.Close()

	content, err := io.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(r.Status)
	fmt.Printf("%T", content)
}

//func getDataFromJsonAPIComplex() {
//	url := "https://animechan.io/api/v1/quotes/random"
//	_, err := http.Post(url)
//	if err != nil {
//		return
//	}
//}

func F2() {
	url := "https://animechan.io/api/v1/quotes/random"
	client := http.Client{Timeout: time.Second * 3}
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(Map2json()))
	if err != nil {
		log.Panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(content))
}

func F3() {
	url := "https://animechan.io/api/v1/quotes/random"
	client := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(content))
	var m map[string]interface{}
	err = json.Unmarshal(content, &m)
	if err != nil {
		log.Panic(err)
	}
	for k, v := range m {
		fmt.Printf("%v: %v\n\n", k, v)
	}
}

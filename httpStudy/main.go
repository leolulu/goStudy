package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	tryHttpClientFull()
}

func tryHttpClientFull() {
	request, err := http.NewRequest(http.MethodGet, "http://www.google.com", nil)
	if err != nil {
		log.Panic(err)
	}
	request.Header.Set("a", "b")
	client := http.Client{Timeout: 3 * time.Second}
	res, err := client.Do(request)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()
	all, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	fmt.Println(string(all))
}

func tryMyReader() {
	bytes.NewBuffer()
}

func tryDownloadBinaryContent() {
	picURL := "https://konachan.net/image/6c5ef06aedc91419426fe4fc739a4715/Konachan.com%20-%20380247%20aqua_eyes%20aqua_hair%20arcaea%20cape%20dress%20eyepatch%20flowers%20gloves%20lobelia_%28saclia%29%20polychromatic%20saya_%28arcaea%29%20short_hair.jpg"
	res, err := http.Get(picURL)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	//content, err := io.ReadAll(res.Body)
	//if err != nil {
	//	log.Panic(err)
	//}
	file, err := os.Create("pic.jpg")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return
	}

}

func tryArgs() {
	args := os.Args
	for _, i := range args {
		fmt.Printf("%T : %v\n", i, i)
	}
}

func tryHttp() error {
	url := "http://43.139.10.228:1127/Saladict/filtered_node.txt1"
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	//defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(res.Status, res.StatusCode, string(body))
	return nil
}

func testSlice(s *[]string) {
	(*s)[0] = "a"
	(*s)[1] = "b"
	*s = append(*s, "c")
	fmt.Println(*s)
}

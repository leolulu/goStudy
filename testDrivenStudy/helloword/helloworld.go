package helloword

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func Hw1() {
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Panic(err)
	}
	os.Stdout = w

	fmt.Println("123")

	os.Stdout = oldStdout

	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(r)
	if err != nil {
		log.Panic(err)
	}
	s := buf.String()
	fmt.Println("caught", s)
}

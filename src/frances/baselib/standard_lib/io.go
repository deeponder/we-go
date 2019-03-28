package main

import (
	"strings"
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
	"io"
)

func fileOpt()  {
	//全部读取
	r := strings.NewReader("hello world.")
	ioutil.ReadAll(r)
	ioutil.ReadFile("path/to/your/file")
	//按行读取
	fi, err := os.Open("path/to/your/file")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a))
	}
}

func main() {



}

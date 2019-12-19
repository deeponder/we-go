package main

import (
	"os/exec"
	"log"
	"fmt"
)

func main()  {
	cmd := exec.Command("python","test.py ")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}

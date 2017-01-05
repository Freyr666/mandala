package main

import (
	"fmt"
	"os"
)

func main()  {
	args := os.Args[1:]
	
	for _, f := range args {
		fmt.Println("File: ", f)
		u := strToUnit(&f)
		v := u.View(15)
		fmt.Print(string(v))
	}
}

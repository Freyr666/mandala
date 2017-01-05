package main

import (
	"fmt"
	"os"
)

func main()  {
	args := os.Args[1:]
	
	for _, f := range args {
		u := strToFile(f)
		v := u.View(15)
		fmt.Print(string(v))
		o := u.Open()
		if o != nil {
			fmt.Print(o)
		}
		//s := u.Status()
		//fmt.Print(s)
		fmt.Print("\n\n")
	}
}

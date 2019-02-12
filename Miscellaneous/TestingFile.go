package main

import (
	"fmt"
	"net/http"
)

func main() {
	fsdf, err := http.Get("https://www.calculator.net/lcm-calculator.html?numberinputs=330,75,450,225&x=48&y=32")

	if err != nil {
		panic(err)
	}

	fmt.Println(fsdf)
}

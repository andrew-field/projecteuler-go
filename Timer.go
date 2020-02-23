package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		panic("There needs to be 1 argument.")
	}

	programToRun := exec.Command(os.Args[1])
	timeBefore := time.Now()
	err := programToRun.Run()
	fmt.Printf("Command finished with error: %v\n", err)
	fmt.Println("Time taken to run ", os.Args[1], " was ", time.Since(timeBefore))
}

package main

import (
	"fmt"
	"time"
)

func main() {
	deploy()
}

// deploy faking a deployment
func deploy() {
	fmt.Println("Deployment action ...")

	time.Sleep(time.Second * 3) // wait 3 seconds

	fmt.Println("Deployment done")
}

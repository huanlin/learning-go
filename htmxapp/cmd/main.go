package main

import (
	"fmt"
	"log"
	"time"

	"github.com/huanlin/learning-go/htmxapp/internal/hardware"
)

func main() {
	fmt.Println("Start running...")
	log.Println("Start running.")

	go func() {
		output, err := hardware.GetSystemInfo()
		if err != nil {
			log.Println("Error: ", err)
			return
		}

		fmt.Println(output)
	}()

	time.Sleep(3 * time.Second)
}

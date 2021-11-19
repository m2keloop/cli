package main

import (
	"cli/cmd"
	"log"
)

func main() {
	err := cmd.Executor()
	if err != nil {
		log.Fatalf("cmd.Executor err:%v", err)
	}
}

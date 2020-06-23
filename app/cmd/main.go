package main

import (
	"os"

	"github.com/mirzaakhena/zapp/app/process"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		process.RunProcess(args[1])
	} else {
		process.RunProcess("skrip.yaml")
	}
}

package main

import (
	"buffalos/src/internal/configurations"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		if args[1] == "--rename" || args[1] == "-rn" {
			configurations.Rename(args)
		} else if args[1] == "--help" || args[1] == "-h" {

		} else if args[1] == "--version" || args[1] == "-v" {

		} else if args[1] == "--install" || args[1] == "-i" {

		} else if args[1] == "--update" || args[1] == "-u" {

		}
	}
}

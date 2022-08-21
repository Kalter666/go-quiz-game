package main

import "flag"

func readArgs() (int, string) {
	limitPtr := flag.Int("limit", 30, "second")
	filenamePtr := flag.String("filename", "problems.csv", "filename")
	flag.Parse()
	return *limitPtr, *filenamePtr
}
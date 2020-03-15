package main

import "log"

func init() {
	log.Println(" 12 factor sample app 0.1")
}

func main() {
	if err := start(); err != nil {
		panic(err)
	}
}

//TODO: sort out the existance checks
//FIXME: Find out why the createSymlink isn't working
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	IsEnvReady()
	PopulateMap()
}

func IsEnvReady() {
	CheckIfEnvSet("ENVIRONMENT")
	CheckIfEnvSet("DOTFILES")
}

func CheckIfEnvSet(key string) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(key + " variable does not exist")
	} else if _, err := os.Stat(value); os.IsNotExist(err) {
		panic(value + " does not exist")
	}
	return true
}

func PopulateMap() {
	dotfiles, exists := os.LookupEnv("DOTFILES")
	if !exists {
		panic("Requried variable does not exist")
	}

	csvFile, _ := os.Open("files.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// CreateSymlink(dotfiles, line[0], line[1])

		fmt.Println(dotfiles, line[0], line[1])
	}
}

// func CreateSymlink(dir, source, target string) {
// 	name := dir, source
// 	println(dir)
// 	println(source)
// 	println(target)
// }

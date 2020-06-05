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

	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {
	dotfiles := CheckIfEnvSet("DOTFILES")
	SymlinkFiles(dotfiles)
}

func CheckIfEnvSet(envvar string) string {
	envvar, exists := os.LookupEnv(envvar)
	if !exists {
		panic(envvar + "Requried variable does not exist")
	} else if _, err := os.Stat(envvar); os.IsNotExist(err) {
		panic(envvar + " directory does not exist")
	}
	return envvar
}

func SymlinkFiles(envvar string) {
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

		CreateSymlink(envvar, line[0], line[1])

	}
}

func CreateSymlink(dir, source, target string) {
	fmt.Println(dir, source, target)
}

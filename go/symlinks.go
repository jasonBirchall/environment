//FIXME: Find out why the createSymlink isn't working
package main

import (
	"bufio"
	"encoding/csv"
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
	// environment := CheckIfEnvSet("ENVIRONMENT")

	PopulateSymlink("files.csv", dotfiles)
	// PopulateSymlink("dir.csv", environment)
}

func CheckIfEnvSet(envvar string) string {
	envvar, exists := os.LookupEnv(envvar)
	if !exists {
		panic(envvar + " requried variable does not exist")
	} else if _, err := os.Stat(envvar); os.IsNotExist(err) {
		panic(envvar + " directory does not exist")
	}
	return envvar
}

func PopulateSymlink(file, envvar string) {
	csvFile, _ := os.Open(file)
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

func CreateSymlink(dir, target, link string) {

	// source := dir + "/" + target

	os.Symlink("/home/json/Documents/workarea/dotfiles/bashrc", "~/.bashrc")

	// println(source, link)
	// out, err := exec.Command("ln", "-h").Output()

	// if err != nil {
	// 	fmt.Printf("%s", err)
	// }

	// output := string(out[:])
	// fmt.Println(output)

	// println("Checking: if symlink already exists")
	// if _, err := os.Lstat(source); err == nil {
	// 	fmt.Errorf("Skipping: symlink already exists: %+v", err)
	// }
	// println("passed")

	// println("Linking: ", source, link)
	// if err := os.Symlink(link, source); err != nil {
	// 	fmt.Errorf("Failed: cannot create symlink", err)
	// } else {
	// 	fmt.Println("Created: ", link)
	// }
	// println("finished")
}

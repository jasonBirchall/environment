package main

import (
	"os"
)

func main() {
	checkIfEnvSet("ENVIRONMENT")
	checkIfEnvSet("DOTFILES")

	// const (
	// 	dotfiles    = os.LookupEnv("DOTFILES")
	// 	environment = os.LookupEnv("ENVIRONMENT")
	// )
	// println(dotfiles, environment)
}

func checkIfEnvSet(key string) {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(key + " variable does not exist")
	}
	createConst(key, value)
}

func createConst(key, value string) {
	const key = value
}

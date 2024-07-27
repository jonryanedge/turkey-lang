package main

import (
	"fmt"
	"os"
	"os/user"
	"turkey-lang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	if user.Username == "jedgeworth" {
		user.Username = "jonryanedge"
	}
	repl.Logo(os.Stdout)
	fmt.Printf("Hello %s! Welcome to the Turkey Programming Language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

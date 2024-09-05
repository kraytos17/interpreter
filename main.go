package main

import (
	"fmt"
	"os"
	"os/user"

	repl "github.com/kraytos17/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Hello %s. This is the monkey PL", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}

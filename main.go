package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/pkg/errors"

	"github.com/dlmiddlecote/monkeylang/repl"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	user, err := user.Current()
	if err != nil {
		return errors.Wrap(err, "getting user")
	}

	fmt.Printf("Hey %s! This is Monkey!\n", user.Username)
	fmt.Printf("Type away ⚡️\n")
	repl.Start(os.Stdin, os.Stdout)

	return nil
}

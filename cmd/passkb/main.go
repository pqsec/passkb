package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pqsec/passkb"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	delay = flag.Int("d", 5, "number of seconds to delay before starting to type")
)

func main() {
	flag.Parse()

	if *delay < 0 {
		fmt.Fprintln(os.Stderr, "Delay cannot be negative")
		os.Exit(1)
	}

	kb, err := passkb.New("passkb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create a virtual keyboard: %v\n", err)
		os.Exit(1)
	}
	defer kb.Close()

	fmt.Print("Enter password: ")
	pass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read the password: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("")

	err = kb.Type(string(pass), time.Duration(*delay)*time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to type the string: %v\n", err)
	}
}

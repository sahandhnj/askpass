package main

import (
	"fmt"
	"log"
	"os"
	"sahand.dev/askpass"
)

func main() {
	filename := "file.pass"
	p := &askpass.Pass{Filename: filename}

	switch {
	case len(os.Args) > 1:
		if len(os.Args) < 3 {
			fmt.Println("You must provide a password to save")
			os.Exit(1)
		}
		password := os.Args[2]
		err := p.Save(password)
		if err != nil {
			log.Fatalf("Error saving password: %v", err)
		}
		fmt.Println("Password saved successfully")
	default:
		password, err := p.Get()
		if err != nil {
			log.Fatalf("Error retrieving password: %v", err)
		}
		fmt.Println(password)
	}
}

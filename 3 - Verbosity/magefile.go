//+build mage

package main

import (
	"errors"
	"log"
	"os"

	"github.com/magefile/mage/mg"
)

func Hello() {
	log.Println("Hi Gophers!")
}

func Errors() error {
	return errors.New("bye gophers!")
}

func Codes() {
	os.Exit(99)
}

func Both() error {
	return mg.Fatal(77, "Goodbye gophers")
}

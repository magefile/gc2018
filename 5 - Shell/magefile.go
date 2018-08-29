//+build mage

package main

import "github.com/magefile/mage/sh"

func Run() error {
	return sh.Run("go", "version")
}

func RunV() error {
	return sh.RunV("go", "version")
}

func Runerr() error {
	return sh.RunV("xxxxx")
}

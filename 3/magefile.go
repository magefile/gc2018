//+build mage

package main

import "github.com/magefile/mage/sh"

func Run() error {
	return sh.RunV("go", "home!")
}

func Out() error {
	s, err := sh.Output("go", "version")
	if err != nil {
		return err
	}
	println(s[11:19])
	return nil
}

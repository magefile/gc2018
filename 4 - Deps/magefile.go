//+build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/target"
)

func Build() {
	mg.Deps(MakeProto, GoGenerate)
	println("building!")
}

func MakeProto() {
	mg.Deps(GoGenerate)
	target.Path()
}

func GoGenerate() error {
}

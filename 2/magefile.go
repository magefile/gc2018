//+build mage

package main

import (
	"github.com/magefile/mage/mg"
)

func Build() {
	mg.Deps(GoGenerate, MakeProto)
	println("building!")
}

func MakeProto() {
	mg.Deps(GoGenerate)
	println("making protos")
}

func GoGenerate() {
	println("generating")
}

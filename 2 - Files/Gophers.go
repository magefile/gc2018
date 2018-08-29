//+build mage

// Mage doesn't care what the filename is, just like go doesn't care.  Mage will grab
// every file in the current directory with the +build mage tag.  Platform
// specifications work in tags and filenames like you're used to.
package main

// Some cool stuff
func SomeTarget() {
	println("hello world!")
}

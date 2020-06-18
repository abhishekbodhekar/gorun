package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		Usage()
		return
	}

}

func CreateFile() {
	ioutil.TempFile(".", "temp")
}
func Usage() {
	fmt.Fprint(os.Stderr, "Usage : gorun [_GoCode_]\n")
	os.Exit(2)
}

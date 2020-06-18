package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const basicCode = "package main \n func main() {\n"

func main() {
	if len(os.Args) != 2 {
		Usage()
		return
	}
	file := CreateFile()
	defer ClearAll(file.Name())
	WriteAndRenameFile(file)

}

func WriteAndRenameFile(file *os.File) *os.File {
	code := basicCode + os.Args[1] + "\n}"
	file.Write([]byte(code))
	err := os.Rename(file.Name(), file.Name()+".go")
	if err != nil {
		fmt.Fprint(os.Stderr, "Can rename temp file : "+err.Error()+"\n")
		os.Exit(2)

	}
	return file

}

func CreateFile() *os.File {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Fprint(os.Stderr, "Can not get currrent directory : "+err.Error()+"\n")
		os.Exit(2)
	}

	file, err := ioutil.TempFile(currentDir, "temp")
	if err != nil {
		fmt.Fprint(os.Stderr, "Can not create temp file : "+err.Error()+"\n")
		os.Exit(2)
	}
	return file

}
func Usage() {
	fmt.Fprint(os.Stderr, "Usage : gorun [_Body_Of_main()_]\n")
	os.Exit(2)
}

func ClearAll(name string) {
	fmt.Println("Clearing, " + name)
	os.Remove(name)
	os.Exit(2)
}

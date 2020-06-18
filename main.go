package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"golang.org/x/tools/imports"
)

const basicCode = "package main \n func main() {\n"

func main() {
	if len(os.Args) != 2 {
		Usage()
		return
	}
	file, err := CreateFile()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	if err = RenameFile(file); err != nil {
		os.Remove(file.Name())
		fmt.Fprint(os.Stderr, err)
		return
	}
	defer ClearAll(file)
	if err = ApplyImports(file); err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	if err = Run(file); err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
}

func Run(file *os.File) error {

	subProcess := exec.Command("go", "run", file.Name()+".go")

	subProcess.Stdout = os.Stdout
	subProcess.Stderr = os.Stderr
	subProcess.Stdin = os.Stdin
	if err := subProcess.Run(); err != nil {
		return errors.New("Can not start sub process : " + err.Error())

	}
	return nil

}

func ApplyImports(file *os.File) error {
	code := basicCode + os.Args[1] + "\n}"
	refinedCode, err := imports.Process(file.Name(), []byte(code), nil)
	if err != nil {
		return errors.New("Can not apply imports to temp file : " + err.Error())
	}
	_, err = file.Write(refinedCode)
	if err != nil {
		return errors.New("Can not write refined code to temp file : " + err.Error())
	}
	return nil
}

func RenameFile(file *os.File) error {
	err := os.Rename(file.Name(), file.Name()+".go")
	if err != nil {
		return errors.New("Can not rename temp file : " + err.Error())
	}
	return nil
}

func CreateFile() (*os.File, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, errors.New("Can not get currrent directory : " + err.Error())
	}
	file, err := ioutil.TempFile(currentDir, "temp")
	if err != nil {
		return nil, errors.New("Can not create temp file : " + err.Error())
	}
	return file, nil
}
func Usage() {
	fmt.Fprint(os.Stderr, `Usage :`+"\t  \t"+` gorun [_Body_Of_main()_]`+"\n"+`Example 1 :`+"\t"+` gorun ' fmt.Println("gorun is fun!") '`+"\n"+`Example 2 :`+"\t"+` gorun ' http.Handle("/", http.FileServer(http.Dir("."))); fmt.Println(http.ListenAndServe(":8080", nil)) '`+"\n")
}

func ClearAll(file *os.File) {
	os.Remove(file.Name() + ".go")
	os.Exit(2)
}

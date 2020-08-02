package main1

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"golang.org/x/tools/imports"
)

// basicCode : A bare minimum code required to run a .go gile
const basicCode = "package main \n func main() {\n"

var filename = ""

func main() {
	if len(os.Args) != 2 {
		Usage()
		return
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go Operate(sigs)
	<-sigs
	os.Remove(filename + ".go")

}

// Operate : Handles high level execution
func Operate(sigs chan os.Signal) {
	file, err := CreateFile()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}
	if err = RenameFile(file); err != nil {
		os.Remove(file.Name()) // removing temp file
		fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}
	filename = file.Name()
	defer func() {
		sigs <- syscall.SIGINT // Notifying that the operation is completed
	}()

	if err = ApplyImports(file); err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	if err = Run(file); err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

}

// Run : Executes "go run ... " on newly creted temp file
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

// ApplyImports : this inserts all the required imports for go code. Also, formats the file.
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

// RenameFile : adding suffix .go to temp file
func RenameFile(file *os.File) error {
	err := os.Rename(file.Name(), file.Name()+".go")
	if err != nil {
		return errors.New("Can not rename temp file : " + err.Error())
	}
	return nil
}

// CreateFile : Creates a temp file in current working directory
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

// Usage : Prints gorun usage
func Usage() {
	fmt.Fprint(os.Stderr, `Usage :`+"\t  \t"+` gorun [_Body_Of_main()_]`+"\n"+`Example 1 :`+"\t"+` gorun ' fmt.Println("gorun is fun!") '`+"\n"+`Example 2 :`+"\t"+` gorun ' http.Handle("/", http.FileServer(http.Dir("."))); fmt.Println(http.ListenAndServe(":8080", nil)) '`+"\n")
}

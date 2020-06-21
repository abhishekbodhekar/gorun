# gorun
 
Go's simple code executioner without .go file.
Executes go code right from CLI.

Similar to python -c.
 
## How can I get it?

```
go get github.com/abhishekbodhekar/gorun
```
## How to use?
After ```go get```,
go to the gorun project and install the binary.
for that, run ```go install``` at root of the project gorun, 

now, hit ```gorun``` and check whether gorun command is installed?

if you see usage of gorun, then all is set and you can use the command now.
Else, check if you have properly set your $GOBIN. Also check if $GOBIN is added in $PATH.

## Examples 
 
Write any Go code (body of main()), gorun will handle your imports and main function for you.

```gorun 'fmt.Println("gorun is fun!")'``` 

You can quickly serve your current directory in one line.
 
```gorun 'http.Handle("/", http.FileServer(http.Dir("."))); fmt.Println(http.ListenAndServe(":8080", nil))'```

Quickly find the date.

```gorun 'fmt.Println(time.Now())'```

# gorun
 
Go's simple code executioner without .go file.
Executes go code right from CLI.

Similar to python -c.
 
## How can I get it?

```
go get github.com/abhishekbodhekar/gorun
```

## Examples 
 
Write any Go code (body of main()), gorun will handle your imports and main function for you.

```gorun 'fmt.Println("Hello gommand")'``` 

You can quickly serve your current directory in one line.
 
```gorun 'http.Handle("/", http.FileServer(http.Dir("."))); fmt.Println(http.ListenAndServe(":8080", nil))'```

Quickly find the date.

```gorun 'fmt.Println(time.Now())'```

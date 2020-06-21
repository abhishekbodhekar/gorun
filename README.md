gorun
Go one liner program similar to python -c

How can I get it?
go get github.com/abhishekbodhekar/gorun
Examples
Write any Go code in a single line context, gorun will handle your imports and main function for you.

gorun 'fmt.Println("Hello gommand")'

You could also quickly serve your current directory in one line.

gorun 'http.Handle("/", http.FileServer(http.Dir("."))); fmt.Println(http.ListenAndServe(":8080", nil))'

Quickly find the date.

gorun 'fmt.Println(time.Now())'

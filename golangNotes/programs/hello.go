package main

import "fmt" // Import fmt package to use Println function

func main() { // Entrypoint, from program execution will start
	fmt.Println("Hello, World!")
}

/*
Commands to run the program -

1. Create a new module. Module path correspond to a repository we plan to
   publish our module.
   `go mod init programs`

	It will create go.mod file - which specify module name and go version we're using.
2. Execute the progam -
   `go run main.go`
*/

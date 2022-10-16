## Introduction to Go

- Programming language created by Google in 2007
- Open sourced in 2009
- Multi-concurrency support
- Designed to run multithreaded concurrent applications that take advantage of
  new performant infrastructure
- Concurrency in Go is cheap & easy
- Main usecase of Go is, to use for performant applications that typically runs on scaled, distributed systems


## Characteristics of Go

- Simple syntax : Easy to learn, read and write code
- Fast build time, start up and run
- Resource Effiecient : Requires fewer resources
- Compiled language : Compiles into single binary (machine code) which can be used across different platforms

## Local setup to write/compile Go code

1. Download & Install Go : Follow instruction from [Go official documentation](https://go.dev/doc/install) to setup Go on your local machine.

2. IDE : Install IDE to write Go code


## Variables and constants in Go

# Variables
- Used for storing values which be changed/updated later if required.
- Give the variable name and reference everywhere in the app
- Update the value at one location and it will get updated where variable is referenced.
- Make our app more flexible
- If variable is declared but not used in the program, compiler shows warning for it.
- Synatx :
    `
        var a = "Hello World"
        a := "Hello World"
        var a string = "Hello World"
    `


# Constants
- Constants are like variables, except that their values cannot be changed once declared
  in the program.

## Pointers in Go

- Pointer is a variable which points to memory address of the another variable.
- Syntax : 
    `&variable_name`


## Print formatted data in Go
- It takes template string & annotation verb that will tell format function how to format 
  the variable passed in.
- Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
- E.g. 
    `fmt.printf("Hello, welcome to golangNotes %s", username}`

- Refer Go official documentation to know more about available [formatting options](https://pkg.go.dev/fmt) for more details.


## How to take user input in Go
- Use fmt.scan() to take user input
- Example:
  `fmt.scan("Enter your name : ")`

## Datatypes in Go

- Go is statically typed language
- We need to tell compiler datatype when declaring the variable
- Syntax :
    `var varname datatype`
- Some of main datatypes :
  `
  - string
  - bool
  - int int8 int16 int32 int64
  - uint uint8(byte) uint16 uint32 uint64 uintptr
  - float32 float64
  - complex64 complex128
  `

## Arrays in Go

- Fixed size single type array
- Syntax
  `var variablename = [size] datatype {comma separated elements}`
  `var variable_name = [size] datatype  <----- Declaring an empty array`

- Adding & accessing element by their index in array

## Slices in Go

- Slice is an abstraction of an array
- Variable sized array
- Can get subarray of its own
- Resized when needed
- Add new element in slice using append() at the end of an slice
- Syntax
  `
  var variablename [] datatype
  var variablename = [] datatype {}
  variablename := [] datatype {}
  `


## Loops in Go

- We only have the for loop in Go
- Syntax :
  `
    1. For loop :
     for {}
    2. For Each Loop :
    for index, index_var := range slice_var{}
    for _, index_var := range slice_var{}

  `

## if else statement in Go
`
if condition {
  // code to be executed if condition is true
} else {
  // code to be executed if condition is false
}
`

## Blank identifier
- Blank identifier `_` used when variable must be declared but will not be used in the program.
- So with Go, we need to make unused variables explicit
## Compare two variables or values in Go
- Use `==` to compare to values and `=` to assign the value to variable

## Break and continue 
- Break will terminate the entire for loop execution and continue with the code right after 
  for loop 
- Continue causes loop to skip the remainder of its body & immediately retesting its condition

## Switch statement
- Allows a variable to be tested for equality against a list of values
- Default case handles the case, if no match is found
- Support multiple values to compare in case statement

## Functions
- Encapsulate code into own container which logically belong together
- Synatx :
  `
    func function_name {

    }

    function_name()   <--- Function calling
  `

## Package level variables
- Defined at the top outside all functions
- They can be accessible from any function
- And in all files which are in same package
- cannot be created with `:=` operator


## Packages in Go
- Go programs are organized into packages
- A package is a collection of Go files
- To use function, variables from another package - declare it with a capital letter

`
Create folder for new packages
export function to use in other package : func First_letter_capital_func_name() {}
import package : module-name/package-name
Use function : package-name.function_name
`

## Varaibles scopes - Local, Package, Global level

## Maps 
- Maps unique values to keys
- Cannot have mixed datatype values
- Syntax :
  `mapVarName := make(map[keyDataType]valuesDataType`

## Struct
- Can hold mixed data types
- Allows to define a structure (which fields to put) of the user type
- Syntax : 
  `
  type structName struct {

  }
  `
- To refer variable from struct use `structName.structVariable`
- Create struct object
  `
  structObj := StructName{properties in key:value format}
  `
- Value receiver methods
  `
  func (structVar StructName) func_name() returnType {

  }
  structObj.func_name()   <---- Calling function with struct object
  `
- Pointer receiver methods
  `
  func (structVar *StructName) func_name() returnType {

  }
  structObj.func_name()   <---- Calling function with struct object

  `

## Maps VS Structs
- Map :
  - All keys must be the same type
  - All values must be the same type
  - Use to represent the collection of related properties
  - Don't need to know all keys at compile time
  - Keys are indexed we can interate over them
  - Reference Type

- Structs :
  - values can be of different type
  - Use to represent a "thing" with a lot of different properties
  - You need to know all the different fields at compile time
  - Keys don't support indexing
  - Value Type

## Goroutines 
- To make function concurrent put keyword `go` in front of it
- `Waitgroup` waits for the launched goroutine to finish
- Package sync provide basic sync functionality
- wg.Add() - Specify number of goroutines to wait for
- wg.Wait() - Blocks until the waitgroup count is 0
- wg.Done() - Decrements waitgroup counter by 1 
              called by the goroutines to indicate that it's finished

- Go uses `Green threads`, abstraction of an actual thread
- Goroutines are cheaper & lightweight as compared to OS threads
- Built-in functionality for Goroutines to talk with each other


## Interfaces
- Similar functionality for functions but function signature can be different
- Example : function for area of circle or rectangle
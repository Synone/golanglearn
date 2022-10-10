# Go CLI

- go build => compile a bunch of go source code files
- go run => compiles and executes one or two files
- go fmt => formats all the code in each file in the current directory
- go install => compile and install a package
- go get => downloads the raw source code of someone else's package
- go test

# Package in learn

1. strings (strings.Join(), strings.Split())
2. bytes
3. os
4.

## Data structure

- Array: fixed length list of things, primitive data structure, can't be resized, rarely used directly
- Slice: An array that can grow and shrink
- Slice: Every element in a slice must be same type
-

## Receiver and Arguments

### Structs

```go
type person struct {
	firstName string
	lastName  string
	contactInfo
}
type contactInfo struct {
	email string
	zipCode int
}
```

### Pointer

1. create a pointer struct to the original struct

   ```go
   structPointer := &originalStruct // note the & character
   ```

2. then the function with receiver:

   ```go
   func (pointerToPerson *person) updateName(newFirstName string){
        (*pointerToPerson).firstName = newFirstName
   }

   structPointer.updateName(newName)
   ```

3. notes on pointer
   | Value Types | Reference Types |
   | :--- | :---: |
   | int | slices |  
    | float | maps |  
    | string | channels |  
    | bool | pointers |  
    | structs | functions |
   > > Value type need to use pointer to change its value
4. [Pointer to a Struct](https://www.geeksforgeeks.org/pointer-to-a-struct-in-golang/)

## Map

1. To initialize the map using the make function (or a map literal) before you can add any elements
   ```go
    m := make(map[string]float64)
    m["pi"] = 3.14
   ```

## Interface:

1. Interfaces are not generic types
2. Interfaces are implicit

## Go routines

- "go" keyword in front of a function will create a new go routine - only infront of a function call
- channel is used for go routines to communicate btw go routines
- channel type <T> only transfer data type <T> with routines
-

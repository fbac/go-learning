# fbac's go resources

- [fbac's go resources](#fbacs-go-resources)
  - [Disclaimer](#disclaimer)
  - [Useful links](#useful-links)
  - [Basics](#basics)
    - [Personal coding guidelines](#personal-coding-guidelines)
    - [Creating a new project](#creating-a-new-project)
    - [Command Line](#command-line)
  - [Code samples](#code-samples)
    - [Types](#types)
    - [Conditions and loops](#conditions-and-loops)
  - [TODO](#todo)

## Disclaimer

This documentation and resources have been collected through the time I've been learning and working with golang, since some years ago.

It's been in my personal laptop for a million years, so in order to get rid of that SPOF I'm creating this repository. If someone finds anything helpful in here that'd be awesome.

I've been updating it from time to time, with the simple purpose of serving as a cheatsheet and braindump some knowledge along the way, as I rarely rely on my (very) volatile memory. Hence this is, by no means, aimed to substitute official go documentation, to be a professional source of knowledge or to be maintained as it's and will be used for personal uses and fast fact-checker.

## Useful links

- [Official documentation](https://go.dev/doc/)
- [Effective go](https://go.dev/doc/effective_go)
- [Standard library](https://pkg.go.dev/std)
- [The Go Blog](https://go.dev/blog/all)
- [Go by Example](https://gobyexample.com)
- [golang-book](https://www.golang-book.com/)
- [web applications with Go](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/index.html)

## Basics

- Go is a statically typed compiled general-purpose language.
- Built-in concurrency through goroutines and channels to communicate them.
- Built-in garbage collection and runtime reflection.
- Powerful standard library and modules.
- Control flow: sequential, loop (iterative), conditional.

### Personal coding guidelines

- `go.mod` and `go.sum` **have** to be versioned.
- Always BET: Benchmark, Example, Test!
- Always format and check the code to meet the standard Go criteria: `gofmt`, `govet`, `golint`
- Always test code and create cover profiles: `go test`, `go test -cover`, `go test -coverprofile=example.out`, `go tool cover -html=example.out`
- Declare variables only when **really** needed, trying always to reduce the memory footprint.
- Keep the scope for each variable as narrow as possible.
  - Always use the short declaration operator inside a code block to declare and assign a value to a variable.
  - If a package level variable is needed, declare the var outside main using the `var <name> <type>` statement.
  - be careful when returning vars from funcs to prevent variables avoiding GC.
- Don't forget the `init()` function, sometimes it's very useful, such as initializing imported packages.
- defer + goroutines are best friends. Use with caution.
- Think about memory: it's usually more performant to allocate everything that you need instead of having the runtime do it dynamically. Create data structures once, not in loops if possible.
- Always write docs and examples.
  - Use docs.go for package docs.
  - Use Examples in test files to write example usage.
  - Check with `go doc [<pkg>.][<sym>.]<method>` and `godoc`, `godoc -http=:PORT`
  - Source code with `go doc -src fmt Println`, `go doc -src fmt`, `go doc cmd/go`

### Creating a new project

- `go mod`
- `go mod init <path>`: create new project
- `go mod tidy`

### Command Line

- `go install`

## Code samples

### Types

- [arrays](builtins/arrays.go)
- [slices](builtins/slices.go)
- [maps](builtins/maps.go)
- [structs](builtins/structs.go)
- [structs - anonymous](builtins/functions-anonymous.go)
- [interfaces & polymorphism](builtins/interfaces-and-polymorphism.go)
- [pointers](builtins/pointers.go)
- [functions & methods](builtins/functions-and-methods.go)
- [functions - anonymous](builtins/functions-anonymous.go)

### Conditions and loops

- if
- for
- [switch](builtins/switch.go)

## TODO

- maybe a go course for beginners in the future? 🤔
- add more lib/utils
- add info about primitives, typed/untyped constant, iota
- add info about aggregated/composite types
- add info about if, switch, for
- more io.Writer examples

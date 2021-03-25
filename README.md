# netroutine [![GoDoc](https://godoc.org/github.com/aidenesco/netroutine?status.svg)](https://godoc.org/github.com/aidenesco/netroutine) [![Go Report Card](https://goreportcard.com/badge/github.com/aidenesco/netroutine)](https://goreportcard.com/report/github.com/aidenesco/netroutine)
This package facilitates the creation and running of network request routines. This is useful for automating web scraping jobs.

## Installation
```sh
go get -u github.com/aidenesco/netroutine
```

## Usage
```go
import "github.com/aidenesco/netroutine"

func main() {
    routine := netroutine.NewRoutine([]netroutine.Runnable{})
    
    env, _ := netroutine.NewEnvironment(map[string]interface{}{})

    routine.Run(env)

    fmt.PrintLn(env.Status)
    //Success
```
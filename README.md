# netroutine [![GoDoc](https://godoc.org/github.com/aidenesco/netroutine?status.svg)](https://godoc.org/github.com/aidenesco/netroutine) [![Go Report Card](https://goreportcard.com/badge/github.com/aidenesco/netroutine)](https://goreportcard.com/report/github.com/aidenesco/netroutine)
This package facilitates the creation and execution of network request routines. This is useful for automating web scraping jobs.

## Installation
```sh
go get -u github.com/aidenesco/netroutine
```

## Usage
```go
import "github.com/aidenesco/netroutine"

func main() {
    routine := netroutine.NewRoutine()
    
    env, _ := netroutine.NewEnvironment(make(map[string]interface{}))

    routine.Run(env)

    fmt.Println(env.Status) // Success
  }
```
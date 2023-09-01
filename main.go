package main

import (
  "fmt"
  "os"
  "io"
)

var out io.Writer = os.Stdout


func main() {
  if len(os.Args) == 1 {
    fmt.Fprintln(out, "\nError: you must pass your name as an arg into this useless CLI")
    fmt.Fprintln(out, "\nUsage: useless <name>")
    fmt.Fprintln(out, "")
    return
  }
  name := os.Args[1]
  fmt.Fprintf(out, "Hello, this is a useless CLI tool. Your name is %s\n", name)
}

package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    if (len(os.Args) <= 1) {
        fmt.Println("Error: please provide the filename as argument")
        return
    }

    filename := os.Args[1]
    f, err := os.Open(filename)
    if (err != nil) {
        panic(err)
    }

    io.Copy(os.Stdout, f)
}

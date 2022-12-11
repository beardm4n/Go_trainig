package main

import (
    "fmt"
    "rsc.io/quote"
    "go_training/greetings"
    "log"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    fmt.Println(quote.Go())

    message, err := greetings.Hello("Amid")

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}

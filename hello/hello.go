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

    message1, err1 := greetings.Hello("Amid")

    if err1 != nil {
        log.Fatal(err1)
    }

    fmt.Println(message1)
    fmt.Println("===================================")

    message2, err2 := greetings.Hello("")

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Println(message2)
}

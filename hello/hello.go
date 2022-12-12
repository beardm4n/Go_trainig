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

    // A slice of names.
    names := []string{"Amid", "Gladys", "Samantha", "Darrin"}

    // Request greeting messages for the names.
    message, err := greetings.Hellos(names)

    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(message)
}

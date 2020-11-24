package main

import (
    "fmt"

    "gitee.com/ouyangshi95/gotest/sdk"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}

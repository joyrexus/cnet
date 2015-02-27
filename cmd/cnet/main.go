package main

import (
    "fmt"
    "flag"

    "github.com/joyrexus/cnet"
)

var (
    user string
    pass string
)

func main() {
    flag.StringVar(&user, "user", "*", "CNet user ID")
    flag.StringVar(&user, "u", "*", "CNet user ID")
    flag.StringVar(&pass, "password", "*", "CNet password")
    flag.StringVar(&pass, "p", "*", "CNet password")
    flag.Parse()

    err := cnet.Authenticate(user, pass)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%s was authenticated!\n", user)
}

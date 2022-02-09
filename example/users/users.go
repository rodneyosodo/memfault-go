package main

import (
    "fmt"
    "os"
    "time"

    mem "github.com/0x6f736f646f/memfault-go/pkg/memfault"
)

var (
    username = os.Getenv("MEMFAULT_USERNAME")
    password = os.Getenv("MEMFAULT_PASSWORD")
)

func main() {
    creds := mem.Credentials{
        Email:    username,
        Password: password,
    }
    d, _ := time.ParseDuration("30s")
    conf := mem.Config{
        ApiURL:          "https://api.memfault.com",
        Credentials:     creds,
        MaxIdleConns:    10,
        IdleConnTimeout: d,
    }
    memfault := mem.NewSDK(conf)
    response1, err := memfault.GetMe()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(response1)

    response2, err := memfault.GenerateUserApiKey()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(response2)

    response3, err := memfault.GetUserApiKey()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(response3)

    response4, err := memfault.DeleteUserApiKey()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(response4)

    response5, err := memfault.GetOrganizationSlug(false)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(response5)

}

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
        APIURL:          "https://api.memfault.com",
        Credentials:     creds,
        MaxIdleConns:    10,
        IdleConnTimeout: d,
    }
    memfault := mem.NewSDK(conf)
    response1, err := memfault.GetMe()
    if err != nil {
        fmt.Println(err)
    }
    // Prints out the user email address
    fmt.Println(response1.Email)

    response2, err := memfault.GenerateUserAPIKey()
    if err != nil {
        fmt.Println(err)
    }
    // Prints out the apikey for the logged in user
    fmt.Println(response2.Data.APIKey)

    response3, err := memfault.GetUserAPIKey()
    if err != nil {
        fmt.Println(err)
    }
    // Prints out the apikey for the logged in user
    fmt.Println(response3.Data.APIKey)

    response4, err := memfault.DeleteUserAPIKey()
    if err != nil {
        fmt.Println(err)
    }
    // print out the response `DELETED`
    fmt.Println(response4)

    // response5, err := memfault.getOrganizationSlug(false)
    // if err != nil {
    //     fmt.Println(err)
    // }
    // fmt.Println(fmt.Sprint(response5))

    // response5, err := memfault.getOrganizationSlug(true)
    // if err != nil {
    //     fmt.Println(err)
    // }
    // fmt.Println(fmt.Sprint(response5))
}

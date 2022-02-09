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

    payload := mem.Project{
        Name:     "SmartSink",
        Slug:     "smartsink",
        Os:       "FreeRTOS",
        Platform: "nRF52",
    }
    response1, err := memfault.CreateProject(payload)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(response1)
}

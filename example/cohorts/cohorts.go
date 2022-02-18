package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	mem "github.com/0x6f736f646f/memfault-go/pkg"
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

	project := mem.Project{
		Name:     "smartsink",
		Slug:     "smartsink",
		Os:       "FreeRTOS",
		Platform: "nRF52",
	}
	cohort := mem.Cohort{
		Name: "demo5",
		Slug: "demo5",
	}
	response1, err := memfault.CreateCohort(project, cohort)
	if err != nil {
		fmt.Println(err)
	} else {
		b, err := json.MarshalIndent(response1, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(b))
	}

	response2, err := memfault.ListCohorts(project)
	if err != nil {
		fmt.Println(err)
	} else {
		b, err := json.MarshalIndent(response2, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(b))
	}
}

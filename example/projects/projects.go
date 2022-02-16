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

	payload := mem.Project{
		Name:     "SmartSdinkamakeabc",
		Slug:     "smartsdinkmskeabc",
		Os:       "FreeRTOS",
		Platform: "nRF52",
	}
	response1, err := memfault.CreateProject(payload)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response1.Data.APIKey)

	response2, err := memfault.ListProject()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response2.Data[0].Slug)

	response3, err := memfault.RetrieveProject(response2.Data[0].Slug)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response3.Data.Slug)

	response4, err := memfault.UpdateProject(payload)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response4.Data.APIKey)

	response6, err := memfault.GetProjectClientKey(response2.Data[0].Slug)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response6)

	response7, err := memfault.RefreshProjectClientKey(response2.Data[0].Slug)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response7)

	response5, err := memfault.DeleteProject(payload)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response5)
}

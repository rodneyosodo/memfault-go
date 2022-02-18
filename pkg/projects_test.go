package memfault

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	username = os.Getenv("MEMFAULT_USERNAME")
	password = os.Getenv("MEMFAULT_PASSWORD")
)

func TestCreateProject(t *testing.T) {
	creds := Credentials{
		Email:    username,
		Password: password,
	}
	d, _ := time.ParseDuration("30s")
	conf := Config{
		APIURL:          "https://api.memfault.com",
		Credentials:     creds,
		MaxIdleConns:    10,
		IdleConnTimeout: d,
	}
	memfault := NewSDK(conf)
	payload := Project{
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

	assert.Equal(t, err, errors.New("A project named `SmartSdinkamakeabc` already exists"), fmt.Sprintf("expected error %s, got nil", err))
}

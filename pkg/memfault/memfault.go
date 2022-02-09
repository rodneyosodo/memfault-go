// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package memfault

import (
    "encoding/base64"
    "errors"
    "io/ioutil"
    "net/http"
    "time"
)

// ContentType represents all possible content types.
type ContentType string

var _ SDK = (*mfSDK)(nil)

// SDK contains Mainflux API.
type SDK interface {
    // GetMe
    GetMe() (string, error)

    // Generate a user api key
    GenerateUserApiKey() (string, error)

    // Get the issued user api key
    GetUserApiKey() (string, error)

    // Deletes the issued user api key
    DeleteUserApiKey() (string, error)

    // Gets the organization slug
    GetOrganizationSlug(multiple bool) (string, error)

    // Creates project
    CreateProject(project Project) (string, error)
}

// Credentials contains the credentials
type Credentials struct {
    Email             string `json:"email"`
    Password          string `json:"password"`
    ApiKey            string `json:"api_key"`
    OrganisationToken string `json:"organisation_token"`
    ProjectKey        string `json:"project_key"`
}

type mfSDK struct {
    apiURL         string
    chunksURL      string
    ingressURL     string
    filesURL       string
    credentials    Credentials
    msgContentType ContentType
    client         *http.Client
}

// Config contains sdk configuration parameters.
type Config struct {
    ApiURL      string
    ChunksURL   string
    IngressURL  string
    FilesURL    string
    Credentials Credentials

    MaxIdleConns    int
    IdleConnTimeout time.Duration
}

type Project struct {
    Name     string `json:"name"`
    Slug     string `json:"slug"`
    Os       string `json:"os"`
    Platform string `json:"platform"`
}

// NewSDK returns new mainflux SDK instance.
func NewSDK(conf Config) SDK {
    return &mfSDK{
        apiURL:      conf.ApiURL,
        chunksURL:   conf.ChunksURL,
        ingressURL:  conf.IngressURL,
        filesURL:    conf.FilesURL,
        credentials: conf.Credentials,

        client: &http.Client{
            Transport: &http.Transport{
                MaxIdleConns:    conf.MaxIdleConns,
                IdleConnTimeout: conf.IdleConnTimeout,
            },
        },
    }
}

func (sdk mfSDK) makeRequest(req *http.Request) ([]byte, error) {
    token, err := sdk.authenticate()
    if err != nil {
        return nil, err
    }
    resp, err := sdk.sendRequest(req, token)
    if err != nil {
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return body, nil
}

func (sdk mfSDK) sendRequest(req *http.Request, token string) (*http.Response, error) {
    if token != "" {
        req.Header.Add("Authorization", "Basic "+token)
    }

    req.Header.Add("Content-Type", "application/json")

    res, err := sdk.client.Do(req)
    if err != nil {
        return res, err
    }
    return res, nil
}

func (sdk mfSDK) authenticate() (string, error) {
    if sdk.credentials.Email != "" && sdk.credentials.Password != "" {
        auth := []byte(sdk.credentials.Email + ":" + sdk.credentials.Password)
        return base64.StdEncoding.EncodeToString(auth), nil
    } else if sdk.credentials.ApiKey != "" && sdk.credentials.Email != "" {
        auth := []byte(sdk.credentials.Email + ":" + sdk.credentials.ApiKey)
        return base64.StdEncoding.EncodeToString(auth), nil
    } else if sdk.credentials.OrganisationToken != "" {
        auth := []byte(":" + sdk.credentials.OrganisationToken)
        return base64.StdEncoding.EncodeToString(auth), nil
    }
    return "", errors.New("Empty credentials")
}

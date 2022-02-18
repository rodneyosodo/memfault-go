// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package memfault

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// ContentType represents all possible content types.
type ContentType string

var _ SDK = (*mfSDK)(nil)

// SDK contains Mainflux API.
type SDK interface {
	// GetMe Return information about the logged in User
	GetMe() (UserRes, error)

	// Generate a user api key
	GenerateUserAPIKey() (UserAPIKeyRes, error)

	// Get a previously generated API Key for the logged in User
	GetUserAPIKey() (UserAPIKeyRes, error)

	// Invalidate the previously generated API Key for the logged in User and do not create another one
	DeleteUserAPIKey() (string, error)

	// Creates a Project under the given Organization
	CreateProject(project Project) (CreateProjectRes, error)

	// List the Projects under a given Organization
	ListProject() (ListProjectRes, error)

	// Retrieves a Project under a given Organization
	RetrieveProject(projectSlug string) (CreateProjectRes, error)

	// Update a Project under a given Organization
	UpdateProject(project Project) (CreateProjectRes, error)

	// Delete a Project under a given Organization
	DeleteProject(project Project) (string, error)

	// Return the Project Client Key
	GetProjectClientKey(projectSlug string) (UserAPIKeyRes, error)

	// Regenerate the Project Client Key
	RefreshProjectClientKey(projectSlug string) (UserAPIKeyRes, error)
	CreateCohort(project Project, cohort Cohort) (CreateCohortRes, error)
	ListCohorts(project Project) (ListCohortRes, error)
}

// Credentials contains the credentials
type Credentials struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	APIKey            string `json:"api_key"`
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
	APIURL      string
	ChunksURL   string
	IngressURL  string
	FilesURL    string
	Credentials Credentials

	MaxIdleConns    int
	IdleConnTimeout time.Duration
}

// Project struct
type Project struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Os       string `json:"os"`
	Platform string `json:"platform"`
}

type Cohort struct {
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

// NewSDK returns new mainflux SDK instance.
func NewSDK(conf Config) SDK {
	return &mfSDK{
		apiURL:      conf.APIURL,
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

func (sdk mfSDK) makeRequest(req *http.Request) ([]byte, int, error) {
	token, err := sdk.authenticate()
	if err != nil {
		return nil, 0, err
	}
	resp, err := sdk.sendRequest(req, token)
	if err != nil {
		return nil, 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	return body, resp.StatusCode, nil
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
	} else if sdk.credentials.APIKey != "" && sdk.credentials.Email != "" {
		auth := []byte(sdk.credentials.Email + ":" + sdk.credentials.APIKey)
		return base64.StdEncoding.EncodeToString(auth), nil
	} else if sdk.credentials.OrganisationToken != "" {
		auth := []byte(":" + sdk.credentials.OrganisationToken)
		return base64.StdEncoding.EncodeToString(auth), nil
	}
	return "", errors.New("Empty credentials")
}

func (sdk mfSDK) getOrganizationSlug(multiple bool) (string, error) {
	endpoint := "auth/me"
	url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	resp, _, err := sdk.makeRequest(req)
	if err != nil {
		return "", err
	}
	var ur UserRes
	if err := json.Unmarshal(resp, &ur); err != nil {
		return "", err
	}
	if multiple {
		var list []string
		for _, organization := range ur.Organizations {
			list = append(list, organization.Slug)
		}
		organizationslug := fmt.Sprint(list)
		return organizationslug, nil
	}
	return ur.Organizations[0].Slug, nil
}

func (sdk mfSDK) getProjectSlugByName(name string) (string, error) {
	projects, err := sdk.ListProject()
	if err != nil {
		return "", nil
	}
	for _, project := range projects.Data {
		if name == project.Name {
			return project.Slug, nil
		}
		return "", nil
	}
	return "", nil
}

func (sdk mfSDK) getProjectSlugByID(id int) (string, error) {
	projects, err := sdk.ListProject()
	if err != nil {
		return "", nil
	}
	for _, project := range projects.Data {
		if id == project.ID {
			return project.Slug, nil
		}
		return "", nil
	}
	return "", nil
}

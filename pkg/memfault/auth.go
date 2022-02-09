package memfault

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
)

// GetMe Return information about the logged in User
func (sdk mfSDK) GetMe() (string, error) {
    endpoint := "auth/me"
    url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return "", err
    }
    resp, err := sdk.makeRequest(req)
    if err != nil {
        return "", err
    }
    return string(resp), err

}

// Generate (or re-generate) an API Key for logged in User.
func (sdk mfSDK) GenerateUserApiKey() (string, error) {
    endpoint := "auth/api_key"
    url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)
    payload := strings.NewReader(``)

    req, err := http.NewRequest(http.MethodPost, url, payload)
    if err != nil {
        return "", err
    }
    resp, err := sdk.makeRequest(req)
    if err != nil {
        return "", err
    }
    var akr UserApiKeyRes
    if err := json.Unmarshal(resp, &akr); err != nil {
        return "", err
    }

    return string(akr.Data.ApiKey), nil
}

// GetUserApiKey Get a previously generated API Key for the logged in User
func (sdk mfSDK) GetUserApiKey() (string, error) {
    endpoint := "auth/api_key"
    url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return "", err
    }
    resp, err := sdk.makeRequest(req)
    if err != nil {
        return "", err
    }
    var akr UserApiKeyRes
    if err := json.Unmarshal(resp, &akr); err != nil {
        return "", err
    }

    return string(akr.Data.ApiKey), nil
}

// DeleteUserApiKey Invalidate the previously generated API Key for the logged in User and do not create another one
func (sdk mfSDK) DeleteUserApiKey() (string, error) {
    endpoint := "auth/api_key"
    url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)
    payload := strings.NewReader(``)

    req, err := http.NewRequest(http.MethodDelete, url, payload)
    if err != nil {
        return "", err
    }
    resp, err := sdk.makeRequest(req)
    if err != nil {
        return "", err
    }
    return string(resp), nil
}

// GetOrganizationSlug
func (sdk mfSDK) GetOrganizationSlug(multiple bool) (string, error) {
    endpoint := "auth/me"
    url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return "", err
    }
    resp, err := sdk.makeRequest(req)
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
        organizationSlugs := fmt.Sprint(list)
        return organizationSlugs, nil
    }
    return string(ur.Organizations[0].Slug), nil
}

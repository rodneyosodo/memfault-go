package memfault

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
)

// GetMe Return information about the logged in User
func (sdk mfSDK) GetMe() (UserRes, error) {
    var ur UserRes
    endpoint := "auth/me"
    url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return ur, err
    }
    resp, err := sdk.makeRequest(req)
    if err != nil {
        return ur, err
    }
    if err := json.Unmarshal(resp, &ur); err != nil {
        return ur, err
    }
    return ur, err

}

// GenerateUserAPIKey generates (or re-generate) an API Key for logged in User.
// To generate this without logging into Memfault, you may use HTTP Basic Auth to call this API.
func (sdk mfSDK) GenerateUserAPIKey() (UserAPIKeyRes, error) {
    var akr UserAPIKeyRes
    endpoint := "auth/api_key"
    url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)
    payload := strings.NewReader(``)

    req, err := http.NewRequest(http.MethodPost, url, payload)
    if err != nil {
        return akr, err
    }
    resp, err := sdk.makeRequest(req)
    if err != nil {
        return akr, err
    }
    if err := json.Unmarshal(resp, &akr); err != nil {
        return akr, err
    }

    return akr, nil
}

// GetUserAPIKey Get a previously generated API Key for the logged in User
func (sdk mfSDK) GetUserAPIKey() (UserAPIKeyRes, error) {
    var akr UserAPIKeyRes
    endpoint := "auth/api_key"
    url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return akr, err
    }
    resp, err := sdk.makeRequest(req)
    if err != nil {
        return akr, err
    }
    if err := json.Unmarshal(resp, &akr); err != nil {
        return akr, err
    }

    return akr, nil
}

// DeleteUserAPIKey Invalidate the previously generated API Key for the logged in User and do not create another one
func (sdk mfSDK) DeleteUserAPIKey() (string, error) {
    endpoint := "auth/api_key"
    url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)
    payload := strings.NewReader(``)

    req, err := http.NewRequest(http.MethodDelete, url, payload)
    if err != nil {
        return "", err
    }
    _, err = sdk.makeRequest(req)
    if err != nil {
        return "", err
    }
    return "DELETED", nil
}

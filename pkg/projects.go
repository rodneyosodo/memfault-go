package memfault

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// CreateProject creates a Project under the given Organization
func (sdk mfSDK) CreateProject(project Project) (CreateProjectRes, error) {
	var exists bool
	slug, err := sdk.getProjectSlugByName(project.Name)
	if err != nil {
		return cpr, nil
	}
	if slug == project.Slug {
		exists = true
	}
	data, err := json.Marshal(project)
	if err != nil {
		return cpr, err
	}
	payload := strings.NewReader(string(data))
	organizationslug, err := sdk.getOrganizationSlug(false)
	if err != nil {
		return cpr, err
	}
	endpoint := "api/v0/organizations/" + string(organizationslug) + "/projects"
	url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return cpr, err
	}
	resp, _, err := sdk.makeRequest(req)
	if err != nil {
		return cpr, err
	}
	if exists {
		if err := json.Unmarshal(resp, &er); err != nil {
			return cpr, err
		}
		return cpr, errors.New(string(er.Error.Message))
	}

	if err := json.Unmarshal(resp, &cpr); err != nil {
		return cpr, err
	}
	return cpr, nil
}

// ListProject list the Projects under a given Organization
func (sdk mfSDK) ListProject() (ListProjectRes, error) {
	organizationslug, err := sdk.getOrganizationSlug(false)
	if err != nil {
		return lpr, err
	}
	endpoint := "api/v0/organizations/" + string(organizationslug) + "/projects"
	url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return lpr, err
	}
	resp, _, err := sdk.makeRequest(req)
	if err != nil {
		return lpr, err
	}
	if err := json.Unmarshal(resp, &lpr); err != nil {
		return lpr, err
	}
	return lpr, nil
}

// RetrieveProject retrieve a Project under a given Organization
func (sdk mfSDK) RetrieveProject(projectSlug string) (CreateProjectRes, error) {
	organizationslug, err := sdk.getOrganizationSlug(false)
	if err != nil {
		return cpr, err
	}
	endpoint := "api/v0/organizations/" + string(organizationslug) + "/projects/" + projectSlug
	url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return cpr, err
	}
	resp, _, err := sdk.makeRequest(req)
	if err != nil {
		return cpr, err
	}
	if err := json.Unmarshal(resp, &cpr); err != nil {
		return cpr, err
	}
	return cpr, nil
}

// UpdateProject updates a Project under a given Organization
func (sdk mfSDK) UpdateProject(project Project) (CreateProjectRes, error) {
	organizationslug, err := sdk.getOrganizationSlug(false)
	if err != nil {
		return cpr, err
	}
	projectSlug, err := sdk.getProjectSlugByName(project.Name)
	if err != nil {
		return cpr, nil
	}
	data, err := json.Marshal(project)
	if err != nil {
		return cpr, err
	}
	payload := strings.NewReader(string(data))

	endpoint := "api/v0/organizations/" + string(organizationslug) + "/projects/" + projectSlug
	url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

	req, err := http.NewRequest(http.MethodPatch, url, payload)
	if err != nil {
		return cpr, err
	}
	resp, _, err := sdk.makeRequest(req)
	if err != nil {
		return cpr, err
	}
	if err := json.Unmarshal(resp, &cpr); err != nil {
		return cpr, err
	}
	return cpr, nil
}

// DeleteProject Delete a Project under a given Organization
func (sdk mfSDK) DeleteProject(project Project) (string, error) {
	organizationslug, err := sdk.getOrganizationSlug(false)
	if err != nil {
		return "", err
	}
	projectSlug, err := sdk.getProjectSlugByName(project.Name)
	if err != nil {
		return "", nil
	}

	endpoint := "api/v0/organizations/" + string(organizationslug) + "/projects/" + projectSlug
	url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

	payload := strings.NewReader(``)

	req, err := http.NewRequest(http.MethodDelete, url, payload)
	if err != nil {
		return "", err
	}
	_, _, err = sdk.makeRequest(req)
	if err != nil {
		return "", err
	}
	return "DELETED", nil
}

// GetProjectClientKey Return the Project Client Key
func (sdk mfSDK) GetProjectClientKey(projectSlug string) (UserAPIKeyRes, error) {
	organizationslug, err := sdk.getOrganizationSlug(false)
	if err != nil {
		return uakr, err
	}
	endpoint := "api/v0/organizations/" + string(organizationslug) + "/projects/" + projectSlug + "/api_key"
	url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return uakr, err
	}
	resp, _, err := sdk.makeRequest(req)
	if err != nil {
		return uakr, err
	}
	if err := json.Unmarshal(resp, &uakr); err != nil {
		return uakr, err
	}
	return uakr, nil
}

// RefreshProjectClientKey Regenerate the Project Client Key
func (sdk mfSDK) RefreshProjectClientKey(projectSlug string) (UserAPIKeyRes, error) {
	organizationslug, err := sdk.getOrganizationSlug(false)
	if err != nil {
		return uakr, err
	}
	endpoint := "api/v0/organizations/" + string(organizationslug) + "/projects/" + projectSlug + "/api_key"
	url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

	payload := strings.NewReader(``)

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return uakr, err
	}
	resp, _, err := sdk.makeRequest(req)
	if err != nil {
		return uakr, err
	}
	if err := json.Unmarshal(resp, &uakr); err != nil {
		return uakr, err
	}
	return uakr, nil
}

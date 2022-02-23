package memfault

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// CreateCohort creates a Cohort for the given Project
func (sdk mfSDK) CreateCohort(project Project, cohort Cohort) (CreateCohortRes, error) {
	slug, err := sdk.getProjectSlugByName(project.Name)
	if err != nil {
		return ccr, err
	}
	organizationslug, err := sdk.getOrganizationSlug(false)
	if err != nil {
		return ccr, err
	}
	data, err := json.Marshal(cohort)
	if err != nil {
		return ccr, err
	}
	payload := strings.NewReader(string(data))
	endpoint := "api/v0/organizations/" + string(organizationslug) + "/projects/" + string(slug) + "/cohorts"
	url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return ccr, err
	}
	resp, statusCode, err := sdk.makeRequest(req)
	if err != nil {
		return ccr, err
	}
	if statusCode != 200 {
		if err := json.Unmarshal(resp, &er); err != nil {
			return ccr, err
		}
		return ccr, errors.New(string(er.Error.Message))
	}
	if err := json.Unmarshal(resp, &ccr); err != nil {
		return ccr, err
	}
	return ccr, nil
}

// ListCohorts List all Cohorts for the given Project
func (sdk mfSDK) ListCohorts(project Project) (ListCohortRes, error) {
	slug, err := sdk.getProjectSlugByName(project.Name)
	if err != nil {
		return lcr, err
	}
	organizationslug, err := sdk.getOrganizationSlug(false)
	if err != nil {
		return lcr, err
	}
	endpoint := "api/v0/organizations/" + string(organizationslug) + "/projects/" + string(slug) + "/cohorts"
	url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return lcr, err
	}
	resp, statusCode, err := sdk.makeRequest(req)
	if err != nil {
		return lcr, err
	}
	if statusCode != 200 {
		if err := json.Unmarshal(resp, &er); err != nil {
			return lcr, err
		}
		return lcr, errors.New(string(er.Error.Message))
	}
	if err := json.Unmarshal(resp, &lcr); err != nil {
		return lcr, err
	}
	return lcr, nil
}

// Retrieve a single Cohort
func (sdk mfSDK) RetrieveCohorts(project Project, cohort Cohort) (CreateCohortRes, error) {
	projectslug, err := sdk.getProjectSlugByName(project.Name)
	if err != nil {
		return ccr, err
	}
	organizationslug, err := sdk.getOrganizationSlug(false)
	if err != nil {
		return ccr, err
	}
	cohortslug, err := sdk.getCohortSlug(project, cohort.Name)
	if err != nil {
		return ccr, err
	}
	endpoint := "api/v0/organizations/" + string(organizationslug) + "/projects/" + string(projectslug) + "/cohorts/" + string(cohortslug)
	url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ccr, err
	}
	resp, statusCode, err := sdk.makeRequest(req)
	if err != nil {
		return ccr, err
	}
	if statusCode != 200 {
		if err := json.Unmarshal(resp, &er); err != nil {
			return ccr, err
		}
		return ccr, errors.New(string(er.Error.Message))
	}
	if err := json.Unmarshal(resp, &ccr); err != nil {
		return ccr, err
	}
	return ccr, nil
}

// UpdateCohorts Update a single Cohort
func (sdk mfSDK) UpdateCohorts(project Project, cohort Cohort, cohortslug string) (CreateCohortRes, error) {
	projectslug, err := sdk.getProjectSlugByName(project.Name)
	if err != nil {
		return ccr, err
	}
	organizationslug, err := sdk.getOrganizationSlug(false)
	if err != nil {
		return ccr, err
	}
	data, err := json.Marshal(cohort)
	if err != nil {
		return ccr, err
	}
	payload := strings.NewReader(string(data))
	endpoint := "api/v0/organizations/" + string(organizationslug) + "/projects/" + string(projectslug) + "/cohorts/" + string(cohortslug)
	url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

	req, err := http.NewRequest(http.MethodPatch, url, payload)
	if err != nil {
		return ccr, err
	}
	resp, statusCode, err := sdk.makeRequest(req)
	if err != nil {
		return ccr, err
	}
	if statusCode != 200 {
		if err := json.Unmarshal(resp, &er); err != nil {
			return ccr, err
		}
		return ccr, errors.New(string(er.Error.Message))
	}
	if err := json.Unmarshal(resp, &ccr); err != nil {
		return ccr, err
	}
	return ccr, nil
}

//DeleteCohorts Delete a single Cohort
func (sdk mfSDK) DeleteCohorts(project Project, cohort Cohort) (string, error) {
	projectslug, err := sdk.getProjectSlugByName(project.Name)
	if err != nil {
		return "", err
	}
	organizationslug, err := sdk.getOrganizationSlug(false)
	if err != nil {
		return "", err
	}
	cohortslug, err := sdk.getCohortSlug(project, cohort.Name)
	if err != nil {
		return "", err
	}
	endpoint := "api/v0/organizations/" + string(organizationslug) + "/projects/" + string(projectslug) + "/cohorts/" + string(cohortslug)
	url := fmt.Sprintf("%s/%s", sdk.apiURL, endpoint)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return "", err
	}
	_, _, err = sdk.makeRequest(req)
	if err != nil {
		return "", err
	}
	return "DELETED", nil
}

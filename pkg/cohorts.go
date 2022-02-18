package memfault

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	ccr CreateCohortRes
	lcr ListCohortRes
	er  ErrorRes
)

// CreateCohort creates a Cohort for the given Project
func (sdk mfSDK) CreateCohort(project Project, cohort Cohort) (CreateCohortRes, error) {
	var (
		ccr CreateCohortRes
		er  ErrorRes
	)
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

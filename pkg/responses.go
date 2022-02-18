package memfault

import "time"

// APIKey provides the apikey
type APIKey struct {
	APIKey string `json:"api_key,omitempty"`
}

// UserAPIKeyRes is information about the logged in user api key
type UserAPIKeyRes struct {
	Data APIKey `json:"data,omitempty"`
}

// provides information about the organisation
type organization struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

// provides the organisation access control list
type organizationACL struct {
	Admin           bool         `json:"admin,omitempty"`
	CustomerSupport bool         `json:"customer_support,omitempty"`
	Organization    organization `json:"organization,omitempty"`
}

// UserRes is information about the logged in User
type UserRes struct {
	CreatedDate       time.Time         `json:"created_date,omitempty"`
	Email             string            `json:"email,omitempty"`
	ID                int               `json:"id,omitempty"`
	Name              string            `json:"name,omitempty"`
	Organizations     []organization    `json:"organizations,omitempty"`
	UpdatedDate       time.Time         `json:"updated_date,omitempty"`
	PrivacyNotice     bool              `json:"has_accepted_privacy_notice,omitempty"`
	TermsOfService    bool              `json:"has_accepted_terms_of_service,omitempty"`
	Impersonated      string            `json:"impersonator,omitempty"`
	OrganizationsAcls []organizationACL `json:"organization_acls,omitempty"`
}

// intergrations of the project
type integrations struct {
	Name string `json:"name,omitempty"`
}

// device link strategies
type deviceLinkStrategies struct {
	Name string `json:"name,omitempty"`
}

// setup datails for the project
type setupDetails struct {
	OperatingSystem           []string `json:"operating_system,omitempty"`
	CompilerOther             string   `json:"compiler_other,omitempty"`
	BuildToolchain            []string `json:"build_toolchain,omitempty"`
	BuildToolchainOther       string   `json:"build_toolchain_other,omitempty"`
	AdditionalChips           string   `json:"additional_chips,omitempty"`
	ProgrammingLanguagesOther string   `json:"programming_languages_other,omitempty"`
	DebuggingToolchain        []string `json:"debugging_toolchain,omitempty"`
	ConnectivityOther         string   `json:"connectivity_other,omitempty"`
	PrimaryChip               string   `json:"primary_chip,omitempty"`
	Compiler                  []string `json:"compiler,omitempty"`
	Connectivity              []string `json:"connectivity,omitempty"`
	PrimaryChipOther          string   `json:"primary_chip_other,omitempty"`
	DebuggingToolchainOther   string   `json:"debugging_toolchain_other,omitempty"`
	GmsCertification          string   `json:"gms_certification,omitempty"`
	OperatingSystemOther      string   `json:"operating_system_other,omitempty"`
	ProgrammingLanguages      []string `json:"programming_languages,omitempty"`
}

// Details to the project
type projectData struct {
	CreatedDate            time.Time              `json:"created_date,omitempty"`
	Slug                   string                 `json:"slug,omitempty"`
	LegacyLatestAPIEnabled bool                   `json:"legacy_latest_api_enabled,omitempty"`
	Integrations           []integrations         `json:"integrations,omitempty"`
	AlertsEnabled          bool                   `json:"alerts_enabled,omitempty"`
	UpdatedDate            time.Time              `json:"updated_date,omitempty"`
	CountDeployments       int                    `json:"count_deployments,omitempty"`
	DashboardEnabled       bool                   `json:"dashboard_enabled,omitempty"`
	DeviceLinkStrategies   []deviceLinkStrategies `json:"device_link_strategies,omitempty"`
	PrimarySoftwareType    string                 `json:"primary_software_type,omitempty"`
	BugReportEnabled       bool                   `json:"bugreport_export_enabled,omitempty"`
	MetricChartEnabled     bool                   `json:"metric_chart_enabled,omitempty"`
	ChunksDebugLogLength   int                    `json:"chunks_debug_log_length,omitempty"`
	EventDebugLogLength    int                    `json:"event_debug_log_length,omitempty"`
	APIKey                 string                 `json:"api_key,omitempty"`
	MultiComponentEnabled  bool                   `json:"multi_component_enabled,omitempty"`
	SetupDetails           setupDetails           `json:"setup_details,omitempty"`
	VersioningScheme       string                 `json:"versioning_scheme,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	Os                     string                 `json:"os,omitempty"`
	Platform               string                 `json:"platform,omitempty"`
	ID                     int                    `json:"id,omitempty"`
}

// CreateProjectRes thirs is a new project created
type CreateProjectRes struct {
	Data projectData `json:"data,omitempty"`
}

// project error message
type errorInformation struct {
	Code      int    `json:"code,omitempty"`
	HTTPCode  int    `json:"http_code,omitempty"`
	Message   string `json:"message,omitempty"`
	ErrorType string `json:"type,omitempty"`
}

// ErrorRes eror message that project already exist
type ErrorRes struct {
	Error errorInformation `json:"error,omitempty"`
}

type paging struct {
	PerPage    int `json:"per_page,omitempty"`
	TotalCount int `json:"total_count,omitempty"`
	PageCount  int `json:"page_count,omitempty"`
	Page       int `json:"page,omitempty"`
	ItemCount  int `json:"item_count,omitempty"`
}

// ListProjectRes struct
type ListProjectRes struct {
	Data   []projectData `json:"data,omitempty"`
	Paging paging        `json:"paging,omitempty"`
}

type CohortData struct {
	ByPassPrimarySoftwareTypeCheck bool      `json:"bypass_primary_software_type_check,omitempty"`
	ByPassVersionChecks            bool      `json:"bypass_version_checks,omitempty"`
	CountActiveDeployments         int       `json:"count_active_deployments,omitempty"`
	CountDevices                   int       `json:"count_devices,omitempty"`
	CountStagedDevices             int       `json:"count_staged_devices,omitempty"`
	CreatedDate                    time.Time `json:"created_date,omitempty"`
	ID                             int       `json:"id,omitempty"`
	IgnoreTraces                   bool      `json:"ignore_traces,omitempty"`
	LastDeployment                 time.Time `json:"last_deployment,omitempty"`
	Name                           string    `json:"name,omitempty"`
	ReturnEqualLatestReleases      bool      `json:"return_equal_latest_releases,omitempty"`
	Slug                           string    `json:"slug,omitempty"`
	UpdatedDate                    time.Time `json:"updated_date,omitempty"`
	UpdatesEnabled                 bool      `json:"updates_enabled,omitempty"`
}

// CreateCohortRes struct
type CreateCohortRes struct {
	Data CohortData `json:"data,omitempty"`
}

// ListCohortRes struct
type ListCohortRes struct {
	Data []CohortData `json:"data,omitempty"`
}

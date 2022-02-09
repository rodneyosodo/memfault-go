package memfault

import "time"

type apiKey struct {
	ApiKey string `json:"api_key,omitempty"`
}
type UserApiKeyRes struct {
	Data apiKey `json:"data,omitempty"`
}

type organization struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug"`
}

type organizationAcl struct {
	Admin           bool         `json:"admin,omitempty"`
	CustomerSupport bool         `json:"customer_support,omitempty"`
	Organization    organization `json:"organization,omitempty"`
}
type organizations []organization

type UserRes struct {
	CreatedDate       time.Time         `json:"created_date,omitempty"`
	Email             string            `json:"email,omitempty"`
	Id                int               `json:"id,omitempty"`
	Name              string            `json:"name,omitempty"`
	Organizations     organizations     `json:"organizations,omitempty"`
	UpdatedDate       time.Time         `json:"updated_date,omitempty"`
	PrivacyNotice     bool              `json:"has_accepted_privacy_notice"`
	TermsOfService    bool              `json:"has_accepted_terms_of_service"`
	Impersonated      string            `json:"impersonator"`
	OrganizationsAcls []organizationAcl `json:"organization_acls"`
}

// type integrations struct {
// 	name string `json:"name,omitempty"`
// }
// type projectData struct {
// 	CreatedDate time.Time `json:"created_date,omitempty"`
// 	Slug string `json:"slug,omitempty": "smartsink"`
//     LegacyLatestApiEnabled bool `json:"legacy_latest_api_enabled"`
//     Integrations []integrations `json:"integrations"`
//     AlertsEnabled bool `json:"alerts_enabled,omitempty"`
//     UpdatedDate time.Time `json:"updated_date,omitempty"`
//     CountDeployments int `json:"count_deployments"`
//     DashboardEnabled bool `json:"dashboard_enabled"`
//     DeviceLinkStrategies [] `json:"device_link_strategies"`
// 	PrimarySoftwareType string `json:"primary_software_type"`
// 	BugReportEnabled bool `json:"bugreport_export_enabled"`
//     MetricChartEnabled bool `json:"metric_chart_enabled"`
//         // "chunks_debug_log_length": 0,
//         // "api_key": "3kfvYHivnNzFGtlKnNdUAAhETQU77Ldp",
//         // "multi_component_enabled": false,

//     VersioningScheme string `json:"versioning_scheme"`
//     Name string `json:"name"`
//     Os       string `json:"os"`
//     Platform string `json:"platform"`
//     Id int `json:"id"`
// }
// type CreateProjectRes struct {
// 	Data projectData `json:"data,omitempty"`
// }

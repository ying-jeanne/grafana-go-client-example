/*
 * Grafana HTTP API.
 *
 * The Grafana backend exposes an HTTP API, the same API is used by the frontend to do everything from saving dashboards, creating users and updating data sources.
 *
 * API version: 0.0.1
 * Contact: hello@grafana.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package goclient

type CustomPermissionsRecordDtoModel struct {
	CustomPermissions string `json:"customPermissions,omitempty"`
	GranteeName       string `json:"granteeName,omitempty"`
	GranteeType       string `json:"granteeType,omitempty"`
	GranteeUrl        string `json:"granteeUrl,omitempty"`
	Id                int64  `json:"id,omitempty"`
	IsFolder          bool   `json:"isFolder,omitempty"`
	OrgId             int64  `json:"orgId,omitempty"`
	OrgRole           string `json:"orgRole,omitempty"`
	Slug              string `json:"slug,omitempty"`
	Title             string `json:"title,omitempty"`
	Uid               string `json:"uid,omitempty"`
	Url               string `json:"url,omitempty"`
	UsersCount        int64  `json:"usersCount,omitempty"`
}

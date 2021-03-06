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

type SearchUserQueryResultModel struct {
	Page       int64                   `json:"page,omitempty"`
	PerPage    int64                   `json:"perPage,omitempty"`
	TotalCount int64                   `json:"totalCount,omitempty"`
	Users      []UserSearchHitDtoModel `json:"users,omitempty"`
}

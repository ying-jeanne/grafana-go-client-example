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

type QueryHistoryDtoModel struct {
	Comment       string     `json:"comment,omitempty"`
	CreatedAt     int64      `json:"createdAt,omitempty"`
	CreatedBy     int64      `json:"createdBy,omitempty"`
	DatasourceUid string     `json:"datasourceUid,omitempty"`
	Queries       *JsonModel `json:"queries,omitempty"`
	Starred       bool       `json:"starred,omitempty"`
	Uid           string     `json:"uid,omitempty"`
}

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

type PostAnnotationsCmdModel struct {
	DashboardId  int64      `json:"dashboardId,omitempty"`
	DashboardUID string     `json:"dashboardUID,omitempty"`
	Data         *JsonModel `json:"data,omitempty"`
	PanelId      int64      `json:"panelId,omitempty"`
	Tags         []string   `json:"tags,omitempty"`
	Text         string     `json:"text,omitempty"`
	Time         int64      `json:"time,omitempty"`
	TimeEnd      int64      `json:"timeEnd,omitempty"`
}

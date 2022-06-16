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

type PostGraphiteAnnotationsCmdModel struct {
	Data string      `json:"data,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	What string      `json:"what,omitempty"`
	When int64       `json:"when,omitempty"`
}

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

type InlineResponse2009Model struct {
	// ID Identifier of the deleted folder.
	Id int64 `json:"id"`
	// Message Message of the deleted folder.
	Message string `json:"message"`
	// Title of the deleted folder.
	Title string `json:"title"`
}

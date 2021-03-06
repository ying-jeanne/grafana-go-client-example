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

// AlertStatus alert status
type AlertStatusModel struct {
	// inhibited by
	InhibitedBy []string `json:"inhibitedBy"`
	// silenced by
	SilencedBy []string `json:"silencedBy"`
	// state
	State string `json:"state"`
}

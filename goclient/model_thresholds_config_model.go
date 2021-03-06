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

// ThresholdsConfig setup thresholds
type ThresholdsConfigModel struct {
	Mode *ThresholdsModeModel `json:"mode,omitempty"`
	// Must be sorted by 'value', first value is always -Infinity
	Steps []ThresholdModel `json:"steps,omitempty"`
}

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

type GettableApiAlertingConfigModel struct {
	Global              *GlobalConfigModel         `json:"global,omitempty"`
	InhibitRules        []InhibitRuleModel         `json:"inhibit_rules,omitempty"`
	MuteTimeProvenances map[string]ProvenanceModel `json:"muteTimeProvenances,omitempty"`
	MuteTimeIntervals   []MuteTimeIntervalModel    `json:"mute_time_intervals,omitempty"`
	// Override with our superset receiver type
	Receivers []GettableApiReceiverModel `json:"receivers,omitempty"`
	Route     *RouteModel                `json:"route,omitempty"`
	Templates []string                   `json:"templates,omitempty"`
}

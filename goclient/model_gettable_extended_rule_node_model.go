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

type GettableExtendedRuleNodeModel struct {
	Alert        string                    `json:"alert,omitempty"`
	Annotations  map[string]string         `json:"annotations,omitempty"`
	Expr         string                    `json:"expr,omitempty"`
	For_         *DurationModel            `json:"for,omitempty"`
	GrafanaAlert *GettableGrafanaRuleModel `json:"grafana_alert,omitempty"`
	Labels       map[string]string         `json:"labels,omitempty"`
	Record       string                    `json:"record,omitempty"`
}

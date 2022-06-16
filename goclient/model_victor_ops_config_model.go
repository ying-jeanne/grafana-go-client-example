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

type VictorOpsConfigModel struct {
	ApiKey            *SecretModel           `json:"api_key,omitempty"`
	ApiKeyFile        *SecretModel           `json:"api_key_file,omitempty"`
	ApiUrl            *UrlModel              `json:"api_url,omitempty"`
	CustomFields      map[string]string      `json:"custom_fields,omitempty"`
	EntityDisplayName string                 `json:"entity_display_name,omitempty"`
	HttpConfig        *HttpClientConfigModel `json:"http_config,omitempty"`
	MessageType       string                 `json:"message_type,omitempty"`
	MonitoringTool    string                 `json:"monitoring_tool,omitempty"`
	RoutingKey        string                 `json:"routing_key,omitempty"`
	SendResolved      bool                   `json:"send_resolved,omitempty"`
	StateMessage      string                 `json:"state_message,omitempty"`
}

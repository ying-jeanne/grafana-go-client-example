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

type OpsGenieConfigModel struct {
	Actions      string                         `json:"actions,omitempty"`
	ApiKey       *SecretModel                   `json:"api_key,omitempty"`
	ApiKeyFile   string                         `json:"api_key_file,omitempty"`
	ApiUrl       *UrlModel                      `json:"api_url,omitempty"`
	Description  string                         `json:"description,omitempty"`
	Details      map[string]string              `json:"details,omitempty"`
	Entity       string                         `json:"entity,omitempty"`
	HttpConfig   *HttpClientConfigModel         `json:"http_config,omitempty"`
	Message      string                         `json:"message,omitempty"`
	Note         string                         `json:"note,omitempty"`
	Priority     string                         `json:"priority,omitempty"`
	Responders   []OpsGenieConfigResponderModel `json:"responders,omitempty"`
	SendResolved bool                           `json:"send_resolved,omitempty"`
	Source       string                         `json:"source,omitempty"`
	Tags         string                         `json:"tags,omitempty"`
	UpdateAlerts bool                           `json:"update_alerts,omitempty"`
}

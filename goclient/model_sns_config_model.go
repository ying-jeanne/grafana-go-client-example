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

type SnsConfigModel struct {
	ApiUrl       string                 `json:"api_url,omitempty"`
	Attributes   map[string]string      `json:"attributes,omitempty"`
	HttpConfig   *HttpClientConfigModel `json:"http_config,omitempty"`
	Message      string                 `json:"message,omitempty"`
	PhoneNumber  string                 `json:"phone_number,omitempty"`
	SendResolved bool                   `json:"send_resolved,omitempty"`
	Sigv4        *SigV4ConfigModel      `json:"sigv4,omitempty"`
	Subject      string                 `json:"subject,omitempty"`
	TargetArn    string                 `json:"target_arn,omitempty"`
	TopicArn     string                 `json:"topic_arn,omitempty"`
}

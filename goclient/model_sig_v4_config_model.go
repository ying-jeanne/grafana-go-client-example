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

// SigV4Config is the configuration for signing remote write requests with AWS's SigV4 verification process. Empty values will be retrieved using the AWS default credentials chain.
type SigV4ConfigModel struct {
	AccessKey string       `json:"AccessKey,omitempty"`
	Profile   string       `json:"Profile,omitempty"`
	Region    string       `json:"Region,omitempty"`
	RoleARN   string       `json:"RoleARN,omitempty"`
	SecretKey *SecretModel `json:"SecretKey,omitempty"`
}

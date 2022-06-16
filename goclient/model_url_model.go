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

type UrlModel struct {
	ForceQuery  bool           `json:"ForceQuery,omitempty"`
	Fragment    string         `json:"Fragment,omitempty"`
	Host        string         `json:"Host,omitempty"`
	Opaque      string         `json:"Opaque,omitempty"`
	Path        string         `json:"Path,omitempty"`
	RawFragment string         `json:"RawFragment,omitempty"`
	RawPath     string         `json:"RawPath,omitempty"`
	RawQuery    string         `json:"RawQuery,omitempty"`
	Scheme      string         `json:"Scheme,omitempty"`
	User        *UserinfoModel `json:"User,omitempty"`
}

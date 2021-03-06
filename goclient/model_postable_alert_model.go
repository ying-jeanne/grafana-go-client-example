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

import (
	"time"
)

// PostableAlert postable alert
type PostableAlertModel struct {
	Annotations *LabelSetModel `json:"annotations,omitempty"`
	// ends at Format: date-time
	EndsAt time.Time `json:"endsAt,omitempty"`
	// generator URL Format: uri
	GeneratorURL string         `json:"generatorURL,omitempty"`
	Labels       *LabelSetModel `json:"labels"`
	// starts at Format: date-time
	StartsAt time.Time `json:"startsAt,omitempty"`
}

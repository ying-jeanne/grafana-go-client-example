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

type GettableAlertModel struct {
	Annotations *LabelSetModel `json:"annotations"`
	// ends at
	EndsAt time.Time `json:"endsAt"`
	// fingerprint
	Fingerprint string `json:"fingerprint"`
	// generator URL Format: uri
	GeneratorURL string         `json:"generatorURL,omitempty"`
	Labels       *LabelSetModel `json:"labels"`
	// receivers
	Receivers []ReceiverModel `json:"receivers"`
	// starts at
	StartsAt time.Time         `json:"startsAt"`
	Status   *AlertStatusModel `json:"status"`
	// updated at
	UpdatedAt time.Time `json:"updatedAt"`
}

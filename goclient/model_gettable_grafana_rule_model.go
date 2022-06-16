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

type GettableGrafanaRuleModel struct {
	Condition       string            `json:"condition,omitempty"`
	Data            []AlertQueryModel `json:"data,omitempty"`
	ExecErrState    string            `json:"exec_err_state,omitempty"`
	Id              int64             `json:"id,omitempty"`
	IntervalSeconds int64             `json:"intervalSeconds,omitempty"`
	NamespaceId     int64             `json:"namespace_id,omitempty"`
	NamespaceUid    string            `json:"namespace_uid,omitempty"`
	NoDataState     string            `json:"no_data_state,omitempty"`
	OrgId           int64             `json:"orgId,omitempty"`
	Provenance      *ProvenanceModel  `json:"provenance,omitempty"`
	RuleGroup       string            `json:"rule_group,omitempty"`
	Title           string            `json:"title,omitempty"`
	Uid             string            `json:"uid,omitempty"`
	Updated         time.Time         `json:"updated,omitempty"`
	Version         int64             `json:"version,omitempty"`
}

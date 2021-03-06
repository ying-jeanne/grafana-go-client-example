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

type AlertRuleModel struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	Condition   string            `json:"condition"`
	Data        []AlertQueryModel `json:"data"`
	//  Alerting AlertingErrState Error ErrorErrState OK OkErrState
	ExecErrState string            `json:"execErrState"`
	FolderUID    string            `json:"folderUID"`
	For_         *DurationModel    `json:"for"`
	Id           int64             `json:"id,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"`
	//  Alerting Alerting NoData NoData OK OK
	NoDataState string           `json:"noDataState"`
	OrgID       int64            `json:"orgID"`
	Provenance  *ProvenanceModel `json:"provenance,omitempty"`
	RuleGroup   string           `json:"ruleGroup"`
	Title       string           `json:"title"`
	Uid         string           `json:"uid,omitempty"`
	Updated     time.Time        `json:"updated,omitempty"`
}

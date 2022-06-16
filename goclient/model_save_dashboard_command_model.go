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

type SaveDashboardCommandModel struct {
	UpdatedAt time.Time  `json:"UpdatedAt,omitempty"`
	Dashboard *JsonModel `json:"dashboard,omitempty"`
	FolderId  int64      `json:"folderId,omitempty"`
	FolderUid string     `json:"folderUid,omitempty"`
	IsFolder  bool       `json:"isFolder,omitempty"`
	Message   string     `json:"message,omitempty"`
	Overwrite bool       `json:"overwrite,omitempty"`
	UserId    int64      `json:"userId,omitempty"`
}

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

type ImportDashboardRequestModel struct {
	Dashboard *JsonModel                  `json:"dashboard,omitempty"`
	FolderId  int64                       `json:"folderId,omitempty"`
	FolderUid string                      `json:"folderUid,omitempty"`
	Inputs    []ImportDashboardInputModel `json:"inputs,omitempty"`
	Overwrite bool                        `json:"overwrite,omitempty"`
	Path      string                      `json:"path,omitempty"`
	PluginId  string                      `json:"pluginId,omitempty"`
}

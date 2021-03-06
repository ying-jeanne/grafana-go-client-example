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

type ImportDashboardResponseModel struct {
	DashboardId      int64  `json:"dashboardId,omitempty"`
	Description      string `json:"description,omitempty"`
	FolderId         int64  `json:"folderId,omitempty"`
	Imported         bool   `json:"imported,omitempty"`
	ImportedRevision int64  `json:"importedRevision,omitempty"`
	ImportedUri      string `json:"importedUri,omitempty"`
	ImportedUrl      string `json:"importedUrl,omitempty"`
	Path             string `json:"path,omitempty"`
	PluginId         string `json:"pluginId,omitempty"`
	Removed          bool   `json:"removed,omitempty"`
	Revision         int64  `json:"revision,omitempty"`
	Slug             string `json:"slug,omitempty"`
	Title            string `json:"title,omitempty"`
	Uid              string `json:"uid,omitempty"`
}

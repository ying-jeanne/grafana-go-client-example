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

// A Field is essentially a slice of various types with extra properties and methods. See NewField() for supported types.  The slice data in the Field is a not exported, so methods on the Field are used to to manipulate its data.
type FieldModel struct {
	Config *FieldConfigModel `json:"config,omitempty"`
	Labels *LabelsModel      `json:"labels,omitempty"`
	// Name is default identifier of the field. The name does not have to be unique, but the combination of name and Labels should be unique for proper behavior in all situations.
	Name string `json:"name,omitempty"`
}
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

// A FrameType string, when present in a frame's metadata, asserts that the frame's structure conforms to the FrameType's specification. This property is currently optional, so FrameType may be FrameTypeUnknown even if the properties of the Frame correspond to a defined FrameType.
type FrameTypeModel struct {
}

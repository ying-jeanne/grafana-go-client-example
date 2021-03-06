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

type HttpClientConfigModel struct {
	Authorization *AuthorizationModel `json:"authorization,omitempty"`
	BasicAuth     *BasicAuthModel     `json:"basic_auth,omitempty"`
	BearerToken   *SecretModel        `json:"bearer_token,omitempty"`
	// The bearer token file for the targets. Deprecated in favour of Authorization.CredentialsFile.
	BearerTokenFile string `json:"bearer_token_file,omitempty"`
	// FollowRedirects specifies whether the client should follow HTTP 3xx redirects. The omitempty flag is not set, because it would be hidden from the marshalled configuration when set to false.
	FollowRedirects bool            `json:"follow_redirects,omitempty"`
	Oauth2          *OAuth2Model    `json:"oauth2,omitempty"`
	ProxyUrl        *UrlModel       `json:"proxy_url,omitempty"`
	TlsConfig       *TlsConfigModel `json:"tls_config,omitempty"`
}

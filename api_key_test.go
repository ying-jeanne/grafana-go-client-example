package gapi

import (
	"testing"

	"github.com/ying-jeanne/grafana-go-client-example/goclient"
)

const (
	createAPIKeyJSON = `{"name":"key-name", "key":"mock-api-key"}`
	deleteAPIKeyJSON = `{"message":"API key deleted"}`

	getAPIKeysJSON = `[
		{
			"id": 1,
			"name": "key-name-2",
			"role": "Viewer"
		},
		{
			"id": 2,
			"name": "key-name-2",
			"role": "Admin",
			"expiration": "2021-10-30T10:52:03+03:00"
		}
	]`
)

func TestCreateAPIKey(t *testing.T) {
	cfg := goclient.NewConfiguration()
	client := goclient.NewAPIClient(cfg)

	ctx := util.GetContextWithBasicAuth()

	_, res, err := client.ApiKeysApi.AddAPIkey(ctx, goclient.AddApiKeyCommandModel{
		Name: "key-name",
		Role: "Viewer",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(pretty.PrettyFormat(res))
}

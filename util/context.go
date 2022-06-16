package util

import (
	"context"

	"github.com/ying-jeanne/grafana-go-client-example/goclient"
)

func GetContextWithBasicAuth() context.Context {
	return context.WithValue(context.Background(), goclient.ContextBasicAuth, goclient.BasicAuth{
		UserName: "admin",
		Password: "admin",
	})
}

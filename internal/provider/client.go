package provider

import (
	"time"

	"terraform-provider-nd/internal/registry"

	nd "github.com/netascode/go-nd"
)

var _ registry.ClientProvider = (*NDClient)(nil)

type NDClient struct {
	URL       string
	Username  string
	Password  string
	Domain    string
	Insecure  bool
	Timeout   time.Duration
	ApiClient *nd.Client
	NDModules map[string]interface{}
}

func (c *NDClient) GetModule(name string) interface{} {
	return c.NDModules[name]
}

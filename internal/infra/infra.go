package infra

import (
	nd "github.com/netascode/go-nd"
)

const ModuleKey = "infra"

type NexusDashboardInfra struct {
	ApiClient *nd.Client
}

var infraInstance *NexusDashboardInfra

func NewInfra(client *nd.Client) *NexusDashboardInfra {
	if infraInstance == nil {
		infraInstance = &NexusDashboardInfra{
			ApiClient: client,
		}
	}
	return infraInstance
}

package manage

import (
	nd "github.com/netascode/go-nd"
)

const ModuleKey = "manage"

type NexusDashboardManage struct {
	ApiClient *nd.Client
}

var manageInstance *NexusDashboardManage

func NewManage(client *nd.Client) *NexusDashboardManage {
	if manageInstance == nil {
		manageInstance = &NexusDashboardManage{
			ApiClient: client,
		}
	}
	return manageInstance
}

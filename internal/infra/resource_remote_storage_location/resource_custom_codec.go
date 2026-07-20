package resource_remote_storage_location

import "encoding/json"

type customNDFCRemoteStorageLocationModel NDFCRemoteStorageLocationModel

type remoteStorageLocationWrappedResponse struct {
	Spec *customNDFCRemoteStorageLocationModel `json:"spec"`
}

func (m *NDFCRemoteStorageLocationModel) UnmarshalJSON(data []byte) error {
	var wrapped remoteStorageLocationWrappedResponse
	if err := json.Unmarshal(data, &wrapped); err == nil && wrapped.Spec != nil {
		*m = NDFCRemoteStorageLocationModel(*wrapped.Spec)
	}
	return nil
}

func (m NDFCRemoteStorageLocationModel) MarshalJSON() ([]byte, error) {
	if m.StorageLocationType != "nfs" {
		if m.Authentication.SshKey != "" {
			m.Authentication.AuthenticationType = "key"
		} else {
			m.Authentication.AuthenticationType = "password"
		}
	}
	return json.Marshal(customNDFCRemoteStorageLocationModel(m))
}

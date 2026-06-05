package resource_remote_storage_location

import (
	"encoding/json"
	"errors"
)

type customNDRemoteStorageLocationModel NDFCRemoteStorageLocationModel

type remoteStorageLocationWrappedResponse struct {
	Spec *customNDRemoteStorageLocationModel `json:"spec"`
}

func (m *NDFCRemoteStorageLocationModel) UnmarshalJSON(data []byte) error {
	var wrapped remoteStorageLocationWrappedResponse
	err := json.Unmarshal(data, &wrapped)
	if err != nil {
		return err
	}

	if wrapped.Spec != nil {
		*m = NDFCRemoteStorageLocationModel(*wrapped.Spec)
	} else {
		return errors.New(string(data))
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
	return json.Marshal(customNDRemoteStorageLocationModel(m))
}
